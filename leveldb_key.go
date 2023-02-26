package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func hexOrRaw(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		raw := strings.TrimPrefix(str, "0x")
		b, _ := hex.DecodeString(raw)
		return b
	} else if strings.HasPrefix(str, "0X") {
		raw := strings.TrimPrefix(str, "0X")
		b, _ := hex.DecodeString(raw)
		return b
	}

	return []byte(str)
}

func main() {

	printUsage := func() {
		fmt.Printf("Usage: %s db_folder_path key_to_read\n", os.Args[0])
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

	if len(os.Args) < 3 {
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

	key := hexOrRaw(os.Args[2])

	val, err := db.Get(key, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hex.EncodeToString(val))
	}
}
