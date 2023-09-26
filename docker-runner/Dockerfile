FROM golang:1.21.1-bullseye

WORKDIR /app

ADD . /app/

RUN go mod tidy

RUN make build

ENTRYPOINT [ "/bin/bash" "make" ,"$mode"]