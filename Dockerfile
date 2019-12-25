############################
# STEP 1 build executable binary
############################

FROM golang:1.13-alpine3.10 as builder

ARG VERSION
ARG BRANCH
ARG COMMIT
ARG REPO

RUN apk update && apk add --no-cache make gcc musl-dev linux-headers git

COPY ./$REPO $GOPATH/src/github.com/dynastiateam/backend/$REPO/
COPY ./vendor $GOPATH/src/github.com/dynastiateam/backend/$REPO/vendor
WORKDIR $GOPATH/src/github.com/dynastiateam/backend/$REPO/

RUN cd cmd && go build -ldflags="-X main.Version=$VERSION -X main.Branch=$BRANCH -X main.Commit=$COMMIT" -a -o /go/bin/svc

############################
# STEP 2 build a small image
############################

FROM alpine:latest

RUN apk add --no-cache ca-certificates

# Copy our static executable
COPY --from=builder /go/bin/svc /svc/
WORKDIR /svc

RUN chmod +x svc

# Run the svc binary.
CMD ["./svc"]
