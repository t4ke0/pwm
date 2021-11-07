package pw_generator

import (
	"math/rand"
	"os"
	"strings"
	"time"

	pb "github.com/t4ke0/pwm/pwm_manager/proto"
)

func readWordFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}
	return strings.Split(string(data), "\n"), nil
}

const (
	letters      string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers      string = "0123456789"
	specialChars string = "~#{[|`^@]=+/"
)

var WordsFilename string = os.Getenv("WORDS_FILENAME")

// Generate generates random password accept length and the mode `easy`, `medium` or `complex`.
func Generate(length int, mode pb.PasswordMode) (out string, err error) {

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)

	var words string
	switch mode {
	case pb.PasswordMode_Easy:
		words = letters + numbers
	case pb.PasswordMode_Medium:
		words = letters + numbers + specialChars
	case pb.PasswordMode_Complex:
		var content []string
		content, err = readWordFile(WordsFilename)
		if err != nil {
			return
		}
		for c := 0; c < length; c++ {
			w := content[rand.Intn(len(content))%len(content)] + specialChars + numbers
			b[c] = w[rand.Intn((len(w)+length)%len(w))]
		}
		out = string(b)
		return
	}

	i := 0
	for ; i < length; i++ {
		b[i] = words[rand.Intn(len(words)+1)%len(words)]
	}
	out = string(b)
	return

}
