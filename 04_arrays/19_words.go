package main

import (
	"fmt"
	"strings"
)

//Write a function to count the frequency of words in a string of text and return
//a map of words with their counts. The function should convert the text to lowercase,
//and punctuation should be trimmed from words. The strings package contains several
//helpful functions for this task, including Fields, ToLower, and Trim.

func main() {
	word_groups := make(map[string]int)
	var words string = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	word_slice := strings.Fields(words)

	fmt.Println(word_slice)
	fmt.Println(len(word_slice))

	for _, word := range word_slice {
		format_word := strings.ToLower(strings.Trim(word, ";.,"))
		word_groups[format_word]++
	}

	fmt.Println(word_groups)
	for word, val := range word_groups {
		fmt.Println(word, val)
	}

}
