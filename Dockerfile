FROM golang:1.17-alpine

RUN apk update && apk upgrade && \
	apk add --no-cache bash git openssh

LABEL maintainer="Lee Li Wei <leeliwei930@gmail.com>"

WORKDIR /go/src/api
COPY . .

# Use CompileDaemon to rebuild and restart server while code changed
RUN go get github.com/githubnemo/CompileDaemon

RUN go get -d -v
RUN go install -v


EXPOSE 8080


CMD ["sh", "init.sh"]
