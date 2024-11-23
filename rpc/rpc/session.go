package rpc

import (
	"encoding/binary"
	"net"
)

// 测试网络中读写数据的情况

// 绘画连接的结构体
type Session struct {
	conn net.Conn
}

// 构造方法
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中去写数据
func (s *Session) Write(data []byte) error {
	// 定义写数据的格式
	// 4字节头部 + 可变体的长度
	buf := make([]byte, 4+len(data))
	// 写入头部，记录数据长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data)))
	// 将整个数据放到4后边
	copy(buf[4:], data)
	return nil
}
