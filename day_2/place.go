package day2

import "unicode/utf8"

type Place struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type ByCityLen []Place

func (a ByCityLen) Len() int {
	return len(a)
}

func (a ByCityLen) Less(i, j int) bool {
	iStr := utf8.RuneCountInString(a[i].City)
	jStr := utf8.RuneCountInString(a[j].City)
	return iStr < jStr
}

func (a ByCityLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
