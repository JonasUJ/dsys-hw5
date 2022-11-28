FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN make discovery

FROM debian:buster-slim

COPY --from=builder /app/bin/discovery /app/discovery

CMD ["/app/discovery"]
