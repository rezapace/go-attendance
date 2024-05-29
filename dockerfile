FROM golang:1.20.2-alpine3.17 as builder

WORKDIR /app
COPY . .

RUN apk add build-base

RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o main.app .

FROM alpine:latest

#  Install tzdata for timezone support
RUN apk add --no-cache tzdata

COPY --from=builder /app/main.app .
COPY serviceAccountKey.json /app/serviceAccountKey.json

CMD ["./main.app"]