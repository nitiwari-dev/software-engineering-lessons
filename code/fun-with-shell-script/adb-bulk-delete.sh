#!/usr/bin/env bash
# A script to bulk delete files from specific folder into local directory
# Usage: ./adb-bulk-delete.sh <remote_folder>
# Example: ./adb-bulk-delete.sh /sdcard/DCIM/Camera

REMOTE_FOLDER=$1

if [ -z "$REMOTE_FOLDER" ] ; then
  echo "Usage: $0 <remote_folder> <local_folder>"
  exit 1
fi

# Prepare the list of files to pull
adb shell "ls $REMOTE_FOLDER | grep -E '^(2022)'" > filelist.txt

count=$(wc -l < filelist.txt)
deleted_files_count=0
batch_size=10
increment_size=0
echo "Found $count files to pull from $REMOTE_FOLDER"
for f in $(cat filelist.txt); do
  echo "Deleting $REMOTE_FOLDER/$f"
  adb shell rm "$REMOTE_FOLDER/$f" &
  ((increment_size++))
  if ((increment_size % batch_size == 0)); then
      wait
  fi

  ((deleted_files_count++))
  echo "Deleted $deleted_files_count out of $count files"
done
wait
