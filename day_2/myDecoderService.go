package day2

import (
	"encoding/json"
	"sort"
)

type Decoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})

	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type DecoderService struct {
	log *Logger
}

func NewDecoder(log Logger) *DecoderService {
	return &DecoderService{&log}
}

func (DecoderService) Decode(data []byte) ([]Person, []Place) {
	persons := struct {
		P []Person `json:"things"`
	}{}
	if err := json.Unmarshal(data, &persons); err != nil {
		persons.P = nil
	}

	places := struct {
		P []Place `json:"things"`
	}{}
	if err := json.Unmarshal(data, &places); err != nil {
		places.P = nil
	}

	pers := FilterPersons(persons.P)
	pls := FilterPlaces(places.P)

	return pers, pls
}

func (DecoderService) Sort(data interface{}) {
	switch v := data.(type) {
	case []Person:
		sort.Sort(ByAge(v))
	case []Place:
		sort.Sort(ByCityLen(v))
	}
}

func (d DecoderService) Print(i interface{}) {
	(*d.log).Println(i)

}

func (d DecoderService) Printlen(persons []Person, places []Place) {
	(*d.log).Println(len(persons), len(places))
}

//filters a slice of objects with a given predicate
func FilterPersons(ps []Person) []Person {
	n := 0
	for _, p := range ps {
		if !(p.Age == 0 && p.Name == "") {
			ps[n] = p
			n++
		}
	}
	ps = ps[:n]
	return ps
}

func FilterPlaces(ps []Place) []Place {
	n := 0
	for _, p := range ps {
		if !(p.City == "" && p.Country == "") {
			ps[n] = p
			n++
		}
	}
	ps = ps[:n]
	return ps
}
