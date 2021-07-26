package task3

import "fmt"

func PrintTriangles(trs []Triangle) {
	fmt.Println("============= Triangles list: ===============")
	for _, tr := range trs {
		if tr.Square == 0 {
			break
		}
		fmt.Println(tr.name, " ", tr.Square)
	}
}
