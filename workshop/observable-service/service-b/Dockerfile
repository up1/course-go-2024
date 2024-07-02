FROM golang:1.22.4-alpine3.20 as step1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

FROM alpine:3.20
WORKDIR /app
COPY --from=step1 /app/api .
EXPOSE 8080
CMD ["/app/api"]