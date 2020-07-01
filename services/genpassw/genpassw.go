package genpassw

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	alphab string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums          = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	spChar string = "&'(-_à)=~$\"*%/!:;,<>"
)

//PwType interface that holds the methods available to generate a password
type PwType interface {
	GenerateChars() string
	GenerateInts() string
	GenerateSpchars() string
	GenerateComplex() string
}

//Password holds the length of the password
type Password struct {
	length int
}

//New returns Password struct with a specified length
func New(length int) *Password {
	return &Password{length}
}

//GenerateChars generate password based on characters
func (p *Password) GenerateChars() string {
	var genStr []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < p.length; i++ {
		RI := rand.Intn(len(alphab))
		genStr = append(genStr, string(alphab[RI]))
	}
	return strings.Join(genStr, "")
}

//GenerateInts generate password base on integers
func (p *Password) GenerateInts() string {
	var genInt []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < p.length; i++ {
		RI := rand.Intn(len(nums))
		genInt = append(genInt, strconv.Itoa(nums[RI]))
	}
	return strings.Join(genInt, "")
}

//GenerateSpchars generate password based on special character
func (p *Password) GenerateSpchars() string {
	var genspChar []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < p.length; i++ {
		RI := rand.Intn(len(spChar))
		genspChar = append(genspChar, string(spChar[RI]))
	}
	return strings.Join(genspChar, "")
}

//GenerateComplex generate password based on the three type before
// string + special characters + integers
func (p *Password) GenerateComplex() string {
	var genComplx []string
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < p.length; i++ {
		RI, RJ, RZ := rand.Intn(len(nums)), rand.Intn(len(alphab)), rand.Intn(len(spChar))
		genComplx = append(genComplx, string(alphab[RJ]), strconv.Itoa(nums[RI]), string(spChar[RZ]))
	}
	return strings.Join(genComplx, "")
}
