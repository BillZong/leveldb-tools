.PHONY: bin
bin: jsonval keys kvs predel

jsonval:
	go build -o bin/leveldb_jsonval leveldb_jsonval.go

keys:
	go build -o bin/leveldb_keys leveldb_keys.go

kvs:
	go build -o bin/leveldb_kvs leveldb_kvs.go

predel:
	go build -o bin/leveldb_predel leveldb_predel.go

.PHONY: install
install: bin
	# make sure it is in binary
	$(eval GOPATH = $(shell go env GOPATH))
	mkdir -p $(GOPATH)/bin
	cp -R bin/ $(GOPATH)/bin/
