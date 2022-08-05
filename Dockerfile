FROM golang:1.18.3-alpine3.16 AS build

WORKDIR /base

COPY . .
RUN go build -o bin/app cmd/main.go

FROM scratch AS release

WORKDIR /

COPY --from=build /base/bin/app .

ENTRYPOINT ["/app"]