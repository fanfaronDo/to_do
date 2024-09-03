FROM golang:1.22

ENV GOPATH=/
WORKDIR ./app
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh
RUN go mod download && go build -o todo-app ./cmd/app/main.go
CMD ["./todo-app"]