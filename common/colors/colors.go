package colors

const (
	color_reset = "\033[0m"
	color_red   = "\033[31m"
	color_green = "\033[32m"
)

func Red(str string) string {
	return color_red + str + color_reset
}

func Green(str string) string {
	return color_green + str + color_reset
}
