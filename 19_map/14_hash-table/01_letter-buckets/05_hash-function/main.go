package main

import "fmt"

func main() {
	n := Hashbucket("Anthony", 12)
	fmt.Println(n)
}

func Hashbucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
