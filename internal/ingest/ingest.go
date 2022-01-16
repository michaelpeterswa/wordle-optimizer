package ingest

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetCurrentWordlists(url string) ([]string, []string, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	choiceList, answerList, err := isolateWordlists(content)
	if err != nil {
		return nil, nil, err
	}

	return cleanAndSplitWordlist(choiceList), cleanAndSplitWordlist(answerList), nil
}

func isolateWordlists(body []byte) (string, string, error) {
	var (
		choiceListRegex string
		answerListRegex string
	)

	choiceListRegex = `Ta=\[("[a-z]{5}",)+("[a-z]{5}"])`
	answerListRegex = `La=\[("[a-z]{5}",)+("[a-z]{5}"])`

	choiceRegex, err := regexp.Compile(choiceListRegex)
	if err != nil {
		return "", "", nil
	}

	answerRegex, err := regexp.Compile(answerListRegex)
	if err != nil {
		return "", "", nil
	}

	choiceResult := choiceRegex.Find(body)
	if choiceResult == nil {
		return "", "", nil
	}

	answerResult := answerRegex.Find(body)
	if answerResult == nil {
		return "", "", nil
	}

	return string(choiceResult), string(answerResult), nil
}

func cleanAndSplitWordlist(s string) []string {
	return strings.Split(strings.ReplaceAll(strings.ReplaceAll(s[4:len(s)-1], "\"", ""), "]", ""), ",")
}
