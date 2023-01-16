package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppend(t *testing.T) {
	obj := Post{Id: int64(11), ParentId: int64(1), Content: "hi", CreateTime: int64(1650437618)}
	output, _ := appendlocalpost(obj)
	expect := true
	assert.Equal(t, expect, output)
}
