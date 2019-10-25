package piscine

func ConcatParams2(args []string) string {
	var str string
	for i := range args {
		str = str + args[i]
	}
	return str
}
