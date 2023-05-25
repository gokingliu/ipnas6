package entity

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk          = Status{0, "Success"}
	ResFail        = Status{-1, ""}
	JsonMarshalErr = Status{-100, "JSON 序列化出错"}
	ClientIPError  = Status{100, "获取 IPv4 和 IPv6 失败"}
)
