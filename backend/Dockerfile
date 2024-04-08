FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/auth-server

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/auth-server .

CMD ["./auth-server"]