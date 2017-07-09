package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

func main() {
	//get the book moby dick
	res, err := http.Get("https://projecteuler.net/project/resources/p022_names.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the file
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	// Split the words in the file into a slice
	var words []string
	var word string
	for _, v := range body {
		if v != 34 && v != 44 {
			word += string(v)
		} else if v == 34 && len(word) > 0 {
			words = append(words, word)
			word = ""
		}
	}

	// sort the words
	sort.Strings(words)

	// Loop through the sorted words and get their score
	var totScore, wrdScore int
	for i, wrd := range words {
		wrdScore = 0
		// Loop through letters of this word and get word score
		for _, letter := range wrd {
			wrdScore += int(letter) - 64
		}
		// Add word score to position score
		totScore += (i + 1) * wrdScore
	}

	fmt.Println(totScore)
}
