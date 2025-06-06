package fortune

import "math/rand"

var fortunes = []string{
	"Fortune favors the bold.",
	"Success is not the key to happiness. Happiness is the key to success.",
	"Your future is bright and full of opportunities.",
}

func Get() string {
	return fortunes[rand.Intn(len(fortunes))]
}
