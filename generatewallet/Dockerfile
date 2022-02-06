FROM golang:1.17.6-stretch AS build
WORKDIR /src
COPY . .
RUN go build -o /app/generateWallet .
CMD ["/app/generateWallet"]
