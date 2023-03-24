FROM ubuntu:22.04

ENV repo=https://github.com/ryanbourdais/gradle-test.git

RUN apt-get update && apt-get install -y \
	openjdk-17-jdk
# RUN sed -i -e 's/us.archive.ubuntu.com/archive.ubuntu.com/g' /etc/apt/sources.list
RUN apt-get update
RUN rm /bin/sh && ln -s /bin/bash /bin/sh
RUN apt-get -qq -y install curl zip unzip

ENV GRADLE_VERSION=8.0.1

RUN curl -s https://get.sdkman.io | bash
RUN chmod a+x "$HOME/.sdkman/bin/sdkman-init.sh"
RUN source "$HOME/.sdkman/bin/sdkman-init.sh" \
	sdk version \
	sdk install gradle ${GRADLE_VERSION} -y

COPY ./code .

RUN apt-get install git -y