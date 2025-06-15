FROM golang:1.24.3-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /todo-list-app

FROM alpine:latest
WORKDIR /app
COPY --from=build /todo-list-app .
COPY web ./web
COPY .env /app/
COPY data ./data

VOLUME /data
ARG TODO_PORT
EXPOSE $TODO_PORT

CMD ["/app/todo-list-app"]