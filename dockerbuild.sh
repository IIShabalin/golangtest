#!/bin/sh
if [ -d dockerbuild/ ];
then
    rm -R dockerbuild
    mkdir dockerbuild
fi
cd dockerbuild
cp ../Dockerfile .
cp ${GOPATH}/bin/test .
docker build -t crc32/golangtest .
docker push crc32/golangtest
