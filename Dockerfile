FROM golang:1-alpine AS development

ENV PROJECT_PATH=/solutions-dapp
ENV PATH=$PATH:$PROJECT_PATH/art-admin/bin

RUN apk add --no-cache make git bash alpine-sdk protobuf

COPY --from=bufbuild/buf:latest /usr/local/bin/buf /usr/local/go/bin/
 
ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir -p $PROJECT_PATH
COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN make install build

FROM alpine:latest AS production

WORKDIR /root/
RUN apk --no-cache add ca-certificates

COPY --from=development /solutions-dapp/art-admin/bin/ .
RUN ["chmod", "+x", "./art-admin"]
ENTRYPOINT ["./art-admin"]
