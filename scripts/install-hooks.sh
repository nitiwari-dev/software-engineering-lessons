#!/usr/bin/env sh

# Find the git directory
GIT_DIR=$(git rev-parse --git-dir)

# Create symlink
ln -s scripts/pre-commit.sh "$GIT_DIR"/hooks/pre-commit

echo "Pre commit hook installed!"

