FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /sudoku-api ./cmd/api/main.go

# Production image
FROM scratch as run

WORKDIR /

COPY --from=build /sudoku-api /sudoku-api

EXPOSE 8080

CMD ["/sudoku-api"]
