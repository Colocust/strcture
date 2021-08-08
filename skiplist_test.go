package strcture

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestSkipList_Insert(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 100000; i++ {
		sl.Insert("s", float32(i))
	}

	n := sl.header.level[0].forward
	for i := 0; i < 100000; i++ {
		assert.Equal(t, float32(i), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}
}

func TestSkipList_UpdateScore(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 100000; i++ {
		sl.Insert("s", float32(i))
	}

	sl.UpdateScore("s", 0, 100000)

	n := sl.header.level[0].forward
	for i := 0; i < 100000; i++ {
		assert.Equal(t, float32(i)+1, n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl.UpdateScore("s", 100000, 0)

	n = sl.header.level[0].forward
	for i := 0; i < 100000; i++ {
		assert.Equal(t, float32(i), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	for i := 0; i < 3; i++ {
		sl.Insert("s", float32(i))
	}
	sl.Insert("s", 7)
	sl.Insert("s", 4)
	sl.UpdateScore("s", 7, 3)

	n = sl.header.level[0].forward
	for i := 0; i < 5; i++ {
		assert.Equal(t, float32(i), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 100000; i++ {
		sl.Insert("s", float32(i))
	}

	res, _ := sl.Delete("s", 0)
	assert.Equal(t, true, res)
	n := sl.header.level[0].forward
	for i := 1; i < 100000; i++ {
		assert.Equal(t, float32(i), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	res, _ = sl.Delete("s", 99999)
	assert.Equal(t, true, res)
	n = sl.header.level[0].forward
	for i := 1; i < 99999; i++ {
		assert.Equal(t, float32(i), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	sl.Insert("s", 1)
	sl.Insert("s", 3)
	sl.Insert("s", 4)
	sl.Insert("s", 5)
	sl.Insert("s", 7)

	res, _ = sl.Delete("s", 4)
	assert.Equal(t, true, res)

	n = sl.header.level[0].forward
	for i := 0; i < 4; i++ {
		assert.Equal(t, float32(i*2+1), n.score)
		n = n.level[0].forward
	}
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	res, _ = sl.Delete("s", 10000)
	assert.Equal(t, false, res)
}

func TestSkipList_FirstInRange(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}

	zrs := NewZRangeSpec(0, 1, false, false)
	res := sl.FirstInRange(zrs)
	assert.Equal(t, float32(0), res.score)

	zrs = NewZRangeSpec(0, 1, true, true)
	res = sl.FirstInRange(zrs)
	assert.Nil(t, res)

	zrs = NewZRangeSpec(0, 1, true, false)
	res = sl.FirstInRange(zrs)
	assert.Equal(t, float32(1), res.score)

	zrs = NewZRangeSpec(0, 2, true, true)
	res = sl.FirstInRange(zrs)
	assert.Equal(t, float32(1), res.score)

	sl = NewSkipList()
	sl.Insert("s", 1)
	sl.Insert("s", 5)
	sl.Insert("s", 6)
	sl.Insert("s", 9)
	zrs = NewZRangeSpec(7, 8, false, false)
	res = sl.FirstInRange(zrs)
	assert.Nil(t, res)
}

func TestSkipList_LastInRange(t *testing.T) {
	sl := NewSkipList()
	sl.Insert("s", 1)
	sl.Insert("s", 5)
	sl.Insert("s", 6)
	sl.Insert("s", 9)
	zrs := NewZRangeSpec(7, 8, false, false)

	res := sl.LastInRange(zrs)
	assert.Nil(t, res)

	sl = NewSkipList()
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}

	zrs = NewZRangeSpec(3, 4, false, false)
	res = sl.LastInRange(zrs)
	assert.Equal(t, float32(4), res.score)

	zrs = NewZRangeSpec(3, 4, false, true)
	res = sl.LastInRange(zrs)
	assert.Equal(t, float32(3), res.score)

	zrs = NewZRangeSpec(3, 4, true, true)
	res = sl.LastInRange(zrs)
	assert.Nil(t, res)
}

func TestSkipList_DeleteByScore(t *testing.T) {
	sl := NewSkipList()
	sl.Insert("s", 1)
	sl.Insert("s", 5)
	sl.Insert("s", 6)
	sl.Insert("s", 9)
	zrs := NewZRangeSpec(7, 8, false, false)

	remove := sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 0, int(remove))

	sl = NewSkipList()
	zrs = NewZRangeSpec(0, 10, false, false)
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}
	remove = sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 11, int(remove))
	n := sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	remove = sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 0, int(remove))

	sl = NewSkipList()
	zrs = NewZRangeSpec(0, 10, true, false)
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}
	remove = sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 10, int(remove))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	zrs = NewZRangeSpec(0, 10, false, true)
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}
	remove = sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 10, int(remove))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	zrs = NewZRangeSpec(0, 10, true, true)
	for i := 0; i < 10000; i++ {
		sl.Insert("s", float32(i))
	}
	remove = sl.DeleteByScore(zrs, NewDict())
	assert.Equal(t, 9, int(remove))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}
}

func TestSkipList_DeleteByRank(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 10; i++ {
		sl.Insert("s", float32(i))
	}
	res := sl.DeleteByRank(1, 1, NewDict())
	assert.Equal(t, 1, int(res))
	n := sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	for i := 0; i < 10; i++ {
		sl.Insert("s", float32(i))
	}
	res = sl.DeleteByRank(4, 7, NewDict())
	assert.Equal(t, 4, int(res))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	for i := 0; i < 10; i++ {
		sl.Insert("s", float32(i))
	}
	res = sl.DeleteByRank(9, 10, NewDict())
	assert.Equal(t, 2, int(res))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	for i := 0; i < 10; i++ {
		sl.Insert("s", float32(i))
	}
	res = sl.DeleteByRank(11, 11, NewDict())
	assert.Equal(t, 0, int(res))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}

	sl = NewSkipList()
	for i := 0; i < 10; i++ {
		sl.Insert("s", float32(i))
	}
	res = sl.DeleteByRank(0, 1, NewDict())
	assert.Equal(t, 1, int(res))
	n = sl.header.level[0].forward
	for l := 0; l < int(sl.level); l++ {
		n = sl.header
		sum := 0
		for n != nil {
			sum += int(n.level[l].span)
			n = n.level[l].forward
		}
		assert.Equal(t, int(sl.length), sum)
	}
}

func TestSkipList_GetRank(t *testing.T) {
	sl := NewSkipList()
	for i := 0; i < 100000; i++ {
		sl.Insert(strconv.Itoa(i), float32(i))
	}

	for i := 0; i < 100000; i++ {
		rank := sl.GetRank(strconv.Itoa(i), float32(i))
		assert.Equal(t, i+1, int(rank))
	}
}
