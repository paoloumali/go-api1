# api1

## GO prep

- install it local to your user

    ```shell
    wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    tar -C $HOME -xzf go1.21.0.linux-amd64.tar.gz
    ```

- add to path

    ```shell
    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
    ```

## Project setup

- one-time
    - go mod init paoloumali/go-api1
    - go get github.com/gin-gonic/gin
    - go mod tidy
    - go run .
