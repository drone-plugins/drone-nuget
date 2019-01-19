FROM alpine:3.4

RUN echo "@testing http://dl-4.alpinelinux.org/alpine/edge/testing" | tee -a /etc/apk/repositories
RUN echo "http://dl-5.alpinelinux.org/alpine/edge/main/" | tee -a /etc/apk/repositories

RUN apk update && \
  apk --no-cache add wget ca-certificates nodejs mono@testing && \
  cert-sync /etc/ssl/certs/ca-certificates.crt

RUN mkdir -p /usr/lib/nuget && \
  wget \
    https://dist.nuget.org/win-x86-commandline/v3.5.0/NuGet.exe \
    -O /usr/lib/nuget/NuGet.exe

WORKDIR /node
ADD package.json /node/
ADD index.js /node/
RUN npm install --production

ENTRYPOINT ["node", "index.js"]