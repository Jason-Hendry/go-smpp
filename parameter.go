package go_smpp

type Parameter struct {
	tag    uint16
	length uint16
	value  []byte
}

func (p *Parameter) GetValue() string {
	return string(p.value)
}

func NewParameter(tag uint16, value []byte) Parameter {
	return Parameter{tag: tag, length: uint16(len(value)), value: value}
}
