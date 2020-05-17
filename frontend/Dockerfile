FROM node:12.16-alpine

#install yarn because its great
RUN npm i -g yarn

# create app directory in the container
RUN mkdir -p /usr/src/club
WORKDIR /usr/src/club

# copy everything except for stuff in .dockerignore
COPY . .

# install and run tests
RUN npm i && npm run test && npm run build

# remove frontend folder its useless now
RUN rm -rf frontend

# expose the express port
EXPOSE 3000

CMD ["npm", "start"]
