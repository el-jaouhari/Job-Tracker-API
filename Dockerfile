FROM golang:1.24.11 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o job-tracker cmd/service/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/job-tracker ./job-tracker
COPY --from=builder /app/db ./db
EXPOSE 8080
CMD ["./job-tracker"]
