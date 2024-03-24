FROM golang:1.22.1-bookworm as builder

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger

CMD ["go", "run", "cmd/api/main.go"]