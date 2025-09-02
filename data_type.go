package commands

type DataType int

const (
	DataTypeInt DataType = iota
	DataTypeString
	DataTypeBool
	DataTypeFloat
)

func (t DataType) String() string {
	switch t {
	case DataTypeInt:
		return "int"
	case DataTypeString:
		return "string"
	case DataTypeBool:
		return "bool"
	case DataTypeFloat:
		return "float"
	default:
		return "unknown"
	}
}
