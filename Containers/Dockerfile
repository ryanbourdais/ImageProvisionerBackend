FROM ubuntu:22.04

ENV repo=https://github.com/ryanbourdais/ImageProvisionerBackend.git

RUN apt-get update
RUN apt-get -qq -y install curl zip unzip

RUN apt install golang -y
RUN go version
RUN apt-get install git -y

RUN apt install docker -y

RUN git clone ${repo}