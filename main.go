//go:generate go run internal/download/cmudict.go
package main

import (
	_ "embed"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gabe565/mnemonic-ninja/internal"
	"github.com/gabe565/mnemonic-ninja/internal/word"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

//go:embed .cmudict.dict
var CMUdict string

func main() {
	var err error

	db, err := internal.SetupDatabase()
	if err != nil {
		panic(err)
	}

	err = internal.ImportWords(db, CMUdict)
	if err != nil {
		panic(err)
	}

	for {
		var queryType string
		err = survey.AskOne(&survey.Select{
			Message: "Choose lookup method:",
			Options: []string{"word", "number", "quit"},
		}, &queryType)
		if err != nil {
			return
		}

		if queryType == "quit" {
			break
		}

		var query string
		err = survey.AskOne(&survey.Input{
			Message: "Enter " + queryType + ":",
		}, &query)

		var words []word.Word
		if err = db.Where(map[string]string{queryType: query}).Find(&words).Error; err != nil {
			panic(err)
		}

		t := table.NewWriter()
		t.SetStyle(table.StyleRounded)
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Word", "Arpabet", "Number"})
		for _, w := range words {
			t.AppendRow(table.Row{w.Word, w.Arpabet, w.Number})
		}
		t.Render()
		fmt.Printf("Count: %d\n", len(words))
	}
	fmt.Println("Exiting")
}
