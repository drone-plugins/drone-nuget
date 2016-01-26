# Docker image for the Drone NuGet plugin

FROM alpine:3.3

RUN echo "@testing http://dl-4.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    apk add 'nodejs>4.2.4' && \
    apk add mono@testing && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /usr/lib/nuget && \
    wget https://dist.nuget.org/win-x86-commandline/v2.8.6/nuget.exe -O /usr/lib/nuget/NuGet.exe

WORKDIR /bin

COPY package.json /bin/
RUN npm install
COPY . /bin/

ENTRYPOINT [ "node", "index.js" ]
