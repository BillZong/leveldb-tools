package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
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
		fmt.Printf("Usage: %s db_folder_path key_to_put hex_data_to_put\n", os.Args[0])
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

	key := hexOrRaw(os.Args[2])
	value := hexOrRaw(os.Args[3])

	if err := db.Put(key, value, nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("put key:", os.Args[2], ", value:", os.Args[3])
	}
}
