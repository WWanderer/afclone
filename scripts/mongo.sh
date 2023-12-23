#!/bin/bash

docker run -d \
    -p 27017:27017 \
    --name mongo_afclone \
    -v mongo-data:/c/Users/Matthias/Documents/Code/afclone/resources/data \
    -e MONGODB_INITDB_ROOT_USERNAME=admin \
    -e MONGODB_INITDB_ROOT_PASSWORD=password \
    mongo:latest