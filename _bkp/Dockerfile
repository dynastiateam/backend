############################
# STEP 1 build executable binary
############################

FROM golang:1.12.7-alpine3.10 as builder

RUN apk update && apk add --no-cache make gcc musl-dev linux-headers git

ENV GO111MODULE=on

COPY .. $GOPATH/src/github.com/dynastiateam/backend/
WORKDIR $GOPATH/src/github.com/dynastiateam/backend/


RUN cd cmd && go build -mod=vendor -a -o /go/bin/svc

############################
# STEP 2 build a small image
############################

FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Copy our static executable
COPY --from=builder /go/bin/svc /svc/
WORKDIR /svc

# Port on which the services will be exposed.
EXPOSE 8080
EXPOSE 8888

RUN chmod +x svc

# Run the svc binary.
CMD ["./svc"]
