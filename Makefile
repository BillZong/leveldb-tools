.PHONY: bin
bin: bindir del jsonval key keys kvs predel

bindir:
	mkdir -p bin/

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

.PHONY: linux-bin
linux-bin: linux-bindir linux-del linux-jsonval linux-key linux-keys linux-kvs linux-predel

linux-bindir:
	mkdir -p bin/linux

linux-del:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_del leveldb_del.go

linux-jsonval:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_jsonval leveldb_jsonval.go

linux-key:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_key leveldb_key.go

linux-keys:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_keys leveldb_keys.go

linux-kvs:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_kvs leveldb_kvs.go

linux-predel:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/leveldb_predel leveldb_predel.go
