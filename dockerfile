# BUILD STAGE
FROM golang:1.21.5-alpine3.19 as build

WORKDIR /app

COPY ./ta_pago.db ./app
COPY . /app

RUN apk add --no-cache gcc libc-dev
RUN CGO_ENABLED=1 go build -o bin/ta_pago_bot cmd/main.go

# RELEASE STAGE
FROM alpine:3.19.1

ENV TZ="America/Sao_Paulo"

WORKDIR /

COPY --from=build /lib/ld-musl-x86_64.so.1 /lib/
COPY --from=build /lib/libc.musl-x86_64.so.1 /lib/
COPY --from=build /app/bin/ta_pago_bot /
COPY --from=build /app/ta_pago.db /ta_pago.db
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 4000

ENTRYPOINT ["./ta_pago_bot"]
