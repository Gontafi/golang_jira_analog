FROM golang:1.21-alpine AS build


RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apk update
RUN apk add postgresql-client


# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

WORKDIR ./app
# build go app
RUN go mod download
RUN go build -o my-jira ./cmd/myjira/main.go

CMD ["./my-jira"]

