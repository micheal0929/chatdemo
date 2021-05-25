package procAgent

import (
	"fmt"
	"time"

	"chatdemo/common"
	"chatdemo/ws/sessionAgent"
	"chatdemo/ws/ws_err"
)

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 11:55 上午
 * @Desc:
 */

type Handler func(session *sessionAgent.Session, body []byte) ([][]byte, error)

// Agent 命令处理代理器
type Agent struct {
	procMap map[common.Command]Handler
}

func (a *Agent) AddCommandHandler(reqCmd, rspCmd common.Command, handler Handler)  {
	if _, ok := a.procMap[reqCmd]; ok {
		panic(fmt.Errorf("duplicated command %d", reqCmd))
	}
	a.procMap[reqCmd] = handler
}


func (a *Agent) Run(session *sessionAgent.Session, reqCmd common.Command, body []byte) error {
	defer session.Update(time.Now())

	if handler, ok := a.procMap[reqCmd]; ok {
		return a.runHandler(handler, session, body)
	}
	return ws_err.ErrorNoSuchCommand
}

func (a *Agent) runHandler(h Handler, session *sessionAgent.Session, body []byte) error {
	rspBuffer, err := h(session, body)
	if err != nil {
		return err
	}
	for _, buf := range rspBuffer {
		if conErr := session.Write(buf); conErr != nil {
			fmt.Printf("session %d write ws_err : %v", session.Sid, err)
		}
	}
	return nil
}