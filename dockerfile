FROM arm32v6/golang:1.16-alpine AS build
RUN apk update
RUN apk upgrade
RUN apk add --update gcc g++

WORKDIR /tmp/app/
COPY ./app/ /tmp/app/
RUN go mod download
RUN CGO_ENABLED=1 go build -o app

FROM arm32v6/alpine
WORKDIR /app/
COPY --from=build /tmp/app/app .

CMD ["./app"]
