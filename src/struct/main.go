package main

import (
	"bytes"
	"fmt"
	"math"
)

type Point struct {
	X, Y float64

	// field and method with the same name Distance
	// Distance int
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

const BITS = 32 << (^uint(0) >> 63)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/BITS, uint(x%BITS)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/BITS, uint(x%BITS)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(values ...int) {
	for _, value := range values {
		s.Add(value)
	}
}

// 并集：元素在集合 A 或集合 B 中出现
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 交集：元素在集合 A 和集合 B 中均出现
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 差集：元素在集合 A 中出现，在集合 B 中未出现
func (s *IntSet) DifferenceWith(t *IntSet) {

}

// 并差集：元素出现在集合 A 但没有出现在集合 B，或者出现在集合 B 但没有出现在集合 A
func (s *IntSet) SymmetricDifference(t *IntSet) {

}

// 返回集合中的所有元素
func (s *IntSet) Elems() (res []uint) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITS; j++ {
			if word&(1<<j) != 0 {
				res = append(res, uint(BITS*i+j))
			}
		}
	}
	return
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITS; j++ {
			if word&(1<<j) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BITS*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet) Len() int {
	return -1
}

// remove x from the set
func (s *IntSet) Remove(x int) {

}
func (s *IntSet) Clear() {

}

func (s *IntSet) Copy() *IntSet {
	return nil
}

func main() {
	point1 := Point{3, 4}
	fmt.Printf("%.3f\n", point1.Distance(Point{5, 6}))

	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	x.AddAll(1, 3, 4)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())
	fmt.Println(x.Has(9), x.Has(123))
}
