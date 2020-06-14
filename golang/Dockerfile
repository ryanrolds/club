FROM golang:1.14-alpine as build

RUN apk add --update make git build-base
RUN apk --no-cache add ca-certificates

COPY . /golang/
WORKDIR /golang/

RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=build /golang/backend .
COPY --from=build /golang/backend.yaml .
COPY --from=build /golang/static static
