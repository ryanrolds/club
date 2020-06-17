FROM golang:1.14-alpine as build

RUN apk add --update make git build-base
RUN apk --no-cache add ca-certificates

COPY . /club/
WORKDIR /club/

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /club/
COPY --from=build /club/backend .
COPY --from=build /club/backend.yaml .
COPY --from=build /club/static static
