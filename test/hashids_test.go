package tests

import (
	"fmt"
	"testing"

	"github.com/speps/go-hashids/v2"
	"github.com/stretchr/testify/assert"
)

func TestHashids(t *testing.T) {
	// Test GenerateHashid()
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	assert.Equal(t, 45, getDecodeID(d))
}

func getDecodeID(arr []int) int {
	if len(arr) > 0 {
		return arr[0]
	}

	return 0
}
