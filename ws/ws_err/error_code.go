package ws_err

import "errors"

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 3:01 下午
 * @Desc:
 */

var (
	ErrorNoSuchCommand = errors.New("no such command")
)