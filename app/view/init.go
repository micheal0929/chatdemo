package view

import (
	"chatdemo/app/model/proto"
	"chatdemo/ws/wsAgent"
)

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 3:07 下午
 * @Desc:
 */

func Init(agent *wsAgent.WsAgent)  {
}

func initHandler(agent wsAgent.WsAgent)  {
	agent.AddCommandHandler(proto.CommandJoinReq, proto.CommandJoinRsp, Join)
}