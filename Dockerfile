FROM golang:1.8.3-alpine

MAINTAINER Kalaiarasan skalaiarasan27@gmail.com

COPY .  /usr/local/go/src/sklrsn.github.com/Robot

RUN apk add --update git

RUN go install sklrsn.github.com/Robot

RUN go build /usr/local/go/src/sklrsn.github.com/Robot/*.go

CMD ["Robot"]
