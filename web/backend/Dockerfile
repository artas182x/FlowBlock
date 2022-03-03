FROM golang:1.17.6-stretch AS build
WORKDIR /src
COPY . .
RUN go build -o /out/backend .
ENV GIN_MODE=release
WORKDIR /out
CMD ["/out/backend"]
