FROM golang:1.21.1-alpine as build

RUN mkdir /var/app
WORKDIR /var/app
COPY . /var/app/

RUN go mod download

RUN go build -o /main ./cmd/api/main.go

FROM alpine:latest

# We need to copy the binary from the build image to the production image.
COPY --from=build /main .
EXPOSE 8080
CMD [ "/main" ]
