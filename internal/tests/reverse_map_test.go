package tests

import (
	"github.com/ozoncp/ocp-response-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	intMap, stringsMap, intStrMap, strIntMap := GetMaps()
	reverseMapInts := utils.ReverseIntInt(intMap)
	reverseMapStrings := utils.ReverseStringString(stringsMap)
	reverseMapIntStr := utils.ReverseIntString(intStrMap)
	reverseMapStrInt := utils.ReverseStringInt(strIntMap)

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
	reverseMap := utils.ReverseIntInt(emptyMap)

	assert.Equal(t, reverseMap, emptyMap)
}
