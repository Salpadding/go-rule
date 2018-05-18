package rule

import (
	"testing"
)

func TestPluckInt(t *testing.T) {
	ctx := newContext(map[string]interface{}{
		"c0": 12,
	})
	c0, err := ctx.pluckInt("c0")
	if err != nil || c0 != 12 {
		panic("fail")
	}
}

func TestPluckInts(t *testing.T) {
	ctx := newContext(map[string]interface{}{
		"c0": 12,
		"c1": 13,
	})
	cArray, err := ctx.pluckInts([]string{"c0", "c1"})
	if err != nil || cArray[0] != 12 || cArray[1] != 13 {
		panic("fail")
	}
}

func TestSetBool(t *testing.T) {
	ctx := newContext(map[string]interface{}{
		"c0": 12,
		"c1": 13,
	})
	ctx.setBool("isSport", true)
	isSport, err := ctx.pluckBool("isSport")
	if err != nil || !isSport {
		panic("fail")
	}
}
