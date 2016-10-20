package controller

type IBase interface {
	SetParam(string, interface{})
	GetParam(string) (interface{}, bool)
}

type Base struct {
	params map[string]interface{}
}

func (base *Base) SetParam(key string, value interface{}) {
	if base.params == nil {
		base.params = make(map[string]interface{})
	}
	base.params[key] = value
}

func (base *Base) GetParam(key string) (value interface{}, ok bool) {
	value, ok = base.params[key]
	return
}
