package sessionAgent

import (
	"time"

	"github.com/gorilla/websocket"
)

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 11:37 上午
 * @Desc:
 */

// Session session连接
type Session struct {
	Conn         *websocket.Conn // 底层ws连接
	UBaseInfo                    // 用户信息
	MachineID    int64           // 设备ID
	Sid          int64           // session id
	LastModified int64           // 最后操作
}

// UBaseInfo 用户基础信息
type UBaseInfo struct {
	Uid   int64
	Coins int64
	Name  string
	Icon  string
}

func (s *Session) Update(t time.Time)  {
	s.LastModified = t.UnixNano()
}

func (s *Session) Write(buf []byte) error {
	return s.Conn.WriteMessage(websocket.TextMessage, buf)
}