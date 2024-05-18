FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . . 

RUN go mod tidy

ENV GOARCH=amd64
ENV GOOS=linux 

RUN go build -o main ./src/main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]