package tests

import "github.com/ozoncp/ocp-response-api/domain"

type TestCase struct {
	file            string
	repeats         uint64
	expectedRepeats int
	isError         bool
}

func GetOneToTenInts() (slice []int) {
	slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return
}

func GetOneToTenStrings() (slice []string) {
	slice = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return
}

func GetNums() (slice []int) {
	slice = []int{13, 1, 2, 16, 2, 12, 3, 34, 5, 22, 7, 8, 49}
	return
}

func GetStrings() (slice []string) {
	slice = []string{
		"thirteen", "one", "two", "twelve", "six",
		"twenty two", "twenty four", "three",
		"thirty four", "five", "seven", "eight", "forty nine"}
	return
}

func GetMaps() (dataInts map[int]int, dataStrings map[string]string, dataIntStr map[int]string, dataStrInts map[string]int) {
	dataInts = map[int]int{0: 1, 1: 2}
	dataStrings = map[string]string{"zero": "one", "one": "two"}
	dataIntStr = map[int]string{1: "one", 2: "two"}
	dataStrInts = map[string]int{"one": 1, "two": 2}
	return
}

func Get10Responses() []domain.Response {
	return []domain.Response{
		*domain.NewResponse(0, 0, 0, "text0"),
		*domain.NewResponse(3, 3, 3, "text3"),
		*domain.NewResponse(2, 2, 2, "text2"),
		*domain.NewResponse(1, 1, 1, "text1"),
		*domain.NewResponse(4, 4, 4, "text4"),
		*domain.NewResponse(5, 5, 5, "text5"),
		*domain.NewResponse(6, 6, 6, "text6"),
		*domain.NewResponse(7, 7, 7, "text7"),
		*domain.NewResponse(8, 8, 8, "text8"),
		*domain.NewResponse(9, 9, 9, "text9"),
	}
}

func GetFileRepeatsData() []TestCase {
	return []TestCase{
		{"mock_data.go", 10000, 10000, false},
		{"mock_data_2.go", 1, 0, true},
		{"", 1, 0, true},
		{"mock_data.go", 0, 0, false},
	}
}
