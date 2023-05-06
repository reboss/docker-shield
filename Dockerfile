FROM golang:alpine as build
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go build

FROM golang:alpine
COPY --from=build /go/src/app/docker-shield /bin/docker-shield

RUN mkdir -p /run/docker/plugins

ENTRYPOINT ["/bin/docker-shield"]
