go install github.com/golang/protobuf/protoc-gen-go@latest 
mv ~/go/bin/protoc-gen-go ~/go/bin/protoc-gen-githubgo
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest 
mv ~/go/bin/protoc-gen-go ~/go/bin/protoc-gen-googlego
go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
go install github.com/gogo/protobuf/protoc-gen-gogofast@latest  
go install github.com/gogo/protobuf/protoc-gen-gogo@latest    
go install github.com/aggronmagi/protoc-gen-gopb/@latest