#!/bin/sh
DOCKER_REGISTRY_SERVER=https://index.docker.io/v1/
DOCKER_USER=crc32
DOCKER_EMAIL=ivan.shabalin@gmail.com
DOCKER_PASSWORD='DCpwd615(@'

kubectl create secret docker-registry golangtest \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL