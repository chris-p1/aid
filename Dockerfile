FROM golang:1.24.5-alpine AS gobuild

WORKDIR /app

COPY ./ ./

# Install dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o main cmd/api/main.go

FROM alpine:latest

# Copy the executable to a stripped down alpine
COPY --from=gobuild /app/main /usr/local/bin/main
WORKDIR /app

EXPOSE 8080
CMD ["main"]
