# mental_wellness_bot

Demo

https://github.com/user-attachments/assets/f3fc2c76-8da8-4283-b86c-f9c5f1c03d7f

Project structure:

```
mentalwellness/
│
├── proto/
│   └── mental_wellness.proto      # gRPC definitions
│
├── static/
│   └── index.html                 # Simple frontend UI
│
├── .env                           # Environment variables
├── go.mod / go.sum                # Go dependencies
├── server.go                      # gRPC + Gemini backend
```

Architecture:

* Go + gRPC backend
* Google Generative AI (Gemini)
* REST proxy for frontend
* Simple web UI
* .env config for API key and port

Steps:

1. Clone or Create Project

```
mkdir mental_wellness_bot && cd mental_wellness_bot
go mod init github.com/ramyasreetejo/mental_wellness_bot
```


2. Install Required Packages

```
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/joho/godotenv
go get github.com/google/generative-ai-go/genai
go get google.golang.org/api/option
```


3. Generate Go gRPC files after creating proto file, defining gRPC Service:

To work with .proto files in a gRPC project on Mac, you'll need two main tools:

- protoc – the Protocol Buffers compiler
- protoc-gen-go – the Go plugin for protoc to generate Go code from .proto files

steps to install(mac specific):

--step 1--
```
brew install protobuf
protoc --version
```

--step 2--
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

--step 3--
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

After all dependent tools installed, run:

```
protoc --go_out=. --go-grpc_out=. proto/mental_wellness.proto
```


4. create server and index html files


5. Update grpc and rest ports, api key in .env file (api key from https://aistudio.google.com/app/apikey)


6. Run Your Bot

```
go run server.go
```

- gRPC running on localhost:50051
- UI and REST API on http://localhost:8080
