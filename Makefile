.PHONY: bin
bin: del jsonval key keys kvs predel

del:
	go build -o bin/leveldb_del leveldb_del.go

jsonval:
	go build -o bin/leveldb_jsonval leveldb_jsonval.go

key:
	go build -o bin/leveldb_key leveldb_key.go

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
