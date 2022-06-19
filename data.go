package main
/**
 * 抽象数据结构 - 传输中的数据
 */
type Data struct {
	Ip       string   `json:"ip"`
	Type     string   `json:"type"`     // Type是标识信息的类型 login - 登录信息, handshake - 握手信息, system - 系统信息, logout - 退出信息, user - 普通用户信息
	From     string   `json:"from"`     // 消息来源
	Content  string   `json:"content"`  // 消息内容
	User     string   `json:"user"`
	UserList []string `json:"user_list"`
}
