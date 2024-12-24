#latest image tag
FROM bitnami/jupyter-base-notebook:latest

#working dir
WORKDIR /home/jovyan/work

#permission
RUN mkdir -p /home/jovyan/work && \
    chown -R 1000:1000 /home/jovyan/work \

#setting up env variables
COPY --chmod=775 setup_env.sh /usr/local/bin/setup_env.sh
ENTRYPOINT ["/usr/local/bin/setup_env.sh"]

CMD jupyter notebook \
    --ip=0.0.0.0 \
    --NotebookApp.port=$JUPYTER_PORT \
   # --NotebookApp.token='' \
   # --NotebookApp.password='' \
    --no-browser \
    --allow-root \

