FROM golang:1.20

WORKDIR /go/src/github.com/patrickodacre/go-dev

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# todo:: can this be done better so as to reduce the size of this container?
COPY . .

