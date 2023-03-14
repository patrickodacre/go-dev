FROM golang:1.20

WORKDIR /go/src/github.com/patrickodacre/go-2

COPY . .

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -polling=true -build="go build -o main ./cmd/main.go" -command="./main"