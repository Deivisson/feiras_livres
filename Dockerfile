FROM golang:1.16.3

ENV APP_PATH /code/
WORKDIR $APP_PATH

RUN export GOROOT=/usr/lib/go \
    export GOPATH=$HOME/go \
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# install build essentials
RUN apt-get update && \
    apt-get install -y wget build-essential pkg-config --no-install-recommends

COPY . $APP_PATH

RUN go get ./...
RUN go get github.com/codegangsta/gin
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs
RUN go get github.com/ramya-rao-a/go-outline
RUN go get github.com/cweill/gotests/gotests
RUN go get github.com/fatih/gomodifytags
RUN go get github.com/josharian/impl
RUN go get github.com/haya14busa/goplay/cmd/goplay
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get github.com/go-delve/delve/cmd/dlv@master
RUN go get honnef.co/go/tools/cmd/staticcheck@2021.1.2
RUN go get golang.org/x/tools/gopls



