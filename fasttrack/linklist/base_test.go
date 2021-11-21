package linklist

import (
	"net/url"
	"testing"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}


//
// 应该写一个 url decode 的小工具🔧
func TestHttp(t *testing.T){
	t.Log(url.QueryUnescape("%F0%9F%88%9A"))
	t.Log(url.QueryUnescape("%F0%9F%88%9A%EF%B8%8F"))

	t.Log(len("Iw9hJz2uGNXlKzIwMDAwMDA5OBYAMIM8D2gl"))

	t.Log((int8)(1))
	t.Log(int8(1))
}