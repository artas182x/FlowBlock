FROM golang:1.17.6-stretch AS build
WORKDIR /src
COPY . .
RUN go build -o /app/generateWallet .

FROM gcr.io/distroless/base-debian10
WORKDIR /app
COPY --from=build /app/generateWallet /app/generateWallet
ENTRYPOINT ["/app/generateWallet"]
