FROM golang:1.19-alpine
RUN go install golang.org/x/tools/cmd/goimports@latest


FROM node:12-alpine

COPY . /home/yeoman/tokend-module
COPY --from=0 /go/bin/goimports /usr/local/bin/
WORKDIR /home/yeoman/tokend-module

RUN true \
    && apk add --no-cache git sudo \
    && npm config set unsafe-perm true \
    && npm install --global --silent yo@4.3.1 \
    && npm link \
    && mkdir /generated \
    && chmod +x /home/yeoman/tokend-module/entrypoint.sh \
    && adduser -D -u 501 yeoman \
    && echo "yeoman ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
    && chown -R yeoman:yeoman /home/yeoman \
    && true
ENV HOME /home/yeoman

USER yeoman
WORKDIR /generated
ENTRYPOINT /home/yeoman/tokend-module/entrypoint.sh

