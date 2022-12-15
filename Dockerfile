FROM golang:1.19.4-alpine3.17 as build
ADD . /app
WORKDIR /app
RUN go build main.go

FROM node:16-alpine3.14
ADD . /app
WORKDIR /app
COPY --from=build /app/main /app/main
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /app
ENV ZONEINFO=/app/zoneinfo.zip
RUN apk update && apk add git && npm i -g pm2 && npm i
RUN git clone https://github.com/sunwu51/words-db.git
CMD pm2 start npm --name "db" -- run "db" && ./main