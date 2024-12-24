#Arguments
ARG NON_ROOT_USER_CODE=1001
#latest image tag
FROM bitnami/jupyter-base-notebook:latest
#change the user to root for admin access to install packages
USER root
#install dependencies
RUN conda install --quiet --yes \
    'matplotlib-base' \
    'scipy' && \
    conda clean --all -f -y

## switch back to non-root user
USER $NON_ROOT_USER_CODE
CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=9001", "--no-browser", "--allow-root"]
