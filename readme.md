# How to generate gRPC codes out of proto file
## Ruby
```
protoc --ruby_out=ruby --grpc_out=ruby --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` xrates.proto
```
