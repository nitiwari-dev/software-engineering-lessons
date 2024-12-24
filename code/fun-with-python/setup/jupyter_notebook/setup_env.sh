#!/usr/bin/env bash

JUPYTER_PORT=9001
JUPYTER_IMAGE_NAME="jupyter_nb"

export JUPYTER_PORT=${JUPYTER_PORT}
export JUPYTER_IMAGE_NAME=${JUPYTER_IMAGE_NAME}
exec "$@" # any command after this will have the env set
