FROM node:6.10-alpine

# Environment variables
ENV NUGET_VERSION 4.0.0

# Install dependencies & clean up
RUN echo "@testing http://dl-4.alpinelinux.org/alpine/edge/testing" | tee -a /etc/apk/repositories \
    && apk --no-cache --update add \
      curl \
      ca-certificates \
      mono@testing

RUN mkdir -p /usr/lib/nuget \
    && cert-sync /etc/ssl/certs/ca-certificates.crt \
    && curl -#SL https://dist.nuget.org/win-x86-commandline/v$NUGET_VERSION/NuGet.exe -o /usr/lib/nuget/NuGet.exe

WORKDIR /node

COPY . .

ENTRYPOINT ["node", "index.js"]
