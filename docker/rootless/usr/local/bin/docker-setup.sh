#!/bin/bash

# Prepare git folder
mkdir -p ${HOME} && chmod 0700 ${HOME}
if [ ! -w ${HOME} ]; then echo "${HOME} is not writable"; exit 1; fi

# Prepare custom folder
mkdir -p ${GITBUNDLE_CUSTOM} && chmod 0500 ${GITBUNDLE_CUSTOM}

# Prepare temp folder
mkdir -p ${GITBUNDLE_TEMP} && chmod 0700 ${GITBUNDLE_TEMP}
if [ ! -w ${GITBUNDLE_TEMP} ]; then echo "${GITBUNDLE_TEMP} is not writable"; exit 1; fi

#Prepare config file
if [ ! -f ${GITBUNDLE_APP_INI} ]; then

    #Prepare config file folder
    GITBUNDLE_APP_INI_DIR=$(dirname ${GITBUNDLE_APP_INI})
    mkdir -p ${GITBUNDLE_APP_INI_DIR} && chmod 0700 ${GITBUNDLE_APP_INI_DIR}
    if [ ! -w ${GITBUNDLE_APP_INI_DIR} ]; then echo "${GITBUNDLE_APP_INI_DIR} is not writable"; exit 1; fi

    # Set INSTALL_LOCK to true only if SECRET_KEY is not empty and
    # INSTALL_LOCK is empty
    if [ -n "$SECRET_KEY" ] && [ -z "$INSTALL_LOCK" ]; then
        INSTALL_LOCK=true
    fi

    # Substitute the environment variables in the template
    APP_NAME=${APP_NAME:-"GitBundle"} \
    RUN_MODE=${RUN_MODE:-"prod"} \
    RUN_USER=${USER:-"git"} \
    SSH_DOMAIN=${SSH_DOMAIN:-"localhost"} \
    HTTP_PORT=${HTTP_PORT:-"4000"} \
    ROOT_URL=${ROOT_URL:-""} \
    DISABLE_SSH=${DISABLE_SSH:-"false"} \
    SSH_PORT=${SSH_PORT:-"2222"} \
    SSH_LISTEN_PORT=${SSH_LISTEN_PORT:-$SSH_PORT} \
    DB_TYPE=${DB_TYPE:-"mysql"} \
    DB_HOST=${DB_HOST:-"localhost:3306"} \
    DB_NAME=${DB_NAME:-"gitbundle"} \
    DB_USER=${DB_USER:-"root"} \
    DB_PASSWD=${DB_PASSWD:-""} \
    INSTALL_LOCK=${INSTALL_LOCK:-"false"} \
    DISABLE_REGISTRATION=${DISABLE_REGISTRATION:-"false"} \
    REQUIRE_SIGNIN_VIEW=${REQUIRE_SIGNIN_VIEW:-"false"} \
    SECRET_KEY=${SECRET_KEY:-""} \
    NSQD_CLUSTER_TCP_ADDR=${NSQD_CLUSTER_TCP_ADDR:-""} \
    REDIS_CONNECTION=${REDIS_CONNECTION:-""} \
    envsubst < /etc/templates/app.ini > ${GITBUNDLE_APP_INI}
fi

# Replace app.ini settings with env variables in the form GITBUNDLE__SECTION_NAME__KEY_NAME
environment-to-ini --config ${GITBUNDLE_APP_INI}
