FROM golang

RUN apt-get update
RUN apt-get install -y g++

WORKDIR /go/src
COPY . github.com/maxnilz/go-tiff2pdf

WORKDIR /go/src/github.com/maxnilz/go-tiff2pdf
RUN make
RUN go install ./tiff2pdf-service
RUN go install ./tiff2pdf-cli

EXPOSE 9090

ENTRYPOINT ["/go/bin/tiff2pdf-service"]
