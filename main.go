package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var searchKey = flag.String("k", "", "The json key to search for")

func main() {
	usage := `jkey

Takes stdin and a key and prints out all instances of that key.

Usage: %s
`
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			break
		}
		buffer.WriteString(strings.Trim(line, " "))
	}

	output := buffer.String()
	var raw interface{}
	err := json.Unmarshal([]byte(output), &raw)
	if err != nil {
		fmt.Println("error serializng json")
	}
	jsonData := raw.(map[string]interface{})
	searchMap(jsonData, *searchKey, ".")
	fmt.Println("Done searching")
}

func searchMap(theMap map[string]interface{}, key string, path string) {
	for k, v := range theMap {
		if k == key {
			fmt.Println(fmt.Sprintf("Found the key at %s", path+k))
		}

		// check for nested maps
		switch vv := v.(type) {
		case map[string]interface{}:
			searchMap(vv, key, fmt.Sprintf("%s%s.", path, k))
		default:
			// It's not a nested map
		}

	}
}
