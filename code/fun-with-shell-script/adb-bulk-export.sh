#!/usr/bin/env bash
# A script to bulk export files from specific folder into local directory
# Usage: ./adb-bulk-export.sh <remote_folder> <local_folder>
# Example: ./adb-bulk-export.sh /sdcard/DCIM/Camera ./CameraBackup

REMOTE_FOLDER=$1
LOCAL_FOLDER=$2

if [ -z "$REMOTE_FOLDER" ] || [ -z "$LOCAL_FOLDER" ]; then
  echo "Usage: $0 <remote_folder> <local_folder>"
  exit 1
fi

# Prepare the list of files to pull
adb shell "ls $REMOTE_FOLDER | grep -E '^(2022)'" > filelist.txt

count=$(wc -l < filelist.txt)
extracted=0
batch_size=10
increment_size=0

echo "Found $count files to pull from $REMOTE_FOLDER to $LOCAL_FOLDER"
for f in $(cat filelist.txt); do
  echo "Pulling $REMOTE_FOLDER/$f to $LOCAL_FOLDER/$f"
  adb pull "$REMOTE_FOLDER/$f" "$LOCAL_FOLDER/$f" &
  ((increment_size++))
  if ((increment_size++ % batch_size == 0)); then
      wait
  fi

  ((extracted++))
  echo "Extracted $extracted out of $count files"
done

wait
