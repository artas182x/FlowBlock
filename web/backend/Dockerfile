FROM golang:1.17.6-stretch AS build
WORKDIR /src
COPY . .
RUN go build -o /out/backend .

FROM gcr.io/distroless/base-debian10
ENV GIN_MODE=release
WORKDIR /out
EXPOSE 8000
COPY --from=build /out/backend /out/backend
ENTRYPOINT ["/out/backend"]