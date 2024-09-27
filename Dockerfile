## base go image
FROM golang:1.23-alpine as builder
#
RUN mkdir /app
#
COPY . /app
#
WORKDIR /app
#
RUN CGO_ENABLED=0 go build -o dolphindance-api .
#
RUN chmod +x /app/dolphindance-api
#
## build a tiny docker image
FROM alpine:latest
#
RUN mkdir /app
#
COPY --from=builder /app/dolphindance-api /app

CMD [ "/app/dolphindance-api"]