package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/TaKeO90/pwm/server/genkey"
	"github.com/TaKeO90/pwm/sqlite"
)

const (
	mainf string = "./server/main.go"
)

func startProcess(args ...string) (p *os.Process, err error) {
	//Make sure the first argument is on out PATH env
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func startServer() (p *os.Process) {
	if p, err := startProcess("go", "run", mainf); err == nil {
		return p
	}
	return nil
}

func initDatabase() {
	db := sqlite.InitDb()
	_, err := sqlite.CreateTables(db)
	if err == nil {
		fmt.Println("Initialized Database")
	} else {
		log.Fatal(err)
	}
}

func main() {
	msg, err := genkey.KeysChecking()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	initDatabase()
	startServer().Wait()
}
