# Simplified REST GRPC Server
## By Chris Cates

The intention of this repository is to demonstrate an easy way to make REST GRPC servers. It's intentionally unidiomatic to focus on ease of use and user experience. Of course, if there are any issues you find in this repository. Ping me or email me at hello@chriscates.ca. Happy coding!

### Requirements

1. Go 1.6+
2. The grpc libraries installed on your computer.

```
# You can just copy and paste this shell script and it should compile and run no problem.
# If you want to compile more then once you only need to run the last command in this script.

# Run this if you don't have these libraries already.
go get google.golang.org/grpc
go get github.com/gogo/protobuf/protoc-gen-gofast
go get github.com/golang/protobuf/{proto,protoc-gen-go}

# Also make sure the $GOPATH and $GOBIN is in your $PATH
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOBIN

# Compile GRPC methods
sh build.sh

# Run GRPC Server and REST Server
sh server.sh

# Test with a cURL command to the HTTP server on PORT 8080? :)
# Also keep in mind both .sh scripts set your GRPC_ROOT to the $PWD variable. So make sure to run these in the root directory of the repo.
```

MIT Licensed
