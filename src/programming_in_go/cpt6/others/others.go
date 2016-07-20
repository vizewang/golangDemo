package others

import (
	"unicode"
	"fmt"
	"strings"
)

type RuneForRuneFunc func(rune)rune

func Demo1() {
	var removePunctuation RuneForRuneFunc
	phrases:=[]string{"Day; dusk, and night.","All day long"}
	removePunctuation= func(char rune) rune{
		if unicode.Is(unicode.Terminal_Punctuation,char){
			return -1
		}
		return char
	}
	processPhrases(phrases,removePunctuation)
}
func processPhrases(phrases []string, function RuneForRuneFunc) {
	for _,phrase:=range phrases{
		fmt.Println(strings.Map(function,phrase))
	}
}