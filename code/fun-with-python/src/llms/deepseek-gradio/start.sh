#!/usr/bin/env bash

folder="text-generation-webui"
if [  -e "$folder" ]; then
    echo "folder exits"
    exit 0
fi

# clone the repo
git clone https://github.com/oobabooga/"${folder}".git
cd text-generation-webui || exit

# install dependencies based on your system
pip3 install -r requirements_apple_silicon.txt


# create venv
python3 -m venv llm_env
source llm_env/bin/activate

# install hugging face and cli
pip3 install huggingface_hub
pip3 install -U "huggingface_hub[cli]"

# download the model
huggingface-cli download TheBloke/Mistral-7B-Instruct-v0.2-GGUF --local-dir ./models/

deactivate

# run call the text generation ui using deep seek model
python server.py --model deepseek-llm-7b

echo "access the ui on " + http://localhost:7860/
