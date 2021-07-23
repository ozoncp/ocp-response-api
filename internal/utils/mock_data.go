package utils

type TestCase struct {
	file            string
	repeats         int
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

func GetStrings() []string {
	return []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"twelve", "thirteen", "twenty two", "twenty four", "thirty four", "forty nine"}
}

func GetMaps() (dataInts map[int]int, dataStrings map[string]string, dataIntStr map[int]string, dataStrInts map[string]int) {
	dataInts = map[int]int{0: 1, 1: 2}
	dataStrings = map[string]string{"zero": "one", "one": "two"}
	dataIntStr = map[int]string{1: "one", 2: "two"}
	dataStrInts = map[string]int{"one": 1, "two": 2}
	return
}

func Get10Responses() []Response {
	return []Response{
		{Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
		{Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
		{Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
		{Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
		{Id: 4, UserId: 4, RequestId: 4, Text: "text4"},
		{Id: 5, UserId: 5, RequestId: 5, Text: "text5"},
		{Id: 6, UserId: 6, RequestId: 6, Text: "text6"},
		{Id: 7, UserId: 7, RequestId: 7, Text: "text7"},
		{Id: 8, UserId: 8, RequestId: 8, Text: "text8"},
		{Id: 9, UserId: 9, RequestId: 9, Text: "text9"},
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
