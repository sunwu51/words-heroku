FROM node:16-alpine3.14
ADD . /app
WORKDIR /app
RUN apk update && apk add git && npm i -g pm2 && npm i
RUN git clone https://github.com/sunwu51/words-db.git
CMD pm2 start npm --name "db" -- run "db" && node index.js