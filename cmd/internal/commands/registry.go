package commands

func Registry() map[string]func([]string) {
	return map[string]func([]string){
		"echo": Echo,
		"type": Type,
		"exit": Exit,
		"pwd":  Pwd,
		"cd":   Cd,
	}
}

var Builtins = []string{
	"echo",
	"type",
	"exit",
	"pwd",
	"cd",
}
