FROM golang:1-alpine AS development

ENV PROJECT_PATH=/art-admin
ENV PATH=$PATH:$PROJECT_PATH/bin

RUN apk add --no-cache make git bash alpine-sdk protobuf

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN make install clean build

FROM alpine:latest AS production

WORKDIR /root/
RUN apk --no-cache add ca-certificates

COPY --from=development /art-admin/bin/ art-admin
RUN ["chmod", "+x", "./art-admin"]
ENTRYPOINT ["./art-admin"]
