package commands

type Cmd struct {
	Name            string
	Command         string
	Icon            string
	UseSelectedList bool
}

func Screencap(filename string) *Cmd {
	return &Cmd{
		Name:            "",
		Command:         "adb exec-out screencap -p >" + filename,
		Icon:            "",
		UseSelectedList: false,
	}
}

func Clear() *Cmd {
	return &Cmd{
		Name:            "",
		Command:         "clear",
		Icon:            "",
		UseSelectedList: false,
	}
}
