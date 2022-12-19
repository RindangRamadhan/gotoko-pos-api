#!/bin/bash

SERVICE_NAME='gotoko-pos-api'

CompileDaemon \
    -exclude-dir="scripts" \
    -color=true \
    -graceful-kill=true \
    -pattern="^(\.env.+|\.env)|(.+\.go|.+\.c)$" \
    -build="go build -o $SERVICE_NAME ." \
    -command="./$SERVICE_NAME start"