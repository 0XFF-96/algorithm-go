package array

import (
	"bytes"
	"fmt"
	"testing"
)

// this snippet is from gopl.
// 这个和

type IntSet struct {
	words []uint64
}

func (s *IntSet)Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int){
	word, bit := x/64, uint(x%64)

	// 扩容
	// 加零
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	//
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func(s *IntSet) UnionWith(t *IntSet){
	for i, tword := range t.words {
		if i < len(s.words){
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// string returns the set as a string of the from
func (s *IntSet) String() string{
	// non-copy
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words{
		if word == 0 {
			continue
		}

		for j :=0;j < 64;j++{
			if word&(1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				// 真正的计算方法
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}


func TestIntSt(t *testing.T){

	//var x IntSet
	//// x.Add(1)
	//x.Add(65)
	//// x.Add(66)
	//
	//t.Log(1 | 1)
	//t.Log(^1)
	//t.Log(1 & 1)

	t.Log(1/64)
	t.Log(1 % 64)

	t.Log(2/64)
	t.Log(2 %64)

	t.Log(3 / 64)
	t.Log(3 % 64)

	t.Log(64 / 64)
	t.Log(64 % 64)
	t.Log(65 / 64)
	t.Log(65 % 64)

	t.Log(1 << 3)

	// 这里相当于  3 + 1 << 3 位
	t.Log(3 | 1 << 3)
}

// https://learnku.com/go/t/23460/bit-operation-of-go
// https://blog.csdn.net/skh2015java/article/details/78451033
// https://www.cnblogs.com/piperck/p/6139369.html

// 比特计算学习。
//
// https://learnku.com/go/t/23460/bit-operation-of-go

// redis:https://redisbook.readthedocs.io/en/latest/compress-datastruct/intset.html
// https://redisbook.readthedocs.io/en/latest/compress-datastruct/intset.html
// 厉害
