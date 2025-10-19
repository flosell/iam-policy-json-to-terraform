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
gh auth status | grep -E 'Token scopes.*workflow' || no_correctly_scoped_token=1
if [ ${no_correctly_scoped_token:-0} -eq 1 ]; then
  echo "Please login to GitHub using 'gh auth login' and afterwards, ensure the token has the 'workflow' scope using 'gh auth refresh -s workflow'"
  exit 1
fi

echo "Checking GPG signing capabilities..."
echo "test" | gpg --sign > /dev/null || gpg_error=1

if [ ${gpg_error:-0} -eq 1 ]; then
  echo "Please configure GPG signing or ensure the key is not expired."
  echo "You might also need to install and configure pinentry-mac to unlock your gpg key without interaction on the terminal."
  exit 1
fi

chag --version || chag_error=1
if [ ${chag_error:-0} -eq 1 ]; then
  echo "Please install chag from https://github.com/mtdowling/chag"
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

for i in "${REPO}.exe" "${REPO}_alpine" "${REPO}_amd64" "${REPO}_arm64" "${REPO}_darwin" "${REPO}_darwin_arm"; do
  echo "Uploading ${i}..."
  gh release upload ${VERSION} ${i}
done

echo "Release to GitHub done."
echo "Homebrew should auto-update after a few hours but better double-check this."