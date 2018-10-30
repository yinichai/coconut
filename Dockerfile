ARG path=/go/src/github.com/yinichai/coconut
FROM golang AS builder

ARG path
WORKDIR $path

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep

RUN chmod +x /usr/bin/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -vendor-only

COPY . .
RUN CGO_ENABLED=0 go build

FROM alpine:latest

ARG path
WORKDIR /app

COPY --from=builder $path/coconut ./

EXPOSE 9060

CMD ["./coconut"]