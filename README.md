
## UPDATE. Test with database

Well, this is interesting and kinda expected. I added new test with the server connect and do some query to database. 

First, you need to setup Couchbase (see https://github.com/arifsetiawan/gosimple#couchbase) and do
```
go get github.com/couchbase/gocb
``` 

Run server
```
go run src/serverdb/main.go 
```

Run client
```
./clientrunnerdb.sh
```

Test result as follows (milliseconds/100 calls) - 10000 would be too long. The first value is using one client to test servers and the second value is using 20 clients to test concurrently.

|  | Golang | 
| ----- | ----- | 
| **Thrift**   | 3036/3104  | 
| **gRPC**     | 2926/3246 | 


### Simple Thrift vs gRPC performance test

Use a simple "helloworld" prototype to test thrift and gRPC. All servers and clients are implemented by Golang

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

Without any database work. gRPC is about 3 times slower.

With database work, the speed is about the same. 

Observation: Raw speed is small factor that make up whole API speed. Database connection seems has bigger weight then raw speed in determining overall speed. Client to server network condition maybe another factor.
