FROM golang:1.27.3-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go test ./... && CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /calculator .

FROM alpine:3.23
RUN addgroup -S calculator && adduser -S -G calculator calculator
COPY --from=build /calculator /usr/local/bin/calculator
USER calculator
EXPOSE 8080
HEALTHCHECK --interval=10s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://127.0.0.1:8080/health || exit 1
ENTRYPOINT ["/usr/local/bin/calculator"]
