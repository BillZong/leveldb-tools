package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {

	printUsage := func() {
		fmt.Printf("Usage: %s db_folder_path\n", os.Args[0])
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

	if len(os.Args) == 1 {
		fmt.Println("Level/Snappy DB folder path is not supplied")
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

	db, err := leveldb.OpenFile(dbPath, &opt.Options{
		ReadOnly: true,
	})
	if err != nil {
		fmt.Printf("Could not open DB from:%s, err: %v\n", dbPath, err)
		printUsage()
		return
	}

	defer db.Close()

	iter := db.NewIterator(nil /* slice range, default get all */, nil /* default read options */)
	for iter.Next() {
		k := iter.Key()
		key := hex.EncodeToString(k)
		fmt.Println(key)
	}

	iter.Release()

	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}
