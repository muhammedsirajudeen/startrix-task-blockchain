FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 3000

CMD ["./app"]
