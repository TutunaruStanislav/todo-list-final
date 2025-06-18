FROM golang:1.24.3-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todo-list-app

FROM alpine:3.14
WORKDIR /app
COPY --from=build /todo-list-app .
COPY web ./web
COPY .env /app/
VOLUME /data

ARG TODO_PORT
EXPOSE $TODO_PORT

CMD ["/app/todo-list-app"]