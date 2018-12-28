FROM golang:1.10

WORKDIR /go/src/github.com/dyweb/sundial

COPY . .

ENV GOPATH /go

ARG CMD_DIR=./cmd

ARG ROOT=github.com/dyweb/sundial

ARG VERSION=

ARG COMMIT=

ARG TARGET=sundial

RUN GOOS=linux GOARCH=amd64                      \
	go build -i -v -o /tmp/${TARGET}                  \
	-ldflags "-s -w -X ${ROOT}/pkg/version.Version=${VERSION}  \
	-X ${ROOT}/pkg/version.Commit=${COMMIT}                    \
	-X ${ROOT}/pkg/version.Package=${ROOT}"                    \
	${CMD_DIR}/${TARGET};

FROM debian:jessie

ARG TARGET=sundial

COPY --from=0 /tmp/${TARGET} /${TARGET}

RUN ln -s /${TARGET} /entrypoint

ENTRYPOINT ["/entrypoint"]
