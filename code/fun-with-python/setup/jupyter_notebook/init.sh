#!/usr/bin/env bash

#setup env variables
chmod +x setup_env.sh
source setup_env.sh
sh setup_env.sh

notebook_dir="$(pwd)/nb"
docker build -t ${JUPYTER_IMAGE_NAME} .

docker run \
-p ${JUPYTER_PORT}:${JUPYTER_PORT} \
-v "$notebook_dir":/home/jovyan/work \
${JUPYTER_IMAGE_NAME}