FROM golang:alpine
ADD src/qntfy /go/src/main
WORKDIR /go/src/main
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get github.com/montanaflynn/stats
CMD ["go", "run", "."]