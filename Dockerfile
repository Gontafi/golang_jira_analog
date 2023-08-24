FROM golang:1.21-alpine AS build

RUN go version

ENV GOPATH=/

COPY ./ ./

#RUN apk update && apk add redis  &&\
#    apk add postgresql && \
#    chmod +x wait-for-db.sh

WORKDIR ./app

RUN go mod download
RUN go build -o my-jira ./cmd/main.go

CMD ["./my-jira"]

