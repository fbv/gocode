#!/bin/sh

curl  -H "Content-Type: application/json"  -d '{"jsonrpc": "2.0", "method": "HelloService.Say", "params": [{"Name":"Bob"}], "id": 1}' http://localhost:9000/rpc
