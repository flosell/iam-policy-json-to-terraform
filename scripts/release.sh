#!/usr/bin/env bash

set -eu -o pipefail

# TODO: make sure the following is installed:
# https://github.com/aktau/github-release
# https://github.com/mtdowling/chag
# $GITHUB_TOKEN is set

USER="flosell"
REPO="iam-policy-json-to-terraform"

SCRIPT_DIR=$(dirname "$0")
VERSION="$1"

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

cd ${SCRIPT_DIR}/..

sed -i "" -e "s/const AppVersion = .*/const AppVersion = \"${VERSION}\"/g" iam-policy-json-to-terraform.go
git commit -m "Release ${VERSION}: Update AppVersion constant" iam-policy-json-to-terraform.go

make clean build

chag update $VERSION
git commit -m "Release ${VERSION}: Update CHANGELOG.md" CHANGELOG.md
chag tag --sign

git push
git push --tags


github-release release \
    --user ${USER} \
    --repo ${REPO} \
    --tag ${VERSION} \
    --name ${VERSION} \
    --description "$(chag contents)"

for i in "${REPO}.exe" "${REPO}_alpine" "${REPO}_amd64" "${REPO}_darwin"; do
  echo "Uploading ${i}..."
  github-release upload \
      --user ${USER} \
      --repo ${REPO} \
      --tag ${VERSION} \
      --name ${i} \
      --file ${i}
done

export HOMEBREW_GITHUB_API_TOKEN="${GITHUB_TOKEN}" # homebrew often uses a readonly token, set the one already used for release instead

archive_url="https://github.com/flosell/iam-policy-json-to-terraform/archive/${VERSION}.tar.gz"
sha=$(curl -sSLf "${archive_url}" | openssl sha256)
brew bump-formula-pr --strict "iam-policy-json-to-terraform" \
                     --url "${archive_url}" \
                     --sha256 "${sha}"
