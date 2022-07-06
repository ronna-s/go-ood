FROM golang:1.18.3-alpine
ENV CGO_ENABLED 0
RUN go install github.com/vektra/mockery/v2@latest
RUN go install golang.org/x/tools/cmd/godoc@latest
RUN go install golang.org/x/lint/golint@latest
RUN go install golang.org/x/tools/cmd/stringer@latest
RUN go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
WORKDIR /root
COPY go.mod /root
COPY go.sum /root
RUN go mod download
COPY . /root

# Add and compile the packages
RUN go install /root/cmd/maze
RUN go install /root/cmd/habitat
CMD maze
