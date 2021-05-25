package wsAgent

import (
	"chatdemo/ws/procAgent"
	"chatdemo/ws/sessionAgent"
)

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 11:33 上午
 * @Desc:
 */

// WsAgent websockets 代理
// 1.记录所有的session
// 2.注册命令路由
type WsAgent struct {
	*procAgent.Agent      // 路由代理
	*sessionAgent.Manager // session代理
}

func NewWsAgent() *WsAgent {
	return &WsAgent{Agent: new(procAgent.Agent), Manager: new(sessionAgent.Manager)}
}
