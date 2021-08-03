Depois de instalar o GO:
caso usar o ZSH no MAC:
export GOPATH=$HOME/go
PATH=$PATH:$GOPATH/bin

1 - go mod init github.com/codeedu/fc2-grpc
2 - go get google.golang.org/protobuf/cmd/protoc-gen-go
3 - go install google.golang.org/protobuf/cmd/protoc-gen-go
4 - protoc --proto_path=proto proto/*.proto --go_out=pb
5 - go run cmd/server/server.go

Client gRCP(Evans):
  Evans auxilia nos testes de métodos e tudo mais sem precisar de um client
  https://github.com/ktr0731/evans#macos

Executar Evans:
  evans -r repl --host localhost --port 50051

Executando métodos:
service UserService

call AddUser

Client gRPC Desenvolvido:
go run cmd/client/client.go
