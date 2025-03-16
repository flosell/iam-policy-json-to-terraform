#!/usr/bin/env bash

set -eu -o pipefail

# TODO: make sure the following is installed:
# https://github.com/mtdowling/chag
# $GITHUB_TOKEN is set # TODO: use gh for that? 

REPO="iam-policy-json-to-terraform"

SCRIPT_DIR=$(dirname "$0")
VERSION="$1"

if [ -z "$VERSION" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

cd ${SCRIPT_DIR}/..

echo "Checking GitHub Login..."
gh auth status | grep -E 'Token scopes.*workflow'
if [ $? -ne 0 ]; then
  echo "Please login to GitHub using 'gh auth login' and afterwards, ensure the token has the 'workflow' scope using 'gh auth refresh -s workflow'"
  exit 1
fi

echo "Checking GPG signing capabilities..."
echo "test" | gpg --sign > /dev/null

if [ $? -ne 0 ]; then
  echo "Please configure GPG signing or ensure the key is not expired"
  exit 1
fi


sed -i "" -e "s/const AppVersion = .*/const AppVersion = \"${VERSION}\"/g" iam-policy-json-to-terraform.go
git commit -m "Release ${VERSION}: Update AppVersion constant" iam-policy-json-to-terraform.go

./go clean
./go cli_build

chag update $VERSION
git commit -m "Release ${VERSION}: Update CHANGELOG.md" CHANGELOG.md
chag tag --sign

git push
git push --tags

gh release create ${VERSION} \
    --title ${VERSION} \
    --notes "$(chag contents)"

for i in "${REPO}.exe" "${REPO}_alpine" "${REPO}_amd64" "${REPO}_darwin" "${REPO}_darwin_arm"; do
  echo "Uploading ${i}..."
  gh release upload ${VERSION} ${i}
done

HOMEBREW_GITHUB_API_TOKEN="$(gh auth token)" # homebrew often uses a readonly token, set the one already used for release instead
export HOMEBREW_GITHUB_API_TOKEN

archive_url="https://github.com/flosell/iam-policy-json-to-terraform/archive/refs/tags/${VERSION}.tar.gz"
sha=$(curl -sSLf "${archive_url}" | sha256sum | awk '{print $1}')
brew bump-formula-pr --strict "iam-policy-json-to-terraform" \
                     --url "${archive_url}" \
                     --sha256 "${sha}"
