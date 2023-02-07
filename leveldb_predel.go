package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func main() {
	printUsage := func() {
		fmt.Printf("Usage: %s db_folder_path key_prefix_to_delete key_length_without_prefix\n", os.Args[0])
	}

	fileExists := func(path string) (bool, error) {
		_, err := os.Stat(path)
		if err == nil {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}

		return true, err
	}

	if len(os.Args) < 4 {
		printUsage()

		return
	}

	dbPath := os.Args[1]

	dbPresent, err := fileExists(dbPath)
	if !dbPresent {
		fmt.Printf("The DB path: %s does not exist.\n", dbPath)
		printUsage()

		return
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	defer db.Close()

	if err != nil {
		fmt.Println("Could not open DB from:", dbPath, ", err:", err)
		printUsage()

		return
	}

	prefix := os.Args[2]
	keylen, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Could not get key(without prefix) length:", err)
		printUsage()

		return
	}

	iter := db.NewIterator(
		bytesPrefixRange([]byte(prefix), make([]byte, keylen)), // from zero
		nil,
	)

	for iter.Next() {
		k := iter.Key()
		if len(k) != len(prefix)+keylen {
			fmt.Println("skip key:", hex.EncodeToString(k))

			continue
		}

		err := db.Delete(iter.Key(), &opt.WriteOptions{
			NoWriteMerge: true,
			// Sync: false,
		})
		if err != nil {
			fmt.Println(err)

			break
		}

		fmt.Printf("delete key %s\n", hex.EncodeToString(k))
	}

	iter.Release()

	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}

// bytesPrefixRange returns key range that satisfy
// - the given prefix, and
// - the given seek position
func bytesPrefixRange(prefix, start []byte) *util.Range {
	r := util.BytesPrefix(prefix)
	r.Start = append(r.Start, start...)

	return r
}
