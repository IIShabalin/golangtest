#!/bin/sh
docker run -p 8080:8080 -e TEST_PORT=8080 golangtest