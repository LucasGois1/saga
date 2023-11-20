FROM arm64v8/golang
LABEL authors="lucasgois"

RUN go install golang.org/x/tools/gopls@latest && \
        go install github.com/go-delve/delve/cmd/dlv@latest && \
        go install honnef.co/go/tools/cmd/staticcheck@latest

WORKDIR /app

ENTRYPOINT [ "tail", "-f", "/dev/null"]