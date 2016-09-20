
### Simple Thrift vs gRPC performance test

use a simple "helloworld" prototype to test thrift and gRPC. All servers and clients are implemented by Golang

Test result as follows (milliseconds/10000 calls). The first value is using one client to test servers and the second value is using 20 clients to test concurrently.

|  | Golang | 
| ----- | ----- | 
| **Thrift**   | 1229/922  | 
| **gRPC**     | 3825/2658 | 


### Note

I removed original Scala and Java code because I am not interested with it. 

### Obvious

Install protoc and thrift compiler

```
brew install --devel protobuf
brew install thrift
```

Install golang package
```
go get git.apache.org/thrift.git/lib/go/thrift/...
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

Run server
```
go run src/server/main.go 
```

Run client
```
./clientrunner.sh
```

Output will be available in `temp.txt` file.

### Modification

Change Thrift definition to use struct. In previous test, thrift definition was
```
namespace go greeter

service Greeter {
    string sayHello(1:string name);
}
```

While gRPC
```
syntax = "proto3";

package greeter;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

So I change thrift to use struct too
```
namespace go greeter

struct HelloRequest {
  1: string message,
}

struct HelloReply {
  1: string message,
}

service Greeter {
    HelloReply sayHello(1:HelloRequest request);
}
```

### Conclusion

gRPC still have worst speed then thrift even I tried to level the field. Although the ratio is not as bad as previous test. About 3 times slower (not 5 times slower as in previous test for one client test)

gRPC CPU usage also higher then thrift (I just glance at %CPU on my Mac, no recorded data)

