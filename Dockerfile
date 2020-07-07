FROM golang:1.14 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o accuweather_exporter accuweather_exporter.go


FROM gcr.io/distroless/static

WORKDIR /root/
COPY --from=builder /app/accuweather_exporter .

ENTRYPOINT ["./accuweather_exporter"]