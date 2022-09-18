FROM golang:1.18.3-alpine
ENV CGO_ENABLED 0
RUN go install github.com/vektra/mockery/v2@v2.14.0
RUN go install golang.org/x/tools/cmd/godoc@v0.1.12
RUN go install golang.org/x/tools/cmd/stringer@v0.1.12
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@v1.6.1
WORKDIR /root
COPY go.mod /root
COPY go.sum /root
RUN go mod download
COPY . /root

# Add and compile the packages
RUN go install /root/cmd/maze
#RUN go install /root/cmd/pnp
#RUN go install /root/cmd/top

CMD maze
