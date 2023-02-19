package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	j "github.com/dave/jennifer/jen"
)

type entry struct {
	Word string `json:"word,omitempty"`
}

func main() {
	// Get dictionary from github
	url := "https://raw.githubusercontent.com/cduica/Oxford-Dictionary-Json/master/dicts.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("unable to retrieve dictionary json: ", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("got non-200 response", resp.StatusCode)
		os.Exit(1)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("unable to read response: ", err)
		os.Exit(1)
	}

	dictionary := make([]entry, 0)
	if err := json.Unmarshal(data, &dictionary); err != nil {
		fmt.Println("unable to unmarshal response: ", err)
		os.Exit(1)
	}

	// Build the go file
	f := j.NewFile("dictionary")

	// Build the values
	values := j.Dict{}
	for _, entry := range dictionary {
		// Filter out some weird entries
		if strings.HasPrefix(entry.Word, "-") {
			continue
		}
		values[j.Lit(entry.Word)] = j.True()
	}

	// Build and render the actual code
	f.Var().Id("Entries").Op("=").Map(j.String()).Bool().Values(values)
	if err := os.WriteFile("entries.go", []byte(f.GoString()), 0644); err != nil {
		fmt.Println("unable to write file: ", err)
		os.Exit(1)
	}

}
