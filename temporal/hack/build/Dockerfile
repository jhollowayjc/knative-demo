FROM golang:1.24 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY --from=build ./framework ./framework
COPY ./ ./

RUN go build -o /worker ./framework

FROM gcr.io/distroless/static-debian12

COPY --from=builder /worker /
CMD ["/worker"]
