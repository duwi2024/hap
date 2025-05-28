//go:build ignore
// +build ignore

// Creates swift code for all HomeKit service and characteristic types and logs
// that to the console.

package main

import (
	"encoding/json"
	"github.com/LUJUNQUAN/hap/gen"
	"github.com/LUJUNQUAN/hap/gen/swift"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var LibPath = os.ExpandEnv("$GOPATH/src/github.com/LUJUNQUAN/hap")
var GenPath = filepath.Join(LibPath, "gen")
var MetadataPath = filepath.Join(GenPath, "metadata.json")

func main() {

	log.Println("Import data from", MetadataPath)

	// Open metadata file
	f, err := os.Open(MetadataPath)
	if err != nil {
		log.Fatal(err)
	}

	// Read content
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Import json
	metadata := gen.Metadata{}
	err = json.Unmarshal(b, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	if b, err := swift.CharacteristicEnumDecl(metadata); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}

	if b, err := swift.ServiceEnumDecl(metadata); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(b))
	}
}
