FROM golang:1.14-alpine

RUN apk add --update make git
RUN apk --no-cache add ca-certificates

COPY . /club
WORKDIR /club

RUN go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=0 /club/club .
COPY --from=0 /club/static static

CMD ["./club"]
