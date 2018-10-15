package gli


type ActiveCommand struct {
	Cmd         Command
	Description string
}


func newActiveCommand(cmd Command, description string) ActiveCommand {
	return ActiveCommand{cmd, description}
}