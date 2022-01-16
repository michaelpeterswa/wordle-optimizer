package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"

	"github.com/michaelpeterswa/wordle-optimizer/internal/calculate"
	"github.com/michaelpeterswa/wordle-optimizer/internal/ingest"
)

func main() {
	choices, answers, err := ingest.GetCurrentWordlists("https://www.powerlanguage.co.uk/wordle/main.c1506a22.js")
	if err != nil {
		log.Printf("failed to acquire current word lists: %s", err)
	}

	master := append(choices, answers...)
	sort.Strings(master)

	fl := calculate.GetCharacterCounts(answers)

	result := calculate.GeneratePowerStarters(answers, *fl, 6)

	f, err := os.Create("result.json")
	if err != nil {
		log.Printf("failed to open output file: %s", err)
	}
	defer f.Close()

	jsonEnc := json.NewEncoder(f)
	jsonEnc.SetIndent("", "\t")

	jsonEnc.Encode(result)
}
