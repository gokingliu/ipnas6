package entity

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk             = Status{0, "Success"}
	ResFail           = Status{-1, ""}
	JsonMarshalErr    = Status{-100, "JSON 序列化出错"}
	ClientIPError     = Status{100, "获取 IPv4 和 IPv6 失败"}
	ClientGetError    = Status{101, "客户端 Get 请求失败"}
	ClientPostError   = Status{102, "客户端 Post 请求失败"}
	ClientPutError    = Status{103, "客户端 Put 请求失败"}
	ClientGetDNSError = Status{104, "获取 Cloudflare DNS 失败"}
)
