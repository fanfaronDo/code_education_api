FROM golang:1.22

RUN go version


ENV GOPATH=/
WORKDIR ./app
COPY ./ ./
RUN chmod +x wait-for-postgres.sh

RUN go mod download && go build -o note-app ./cmd/app/main.go
ENTRYPOINT ["./note-app"]
