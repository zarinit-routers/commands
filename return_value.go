package commands

type ReturnValue struct {
	Name string
	Type DataType
}

func IntValue(name string) *ReturnValue {
	return &ReturnValue{
		Name: name,
		Type: DataTypeInt,
	}
}

func StringValue(name string) *ReturnValue {
	return &ReturnValue{
		Name: name,
		Type: DataTypeString,
	}
}
