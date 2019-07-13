FROM golang:latest AS builder
ADD . /app/ingresso-watcher
WORKDIR /app/ingresso-watcher
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /ingresso-watcher cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /ingresso-watcher ./
RUN chmod +x ./ingresso-watcher
ENTRYPOINT ["./ingresso-watcher"]
