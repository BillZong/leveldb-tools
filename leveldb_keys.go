package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
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

	db, err := leveldb.OpenFile(dbPath, nil)
	defer db.Close()

	if err != nil {
		fmt.Println("Could not open DB from:", dbPath)
		printUsage()
		return
	}

	iter := db.NewIterator(nil /* slice range, default get all */, nil /* default read options */)
	for iter.Next() {
		key := iter.Key()
		keyName := string(key[:])
		fmt.Println(keyName)
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}
