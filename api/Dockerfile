FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /sudoku-api ./cmd/api/main.go

EXPOSE 8080

CMD ["/sudoku-api"]