FROM golang:1.20-alpine3.17 as builder

WORKDIR /app
COPY go.mod ./

RUN go mod download

COPY src/main/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /get-balance

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /get-balance ./

ENV token ""
ENV account_id ""

#EXPOSE 8080

CMD ["/get-balance"]