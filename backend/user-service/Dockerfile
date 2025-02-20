FROM golang:1.24

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy the service code
COPY . .

# Generate swagger docs
RUN swag init

# Build with verbose output
RUN CGO_ENABLED=0 go build -v -o user-service && \
    ls -l user-service && \
    chmod +x user-service

EXPOSE 8082

CMD ["./user-service"]