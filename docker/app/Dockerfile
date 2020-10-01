FROM golang:1.15-alpine as build

WORKDIR /go/berbagi

COPY src/app ./src/app
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git \
  && go get github.com/oxequa/realize

WORKDIR /go/berbagi/src/app

RUN go build -o app

FROM alpine as prod

WORKDIR /app

COPY --from=build /go/berbagi/app .

RUN addgroup go \
  && adduser -D -G go go \
  && chown -R go:go /berbagi/app

CMD ["./app"]