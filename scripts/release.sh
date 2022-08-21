#!/bin/bash

set -e

VERSION=$1

if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

cd npm

yarn version --new-version $VERSION

cd ..

git commit -am "chore: release version $VERSION"
git tag -s -a v$VERSION -m "chore: release version $VERSION"
git push
git push origin v$VERSION
