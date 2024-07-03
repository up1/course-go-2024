FROM golang:1.22.5 AS step1
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /api/main ./cmd/main.go

FROM debian:bookworm
WORKDIR /app
COPY --from=step1 /api/main /app/main
EXPOSE 8080
CMD [ "./main" ]
