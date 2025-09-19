FROM golang:1.22-alpine AS build
WORKDIR /app
COPY services/go/serving ./services/go/serving
WORKDIR /app/services/go/serving
RUN go build -o /serving
FROM alpine:3.20
COPY --from=build /serving /serving
EXPOSE 8080
ENTRYPOINT ["/serving"]
