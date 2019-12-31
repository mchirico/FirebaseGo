FROM golang:1.13.5-alpine3.11 AS build

RUN apk add --no-cache git

ARG project
WORKDIR /go/src/$project

# Copy the entire FirebaseGo and build it
# This layer is rebuilt when a file changes in the FirebaseGo directory
COPY . /go/src/$project/
RUN go get -v -t -d ./...
RUN go build -o /bin/$project

# This results in a single layer image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /bin/$project /bin/$project
ENTRYPOINT ["/bin/$project"]
CMD ["--help"]

