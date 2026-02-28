FROM golang:1.25.7

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /app/main ./cmd

EXPOSE 3000
CMD ["./main"]