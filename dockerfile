# BUILD STAGE
FROM golang:1.21.5-alpine3.19 as build_stage

WORKDIR /app/ta_pago_bot

COPY . .

RUN go mod download

EXPOSE 4000

RUN go build -o /bin cmd/main.go

# DEPLOY STAGE

FROM alpine:3.19.0

WORKDIR /

COPY --from=build_stage /bin /bin

EXPOSE 4000

USER nonroot:nonroot

ENTRYPOINT [ "/bin" ]