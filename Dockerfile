# FROM golang:1.15.2-alpine3.12
FROM golang:alpine as builder

LABEL maintainer="Sammarth Kapse <sammarth.kapse@razorpay.com>"

RUN apk update && apk add --no-cache git

WORKDIR /UptimeMonitoringService

COPY . .

RUN go mod download

# RUN go build .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ums .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /UptimeMonitoringService/ums .
COPY --from=builder /UptimeMonitoringService/.env .

EXPOSE 8080

CMD ["./ums"]