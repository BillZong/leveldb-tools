.PHONY: bin
bin: jsonval keys kvs predel

jsonval:
	go build -o leveldb_jsonval leveldb_jsonval.go

keys:
	go build -o leveldb_keys leveldb_keys.go

kvs:
	go build -o leveldb_kvs leveldb_kvs.go

predel:
	go build -o leveldb_predel leveldb_predel.go

.PHONY: install
install: bin
	# make sure it is in binary
	$(eval GOPATH = $(shell go env GOPATH))
	mkdir -p $(GOPATH)/bin
	cp leveldb_jsonval $(GOPATH)/bin/
	cp leveldb_keys $(GOPATH)/bin/
	cp leveldb_kvs $(GOPATH)/bin/
	cp leveldb_predel $(GOPATH)/bin/