FROM golang:1.16.5-alpine3.13 AS build
ENV GOPATH=
WORKDIR /app
COPY . .
RUN go build -o serv-e .

FROM alpine:3.12
COPY --from=build app/serv-e .
COPY --from=build app/request_layout.html .
CMD ["./serv-e"]
