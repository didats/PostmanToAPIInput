package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func scanInput(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + ":")
	text, _ := reader.ReadString('\n')

	return strings.ReplaceAll(text, "\n", "")
}

func main() {
	//location := scanTextLocation()
	//fileLocation := "/Users/didats/Downloads/CATERING.postman_collection.json"
	//baseStruct := "APIDetail"

	fileLocation := scanInput("Postman file path")

	file, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		fmt.Println("Could not locate '"+fileLocation+"' to read", err)
		os.Exit(3)
	}

	dir, _ := os.Getwd()

	saveLocation := scanInput("Path to save the file (default " + dir + ")")
	baseStruct := scanInput("Base struct name (default: APIInputType)")
	if len(baseStruct) < 1 {
		baseStruct = "APIInputType"
	}

	content := string(file)
	res := Postman{}
	json.Unmarshal([]byte(content), &res)

	swift := NewSwiftExporter(res, baseStruct)
	fileContent := swift.fileContent()

	f, err := os.Create(saveLocation + "/APIInput.swift")
	if err != nil {
		fmt.Println("Could not open the file with error: ", err)
		os.Exit(3)
	}

	f.WriteString(fileContent)
	defer f.Close()

	fmt.Println("File saved to: " + saveLocation + "/APIInput.swift")
}
