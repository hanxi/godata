go/example/example.pb.go: example.proto
	mkdir -p go/example # make directory for go package
	protoc $$PROTO_PATH --go_opt=paths=source_relative --go_out=go/example example.proto
