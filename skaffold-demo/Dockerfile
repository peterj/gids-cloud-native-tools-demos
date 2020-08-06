FROM golang:1.14.4-alpine3.12 as builder
COPY main.go .
RUN go build -o /app main.go

FROM alpine:3.12
CMD ["./app"]
COPY --from=builder /app .