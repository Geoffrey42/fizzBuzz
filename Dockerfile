FROM golang:1.15.6-buster

ARG HTTP_PROXY
ARG HTTPS_PROXY

WORKDIR /fizzBuzz

COPY . .

RUN apt install gnupg ca-certificates
RUN apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 379CE192D401AB61
RUN echo "deb https://dl.bintray.com/go-swagger/goswagger-debian ubuntu main" | tee /etc/apt/sources.list.d/goswagger.list
RUN apt update 
RUN apt install swagger

RUN go get -d -v ./...

RUN go install ./cmd/fizzbuzz-server

RUN which fizzbuzz-server

FROM debian:buster-slim

WORKDIR /fizzBuzz

COPY --from=0 /go/bin/fizzbuzz-server .

RUN ls

CMD ./fizzbuzz-server --host=0.0.0.0 --port=5000