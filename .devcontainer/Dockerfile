FROM mcr.microsoft.com/vscode/devcontainers/go:latest

RUN apt-get update -y

#RUN apt-get install python3 python3-dev python3-pip -y
#RUN pip3 install pre-commit
RUN apt-get install -y protobuf-compiler

RUN go install mvdan.cc/gofumpt@latest
RUN go install github.com/boumenot/gocover-cobertura@latest
RUN go install golang.org/x/tools/cmd/goimports@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install github.com/dkorunic/betteralign/cmd/betteralign@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
