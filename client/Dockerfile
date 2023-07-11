FROM golang:1.20.5-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o /client


FROM alpine
WORKDIR /app
COPY .env /app/.env
COPY --from=build ./client /app/client
EXPOSE 8081/tcp
CMD ["/app/client"]