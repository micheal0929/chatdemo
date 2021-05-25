package sessionAgent

/**
 * @Author: michael.plr
 * @Date: 2021/5/22 11:35 上午
 * @Desc:
 */

// Manager session 代理
// 管理session链接
type Manager struct {
	sessions map[int64]*Session // session信息管理
}
