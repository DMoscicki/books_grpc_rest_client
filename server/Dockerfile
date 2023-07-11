FROM golang:1.20.5-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /server


FROM alpine
WORKDIR /app
COPY .env /app/.env
COPY --from=build ./server /app/server
EXPOSE 8080/tcp
CMD ["/app/server"]

