FROM node:6.10-alpine

# Install dependencies & clean up
RUN echo "@testing http://dl-4.alpinelinux.org/alpine/edge/testing" | tee -a /etc/apk/repositories \
    && echo "http://dl-5.alpinelinux.org/alpine/edge/main/" | tee -a /etc/apk/repositories \
    && apk --no-cache --update add \
      curl \
      ca-certificates \
      mono@testing

RUN mkdir -p /usr/lib/nuget \
    && cert-sync /etc/ssl/certs/ca-certificates.crt \
    && curl -#SL https://dist.nuget.org/win-x86-commandline/v3.5.0/NuGet.exe -o /usr/lib/nuget/NuGet.exe

WORKDIR /node

COPY . .

ENTRYPOINT ["node", "index.js"]
