package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
	"os"
)

type Body []string

func (o *Body) UnmarshalJSON(data []byte) error {

	var ss []string
	if err := json.Unmarshal(data, &ss); err != nil {
		var s string
		err := json.Unmarshal(data, &s)
		if err != nil {
			return err
		}
		*o = []string{s}
		return nil
	}
	*o = ss

	return nil
}

type Snippet struct {
	name        string
	Prefix      string `json:"prefix"`
	Description string `json:"description"`
	Body        Body   `json:"body"`
}
type Snippets struct {
	snippets []Snippet
}

func (o *Snippets) UnmarshalJSON(data []byte) error {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	o.snippets = make([]Snippet, 0)
	for key := range m {
		var snippet Snippet
		err := json.Unmarshal(m[key], &snippet)
		if err != nil {
			return err
		}
		snippet.name = key
		o.snippets = append(o.snippets, snippet)
	}
	return nil
}

func main() {
	// Open our jsonFile
	snippetFile := flag.String("f", "", "location of the snippet file")
	flag.Parse()
	if len(*snippetFile) == 0 {
		flag.PrintDefaults()
		return
	}
	jsonFile, err := os.Open(*snippetFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)
	var parsed Snippets
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &parsed)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Prefix", "Name", "Description"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	for _, snippet := range parsed.snippets {
		body := ""
		for i, s := range snippet.Body {
			if i != 0 {
				body += "<br />"
			}
			body += fmt.Sprintf("`%s`", s)
		}

		table.Append([]string{fmt.Sprintf("`%s`", snippet.Prefix), fmt.Sprintf("%s", snippet.name), body})
	}
	table.Render()
}
