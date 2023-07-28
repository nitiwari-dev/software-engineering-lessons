#!/usr/bin/env sh

local_branch=$(git rev-parse --abbrev-ref HEAD)
accepted_branch_name="HEAD|feature|hotfix|conflict|bumpversion|revert|bug|fix|release|doc"
accepted_branch_regex="^(($accepted_branch_name)(\/|-)[A-Za-z0-9._-]+|^main$)"

if [[ ! $local_branch =~ $accepted_branch_regex ]];
then
  echo "Branch name \`$local_branch\` is not in accepted format viz - $accepted_branch_name"
  exit 1
fi