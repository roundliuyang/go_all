package main

import (
	"bytes"
	"sync"
)

/*
可以使用 sync.Pool 做 buffer 池，但是，如果用这种方式做 buffer 池的话，可能会有内存泄漏的风险。

取出来的 bytes.Buffer 在使用的时候，我们可以往这个元素中增加大量的 byte 数据，这会导致底层的 byte slice 的容量可能会变得很大。这个时候，
即使 Reset 再放回到池子中，这些 byte slice 的容量不会改变，所占的空间依然很大。而且，因为 Pool 回收的机制，这些大的 Buffer 可能不被回收，
而是会一直占用很大的空间，这属于内存泄漏的问题。
*/
var buffers = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func GetBuffer() *bytes.Buffer {
	return buffers.Get().(*bytes.Buffer)
}

func PutBuffer(buf *bytes.Buffer) {
	buf.Reset()
	buffers.Put(buf)
}
