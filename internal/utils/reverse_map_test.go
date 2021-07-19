package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getMaps() (dataInts map[int]int, dataStrings map[string]string, dataIntStr map[int]string, dataStrInts map[string]int) {
	dataInts = map[int]int{0: 1, 1: 2}
	dataStrings = map[string]string{"zero": "one", "one": "two"}
	dataIntStr = map[int]string{1: "one", 2: "two"}
	dataStrInts = map[string]int{"one": 1, "two": 2}
	return
}

func TestReverse(t *testing.T) {
	intMap, stringsMap, intStrMap, strIntMap := getMaps()
	reverseMapInts := ReverseIntInt(intMap)
	reverseMapStrings := ReverseStringString(stringsMap)
	reverseMapIntStr := ReverseIntString(intStrMap)
	reverseMapStrInt := ReverseStringInt(strIntMap)

	assert.Equal(t, len(intMap), len(reverseMapInts))
	assert.Equal(t, len(stringsMap), len(reverseMapStrings))
	assert.Equal(t, len(intStrMap), len(reverseMapIntStr))
	assert.Equal(t, len(strIntMap), len(reverseMapStrInt))

	for k, v := range intMap {
		assert.Equal(t, k, reverseMapInts[v])
	}
	for k, v := range stringsMap {
		assert.Equal(t, k, reverseMapStrings[v])
	}
	for k, v := range intStrMap {
		assert.Equal(t, k, reverseMapIntStr[v])
	}
	for k, v := range strIntMap {
		assert.Equal(t, k, reverseMapStrInt[v])
	}
}

func TestEmptyMap(t *testing.T) {
	emptyMap := map[int]int{}
	reverseMap := ReverseIntInt(emptyMap)

	assert.Equal(t, reverseMap, emptyMap)
}
