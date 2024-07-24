FROM golang:1.22.5-alpine as builder

ARG GO_BUILD_COMMAND="go build -tags static_all"

# Install some build deps + ssh tools for the setup below.
RUN apk update && apk --no-cache add  build-base  git bash  coreutils openssh  openssl


# this command if you get source from bitbucket repos
# Create the directory where the application will reside
RUN mkdir -p /go/src/github.com/feildrixliemdra/go-boilerplate


WORKDIR /go/src/github.com/feildrixliemdra/go-boilerplate

COPY . .


# application builder step
RUN go mod tidy && go mod download && go mod vendor
RUN eval $GO_BUILD_COMMAND


# STEP 2 build a small image
# Set up the final (deployable/runtime) image.
FROM alpine:3.10.2


# setup package dependencies
RUN apk --no-cache update && apk --no-cache  add  ca-certificates bash jq curl

ENV BUILDDIR=/go/src/github.com/feildrixliemdra/go-boilerplate
ENV PROJECT_DIR=/opt/go-boilerplate

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

#create project directory
RUN mkdir -p $PROJECT_DIR/config
RUN #mkdir -p $PROJECT_DIR/database/migration

WORKDIR $PROJECT_DIR

COPY --from=builder $BUILDDIR/go-boilerplate go-boilerplate
COPY --from=builder $BUILDDIR/config/config.yaml $PROJECT_DIR/config/config.yaml
#COPY --from=builder $BUILDDIR/database/migration $PROJECT_DIR/database/migration
#COPY --from=builder $BUILDDIR/config/msg.yaml $PROJECT_DIR/config/msg.yaml

CMD ["sh","-c", "/opt/go-boilerplate/go-boilerplate serve-http"]
