package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/michaelpeterswa/wordle-optimizer/internal/calculate"
	"github.com/michaelpeterswa/wordle-optimizer/internal/ingest"
)

func main() {
	depthPtr := flag.Int("depth", 6, "How deep do you want to search? (max:26)")
	outfilePtr := flag.String("out", "default.json", "The file you want to output to...")

	flag.Parse()
	fmt.Printf("running with config: depth=%d, out=\"%s\"\n", *depthPtr, *outfilePtr)

	choices, answers, err := ingest.GetCurrentWordlists("https://www.powerlanguage.co.uk/wordle/main.c1506a22.js")
	if err != nil {
		log.Printf("failed to acquire current word lists: %s", err)
	}

	master := append(choices, answers...)
	sort.Strings(master)

	fl := calculate.GetCharacterCounts(answers)

	result := calculate.GeneratePowerStarters(answers, *fl, *depthPtr)

	f, err := os.Create(*outfilePtr)
	if err != nil {
		log.Printf("failed to open output file: %s", err)
	}
	defer f.Close()

	jsonEnc := json.NewEncoder(f)
	jsonEnc.SetIndent("", "\t")

	jsonEnc.Encode(result)
}
