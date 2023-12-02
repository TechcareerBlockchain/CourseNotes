package utils

import "fmt"

type Functions struct {
	V1 int `json:"v_1"`
	V2 int `json:"v_2"`
}

func (f *Functions) Plus() int {
	return f.V1 + f.V2
}

func (f *Functions) PlusWithoutReturn() {
	fmt.Println(f.V1 + f.V2)
}

func (f *Functions) ReverseData() (int, int) {
	return f.V2, f.V1
}

func (f *Functions) ReverseDataV2() (res1 int, res2 int) {
	res1 = f.V2
	res2 = f.V1
	return
}
