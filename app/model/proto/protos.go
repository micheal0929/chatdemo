package proto

import "chatdemo/common"

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 3:10 下午
 * @Desc:
 */

// req and rsp
const (
	CommandJoinReq = common.Command(1) // 加入聊天室
	CommandJoinRsp = common.Command(2) //
	CommandQuitReq = common.Command(3) // 退出聊天室
	CommandQuitRsp = common.Command(4) //
	CommandChatReq = common.Command(5) // 发送聊天信息
	CommandChatRsp = common.Command(6) //
	CommandGiftReq = common.Command(7) // 赠送礼物信息
	CommandGiftRsp = common.Command(8) //
)

// push -- system
const (
	CommandRoomClosePush     = common.Command(100) // 聊天室关闭推送
	CommandRoomBroadCastPush = common.Command(102) // 聊天室全平台广播推送
)
