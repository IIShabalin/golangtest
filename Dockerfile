FROM ubuntu
WORKDIR /opt/test
ADD . /opt/test
ENTRYPOINT ["/opt/test/test"]
