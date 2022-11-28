FROM golang:1.19-buster as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN make frontend

FROM debian:buster-slim

COPY --from=builder /app/bin/frontend /app/frontend

CMD ["/app/frontend"]
