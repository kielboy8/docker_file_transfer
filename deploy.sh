#!/bin/bash
set -v

docker-compose build
docker-compose up -d
docker-compose logs

# For running the app manually
# export RECEIVER=$(docker ps --format '{{.Names}}' | grep transfer_receiver-app)
# export SENDER=$(docker ps --format '{{.Names}}' | grep transfer_sender-app)