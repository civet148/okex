package wInterface

import (
	"github.com/civet148/okex/types"
)

// 请求数据
type WSParam interface {
	EventType() types.Event
	ToMap() *map[string]string
}
