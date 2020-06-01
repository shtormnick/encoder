package main

import (
	"fmt"
	"strings"
)

func Splitter(s string) []string {
	words := strings.Fields(s)
	return words
}

func Encoder(words []string) string {
	for i, _ := range words {
		w1 := words[i]
		s := w1[len(w1)-1:]
		for _, ch := range w1 {
			char := string([]rune{ch})
			switch {
			case char == "a":
				w1 = strings.ReplaceAll(w1, "a", "1")
			case char == "e":
				w1 = strings.ReplaceAll(w1, "e", "2")
			case char == "i":
				w1 = strings.ReplaceAll(w1, "i", "3")
			case char == "o":
				w1 = strings.ReplaceAll(w1, "o", "4")
			case char == "u":
				w1 = strings.ReplaceAll(w1, "u", "5")
			default:
			}
		}
		if s == "," || s == "." || s == "-" || s == "?" || s == "!" || s == " " || s == ":" || s == "..." {
			words[i] = w1
		} else {
			words[i] = w1
		}
	}
	encoder_word := strings.Join(words, " ")
	return encoder_word
}

func Decoder (words []string) string {
	for i, _ := range words {
		w1 := words[i]
		s := w1[len(w1)-1:]
		for _, ch := range w1 {
			char := string([]rune{ch})
			switch {
			case char == "1":
				w1 = strings.ReplaceAll(w1, "1", "a")
			case char == "2":
				w1 = strings.ReplaceAll(w1, "2", "e")
			case char == "3":
				w1 = strings.ReplaceAll(w1, "3", "i")
			case char == "4":
				w1 = strings.ReplaceAll(w1, "4", "o")
			case char == "5":
				w1 = strings.ReplaceAll(w1, "5", "u")
			default:
			}
		}
		if s == "," || s == "." || s == "-" || s == "?" || s == "!" || s == " " || s == ":" || s == "..." {
			words[i] = w1
		} else {
			words[i] = w1
		}
	}
	dencoder_word := strings.Join(words, " ")
	return dencoder_word
}


func main() {
	word := []string(Splitter("hello, world!"))
	fmt.Printf("%v \n",Encoder(word))
	encoder_word := Decoder(word)
	fmt.Println(encoder_word)
}
