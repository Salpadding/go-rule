package rule

import (
	"testing"
)

func TestPluckInt(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
	})
	c0, err := ctx.PluckInt("c0")
	if err != nil || c0 != 12 {
		panic("fail")
	}
}

func TestPluckInts(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
		"c1": 13,
	})
	cArray, err := ctx.PluckInts([]string{"c0", "c1"})
	if err != nil || cArray[0] != 12 || cArray[1] != 13 {
		panic("fail")
	}
}

func TestSetBool(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12,
		"c1": 13,
	})
	ctx.SetBool("isSport", true)
	isSport, err := ctx.PluckBool("isSport")
	if err != nil || !isSport {
		panic("fail")
	}
}

func TestPluckFloat64(t *testing.T) {
	ctx := NewContext(map[string]interface{}{
		"c0": 12.0,
	})
	c0, err := ctx.PluckFloat64("c0")
	if err != nil || c0 != 12 {
		panic("fail")
	}
}
