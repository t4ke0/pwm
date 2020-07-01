package readcredfile

import (
	"../pwshow"
	"strings"
)

type DumpCreds interface {
	ExtractCreds() CredFileCntList
}

type CredFileCnt struct {
	Username string
	Password string
	Category string
}

type CredFile struct {
	ByteCnt []byte
}

type CredFileCntList []CredFileCnt

func New(filebyte []byte) *CredFile {
	return &CredFile{ByteCnt: filebyte}
}

func (c *CredFile) ExtractCreds() CredFileCntList {
	var ncontent string
	var dumpedCreds CredFileCntList
	content := string(c.ByteCnt)
	credFormat := &CredFileCnt{}
	for _, n := range content {
		if string(n) != "\n" && string(n) != "EOF" {
			ncontent += string(n)
		} else if string(n) == "\n" {
			s := strings.Split(ncontent, ",")
			for i, j := range s {
				switch i {
				case 0:
					credFormat.Username = strings.TrimSpace(j)
				case 1:
					credFormat.Password = strings.TrimSpace(j)
				case 2:
					credFormat.Category = strings.TrimSpace(j)
				}
			}
			dumpedCreds = append(dumpedCreds, *credFormat)
			ncontent = ""
		}
	}
	return dumpedCreds
}

func (cl CredFileCntList) Compare(username string) ([]int, bool) {
	var ok bool
	old := pwshow.ShowCreds(username, "")
	var repeated []int

	for i, n := range cl {
		for _, j := range old {
			if n.Username == j.Username && n.Password == j.Password {
				repeated = append(repeated, i+1)
				ok = true
				if i == len(cl)-1 && len(repeated) != 0 {
					break
				}
			} else if n.Username != j.Username && n.Password != j.Password && len(repeated) == 0 {
				ok = false
			}
		}
	}
	return repeated, ok
}
