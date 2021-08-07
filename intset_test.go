package strcture

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntSet_Add(t *testing.T) {
	is := NewIntSet()
	for i := 1; i < 9; i++ {
		is.Add(i)
	}
	for i := 1; i < 9; i++ {
		assert.Equal(t, i, is.Get(i-1))
	}

	is.Add(0)
	assert.Equal(t, 0, is.Get(0))

	is.Add(9)
	assert.Equal(t, 9, is.Get(9))

	for i := 0; i < 10; i++ {
		assert.Equal(t, i, is.Get(i))
	}

	is = NewIntSet()
	for i := 0; i < 5; i++ {
		is.Add(i)
	}
	for i := 6; i < 10; i++ {
		is.Add(i)
	}
	is.Add(5)

	assert.Len(t, is.contents, 10)

	for i := 0; i < 10; i++ {
		assert.Equal(t, i, is.Get(i))
	}
}

func TestIntSet_Remove(t *testing.T) {
	is := NewIntSet()

	for i := 0; i < 10; i++ {
		is.Add(i)
	}

	res := is.Remove(-1)
	assert.Equal(t, false, res)

	res = is.Remove(9)
	assert.Equal(t, true, res)
	assert.Equal(t, 8, is.Get(is.Len()-1))

	is.Remove(0)
	assert.Equal(t, true, res)
	assert.Equal(t, 1, is.Get(0))
}

func TestIntSet_Find(t *testing.T) {
	is := NewIntSet()
	for i := 0; i < 10; i++ {
		is.Add(i)
	}

	pos := is.Find(0)
	assert.Equal(t, true, pos)

	pos = is.Find(11)
	assert.Equal(t, false, pos)
}
