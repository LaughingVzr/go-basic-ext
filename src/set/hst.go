package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Clear() {
	// map重新赋值
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Equal(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return false
}

func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			// 由于并发的原因，这里的initialLen有可能发生变化
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		// 截取方法调用时的实际长度的切片，用于应对切片长度变小的情况
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil {
		return false
	}
	oneLen := set.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) {
			return false
		}
	}
	return true
}

func (set *HashSet) Union(other *HashSet) *HashSet {
	if set == nil || other == nil {
		return nil
	}
	union := NewHashSet()
	for v, key := range set {
		union.set[key] = v
	}
	if other.Len() == 0 {
		return union
	}
	for key := range other {
		if !set.Contains(key) {
			union.Add(key)
		}
	}
	return union
}

func (set *HashSet) Intersect(other *HashSet) *HashSet {
	if set == nil || other == nil {
		return nil
	}
	compSign := SetCompare(set, other)
	inst := NewHashSet()
	switch compSign {
	case 0:
		inst = DoIntersect(set, other)
	case 1:
		inst = DoIntersect(other, set)
	case -1:
		inst = DoIntersect(set, other)
	}
	if inst == nil {
		return nil
	}
	return inst
}

func (set *HashSet) Difference(other *HashSet) *HashSet {
	if set == nil || other == nil {
		return nil
	}
	diffSet := NewHashSet()
	if other.Len() == 0 {
		for key := range set {
			diffSet.Add(key)
		}
		return diffSet
	}
	for key := range set {
		if !other.Contains(key) {
			diffSet.Add(key)
		}
	}
	return diffSet
}

func (set *HashSet) SymmetricDifference(other *HashSet) *HashSet {
	if set == nil || other == nil {
		return nil
	}
	diffset := set.Difference(other)
	if other.Len() == 0 {
		return diffset
	}
	diffother := other.Difference(set)
	return diffset.Union(diffother)
}

/*
* 比较两个集合长度大小.
* 1：x集合大
*-1：y集合大
* 0：相等
**/
func SetCompare(x *HashSet, y *HashSet) int {
	xlen := x.Len()
	ylen := y.Len()
	if x > y {
		return 1
	} else if x < y {
		return -1
	} else {
		return 0
	}
}

/**
做交集处理
*/
func DoIntersect(min *HashSet, max *HashSet) (inset *HashSet) {
	inset = NewHashSet()
	if min.Len() == 0 {
		return
	}
	for key := range min {
		if max.Contains(key) {
			inset.Add(key)
		}
	}
	return
}

/**
差集处理
*/
func DoDiff(min *HashSet, max *HashSet) (diffset *HashSet) {
	diffset = NewHashSet()
	if condition {

	}
}
