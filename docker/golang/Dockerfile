FROM ubuntu:22.04

ENV DEBIAN_FRONTEND=noninteractive

ARG ID_USER=1000
ARG ID_GROUP=1000
ARG GOLANG_VERSION=1.20.4
RUN \
    echo "creating group and user" &&\ 
    groupadd -g ${ID_GROUP} app &&\
    useradd -u ${ID_USER} -m app -g app &&\
    usermod -a -G www-data app &&\
    \
    echo "run apt-get update for installing packages" &&\
    apt-get update &&\
    \
    echo "install build-essential wich includes make utitity to run commands from Makefile" &&\
    apt-get install -y build-essential &&\
    \
    echo "installing git for displaying changes for commit in vscode" &&\
    apt-get install -y git &&\
    \
    echo "installing curl to download programs" &&\
    apt-get install curl -y &&\
    \
    echo "installing golang" &&\
    curl -OL https://go.dev/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz &&\
    tar -xvzf go${GOLANG_VERSION}.linux-amd64.tar.gz &&\
    mv go /usr/local/go &&\
    \
    echo "installing vscode server" &&\
    curl -fsSL https://code-server.dev/install.sh | sh

# export go path to use go
ENV PATH /usr/local/go/bin:${PATH}

# installing golang extensions for vscode
USER app
RUN go install -v golang.org/x/tools/gopls@latest &&\
    go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949 &&\
    go install -v github.com/go-delve/delve/cmd/dlv@latest
