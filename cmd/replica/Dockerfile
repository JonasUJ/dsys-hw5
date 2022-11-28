FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN make replica

FROM debian:buster-slim

COPY --from=builder /app/bin/replica /app/replica

CMD ["/app/replica"]
