# syntax=docker/dockerfile:1
FROM golang:1.18 AS builder
WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app


FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/app /
CMD ["/app"]
