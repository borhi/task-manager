FROM golang:1.14-alpine as build
WORKDIR /src/task-manager
COPY . .
RUN go mod vendor
RUN go build -mod=vendor

FROM alpine:latest
COPY --from=build /src/task-manager/.env .
COPY --from=build /src/task-manager/task-manager .
COPY --from=build /src/task-manager/swaggerui ./swaggerui
CMD ["./task-manager"]
EXPOSE 3000
