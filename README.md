# gocode

Some libraries in GO

### hello

```text
$ go run cmd/hello/main.go
HelloService.Say executed with args: Bob
Hello Bob
```

### hello-server

```text
$ go run cmd/hello-server/main.go 
Starting server...
Listen on port 9000
```

```text
$ curl  -H "Content-Type: application/json"  -d '{"jsonrpc": "2.0", "method": "HelloService.Say", "params": [{"Name":"Bob"}], "id": 1}' http://localhost:9000/rpc
{"result":{"Message":"Hello Bob"},"error":null,"id":1}
```

```text
HelloService.Say executed with args: Bob
```
