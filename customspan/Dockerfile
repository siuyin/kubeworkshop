FROM golang:1.26 AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/app

FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /
COPY --from=builder /build/app /app
ENV PORT=8080
EXPOSE ${PORT}
ENTRYPOINT ["/app"]

