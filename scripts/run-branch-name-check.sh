#!/usr/bin/env sh
echo "  Check local branch name..."

local_branch=$(git rev-parse --abbrev-ref HEAD)
accepted_branch="^((hotfix|conflict|bumpversion|revert|bug|fix|release|doc)(\/|-)[A-Za-z0-9._-]+$)"

#check if the branch matching accepted branch pattern
if [[ ! $local_branch =~ $accepted_branch ]];
then
  echo "Branch name should start with one below accepted format in regex!"
  echo "$accepted_branch"
  exit 1
fi