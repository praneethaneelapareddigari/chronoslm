FROM golang:1.22-alpine AS build
WORKDIR /app
COPY services/go/ingestor ./services/go/ingestor
WORKDIR /app/services/go/ingestor
RUN go build -o /ingestor
FROM alpine:3.20
COPY --from=build /ingestor /ingestor
EXPOSE 7070
ENTRYPOINT ["/ingestor"]
