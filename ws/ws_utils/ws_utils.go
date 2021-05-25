package ws_utils

import (
	"chatdemo/common"
	"chatdemo/ws/wsAgent"
)

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 3:28 下午
 * @Desc:
 */


func AddSimpleHandler(agent *wsAgent.WsAgent, reqCmd, rspCmd common.Command, proc interface{}, funcList ...func(*wsAgent.WsAgent) ([]byte, error)) {

}