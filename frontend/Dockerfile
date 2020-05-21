FROM node:12-alpine as build

RUN apk add --update make git
RUN apk --no-cache add ca-certificates

COPY . /frontend/
WORKDIR /frontend/

RUN make install
RUN make build

# production environment
FROM nginx:stable-alpine
COPY --from=build /frontend/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
