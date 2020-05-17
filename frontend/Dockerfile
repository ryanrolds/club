FROM node:12-alpine

RUN apk add --update make git
RUN apk --no-cache add ca-certificates

COPY . /frontend/
WORKDIR /frontend/

RUN make build
