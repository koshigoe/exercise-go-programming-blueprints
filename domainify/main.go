package main

import (
	"bufio"
	"math/rand"
	"os"
	"time"
	"unicode"
	"strings"
	"fmt"
)

var tlds = []string { "com", "net" }
const allowedChars = "abcdefghijklmnopqrstuvwxyz01234567890_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
