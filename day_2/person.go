package day2

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ByAge []Person

func (ba ByAge) Len() int {
	return len(ba)
}

func (ba ByAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba ByAge) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}
