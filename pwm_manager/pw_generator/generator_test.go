package pw_generator

import (
	"testing"

	pb "github.com/t4ke0/pwm/pwm_manager/proto"
)

func init() {
	WordsFilename = "words.txt"
}

const LENGTH int = 20

func TestGenerateEasyPw(t *testing.T) {
	p, err := Generate(LENGTH, pb.PasswordMode_Easy)
	if err != nil {
		t.Logf("failed to generate easy password [%v]", err)
		t.Fail()
		return
	}
	t.Logf("[+] generated password %s", p)
}

func TestGenerateMediumPw(t *testing.T) {
	p, err := Generate(LENGTH, pb.PasswordMode_Medium)
	if err != nil {
		t.Logf("failed to generate medium password [%v]", err)
		t.Fail()
		return
	}
	t.Logf("[+] generated password %s", p)
}

func TestGenerateComplexPw(t *testing.T) {
	p, err := Generate(LENGTH, pb.PasswordMode_Complex)
	if err != nil {
		t.Logf("failed to generate complex password [%v]", err)
		t.Fail()
		return
	}
	t.Logf("[+] generated password %s", p)
}
