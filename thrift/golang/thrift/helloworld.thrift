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