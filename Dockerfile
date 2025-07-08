FROM golang:1.24.4-alpine3.21
WORKDIR /app
COPY  go.mod go.sum ./
RUN go mod download
COPY . .
RUN ls cmd/app
RUN CGO_ENABLED=0 GOOS=linux go build -o bin cmd/app/main.go
EXPOSE 8080
CMD ["/app/bin"]
