package stock_telegram_bot

import (
	"strconv"
	"strings"
)

type Command struct {
	args []string
}

func (c Command) exec() string {
	panic("not support command")
}

type InfoCommand struct {
	Command
}

func (i InfoCommand) exec() string {
	stockData := GetConfig().StockDataApi.BatchQuery(i.args)
	var result string
	for index, data := range stockData {
		result += strconv.FormatInt(int64(index), 10) + "." + new(SimpleFormatter).Format(data)
	}
	return result
}

func ExecCommand(command string) string {
	cmd := parseArgs(command)
	if len(cmd) > 0 {
		if strings.HasPrefix(cmd[0], "/info") {
			infoCommand := &(InfoCommand{
				Command: Command{
					args: cmd[1:],
				},
			})
			return infoCommand.exec()
		}
	}
	return "exec command error"
}

func parseArgs(cmd string) []string {
	var args []string
	var closedChar byte      // 闭合字符 ' or "
	var needClose bool       // 当前是否需要闭合
	var toCompleteArg string // 要拼接的完成的参数，就是 ' 或 " 包围的参数值

	addToArgs := func() {
		toCompleteArg = strings.ReplaceAll(toCompleteArg, `\'`, `'`)
		toCompleteArg = strings.ReplaceAll(toCompleteArg, `\"`, `"`)
		args = append(args, toCompleteArg)
		toCompleteArg = ""
	}

	for i := 0; i < len(cmd); i++ {
		if (cmd[i] == '\'' || cmd[i] == '"') && cmd[i-1] != '\\' { // 不用判断 i > 0
			if !needClose {
				needClose = true
				closedChar = cmd[i]
				continue
			} else if cmd[i] == closedChar {
				needClose = false
				addToArgs()
				continue
			}
		}
		if cmd[i] == ' ' && needClose { // 空格作为参数值的一部分
			toCompleteArg += " "
		} else if cmd[i] == ' ' && toCompleteArg != "" { // 非引号闭合的参数结束
			addToArgs()
		} else if cmd[i] != ' ' {
			toCompleteArg += string(cmd[i])
			if i == len(cmd)-1 {
				addToArgs()
			}
		}
	}
	return args
}
