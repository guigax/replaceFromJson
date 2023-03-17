package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// Read the JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Read the text file to be searched
	textFile, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer textFile.Close()

	textData, err := ioutil.ReadAll(textFile)
	if err != nil {
		log.Fatal(err)
	}

	// Replace the keys with the corresponding values
	for key, value := range data {
		replaceKey := fmt.Sprintf("%s", key)
		var replaceValue string
		switch v := value.(type) {
		case string:
			replaceValue = fmt.Sprintf("\"%s\"", v)
		case bool:
			replaceValue = fmt.Sprintf("%t", v)
		case int, int8, int16, int32, int64:
			replaceValue = fmt.Sprintf("%d", v)
		case uint, uint8, uint16, uint32, uint64:
			replaceValue = fmt.Sprintf("%d", v)
		case float32, float64:
			replaceValue = fmt.Sprintf("%f", v)
		default:
			log.Printf("Unsupported value type for key %s: %v", key, v)
			continue
		}
		textData = []byte(strings.ReplaceAll(string(textData), replaceKey, replaceValue))
	}

	// Write the modified text back to the file
	err = ioutil.WriteFile("text.txt", textData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
