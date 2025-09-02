package main

import (
	"fmt"

	"github.com/zarinit-routers/commands"
)

func main() {
	for _, cmd := range commands.Commands {

		printHeader(cmd.String())

		printDescription(cmd)

		printArgs(cmd.Args)
		printRequestExample(cmd)
		printReturns(cmd.ReturnValues)
	}
}

func printHeader(h string) {
	fmt.Println("## ", h)
	fmt.Println()
}
func printDescription(cmd *commands.Command) {
	if !cmd.ReadyToUse {
		fmt.Println("_**Not ready to use!**_")
		fmt.Println()
	}
	if cmd.Description != "" {
		fmt.Println(cmd.Description)
		fmt.Println()
	}

}

func printArgs(args []*commands.Argument) {

	if len(args) == 0 {
		fmt.Println("")
		return
	}

	fmt.Println("**Accepts:**")
	fmt.Println("")
	for _, arg := range args {
		printArg(arg)
		fmt.Println()
	}
}

func printRequestExample(cmd *commands.Command) {
	if len(cmd.Args) == 0 {
		return
	}
	fmt.Println("**Request example:**")
	fmt.Println()
	fmt.Println("```json")
	fmt.Println("{")

	for index, arg := range cmd.Args {

		requiredStr := "required"
		if !arg.Required {
			requiredStr = "not required"
		}

		fmt.Printf("  \"%s\": %s", arg.Name, arg.GetDefaultValue())
		if index < len(cmd.Args)-1 {
			fmt.Printf(",")
		}
		fmt.Printf(" // %s", requiredStr)
		fmt.Println()
	}
	fmt.Println("}")
	fmt.Println("```")
	fmt.Println()
}

func printArg(arg *commands.Argument) {
	fmt.Printf("`%s` of type %s. %s\n", arg.Name, arg.Type, arg.DocsDescription)
}

func printReturns(vals []*commands.ReturnValue) {
	if len(vals) == 0 {
		return
	}
	fmt.Println("**Returns:**")
	fmt.Println()
	for _, val := range vals {
		printReturnVal(val)
		fmt.Println()
	}
}

func printReturnVal(val *commands.ReturnValue) {
	fmt.Printf("`%s` of type %s.\n", val.Name, val.Type)
}
