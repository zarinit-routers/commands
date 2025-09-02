package commands

import "fmt"

type Argument struct {
	Name         string
	Required     bool
	Type         DataType
	AcceptValues []string

	DocsDefault     any
	DocsDescription string
}

func (a *Argument) WithDocsDefault(val any) *Argument {
	a.DocsDefault = val
	return a
}

func StringArg(name string) *Argument {
	return &Argument{
		Name:     name,
		Required: true,
		Type:     DataTypeString,
	}
}

func (a *Argument) String() string {
	requiredStr := ""
	if a.Required {
		requiredStr = "(required)"
	}
	return fmt.Sprintf("%s %s %s", a.Type.String(), a.Name, requiredStr)
}

func (a *Argument) NotRequired() *Argument {
	a.Required = false
	return a
}

func (a *Argument) GetType() *DataType {
	return &a.Type
}

func IntArg(name string) *Argument {
	return &Argument{
		Name:     name,
		Required: true,
		Type:     DataTypeInt,
	}
}

func (a *Argument) WithValues(vals ...string) *Argument {
	a.AcceptValues = vals
	return a
}

func (a *Argument) GetDefaultValue() string {
	if a.AcceptValues != nil {
		if a.Type == DataTypeString {
			return "\"" + a.AcceptValues[0] + "\""
		}
		return a.AcceptValues[0]
	}
	// if a.DocsDefault != nil {
	// 	return fmt.Sprintf("%v", a.DocsDefault)
	// }
	switch a.Type {
	case DataTypeString:
		if a.DocsDefault == nil {
			return "\"\""
		}
		return fmt.Sprintf("\"%s\"", a.DocsDefault.(string))
	case DataTypeInt:
		return fmt.Sprintf("%d", a.DocsDefault.(int))
	case DataTypeBool:
		return "false"
	}
	return ""
}

func (a *Argument) WithDescription(desc string) *Argument {
	a.DocsDescription = desc
	return a
}
