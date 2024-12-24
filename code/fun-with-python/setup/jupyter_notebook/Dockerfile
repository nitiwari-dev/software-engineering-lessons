#latest image tag
FROM bitnami/jupyter-base-notebook:latest

USER root
#working dir
WORKDIR /home/jovyan/work

#permission
RUN chown -R 1000:1000 /home/jovyan/work

#setting up env variables
COPY setup_env.sh /usr/local/bin/setup_env.sh
RUN chmod 775 /usr/local/bin/setup_env.sh
ENTRYPOINT ["/usr/local/bin/setup_env.sh"]

#Change the user
USER 1000

CMD jupyter notebook \
    --ip=0.0.0.0 \
    --NotebookApp.port=$JUPYTER_PORT \
   # --NotebookApp.token='' \
   # --NotebookApp.password='' \
    --no-browser \
    --allow-root \

