from golang:alpine
RUN apk add git
RUN go get -u github.com/basraven/ceryx && go build github.com/basraven/ceryx
CMD ceryx