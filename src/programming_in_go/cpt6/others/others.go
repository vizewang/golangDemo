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

type Place struct  {
	latitude,longitude float64
	Name string
}

func New(latitude, longitude float64, name string) *Place {
	return &Place{latitude,longitude,name}
}
func (place *Place) Latitude() float64 {
	return place.latitude
}
func (place *Place) SetLatitude(latitude float64) {
	place.latitude=latitude;
}
func (place *Place) Longitude()float64{
	return place.longitude
}
func (place *Place) SetLongitude(longitude float64) {
	place.longitude=longitude
}

