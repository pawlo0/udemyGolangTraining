package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]int)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	j := 0
	for sc.Scan() {
		j++
		words[sc.Text()] = j
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k, v := range words {
		fmt.Println(k, " ", v)
		if i == 200 {
			break
		}
		i++
	}
}
