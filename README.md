### Modification

Modified to work with Thrift 0.9.3, the latest gRPC 1.0 (Golang), and the latest Cap'nProto.

### Simple Thrift vs gRPC vs Cap'n Proto performance test

use a simple "helloworld" prototype to test thrift, gRPC and Cap'nProto.
All servers and clients are implemented by Golang

Test result as follows (milliseconds/10000 calls). The first value is using one client to test servers and the second value is using 20 clients to test concurrently.

|  | Golang | 
| ----- | ----- | 
| **Thrift**   | 261/169  | 
| **gRPC**     | 1456/813 | 
| **Cap'n Proto**   | 1828/1143 | 


