FROM --platform=linux/x86_64 golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN GOARCH=amd64 go build -o app

FROM  --platform=linux/x86_64 alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
