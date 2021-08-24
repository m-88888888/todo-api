FROM golang:1.15-alpine
COPY src/api /go/src/api
WORKDIR /go/src/api
RUN apk update \
  && apk add --no-cache git \
  && go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql
EXPOSE 8080
