package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/TaKeO90/pwm/server/genkey"
	"github.com/TaKeO90/pwm/sqlite"
)

const (
	mainf        string = "./server/main.go"
	frontEndPath string = "./myfrontend/"
	sslProxy     string = "./ssl-proxy/"
	frontEndL    string = "127.0.0.1:5000"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type chanItems struct {
	p   *os.Process
	err error
}

func startProcess(dir string, args ...string) (p *os.Process, err error) {
	//Make sure the first argument is on out PATH env
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		if dir == "" {
			procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
		} else {
			procAttr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
			procAttr.Dir = dir
		}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func startServer(c chan chanItems, wg *sync.WaitGroup) {
	cI := new(chanItems)
	defer wg.Done()
	p, err := startProcess("", "go", "run", mainf)
	if err != nil {
		cI.err = err
		cI.p = nil
		c <- *cI
	}
	cI.err = nil
	cI.p = p
	c <- *cI
}

func runFrontEnd(c chan chanItems, wg *sync.WaitGroup) {
	cI := new(chanItems)
	defer wg.Done()
	cmd := exec.Command("npm", "npm", "run", "build")
	err := cmd.Run()
	if err != nil {
		cI.err = err
		cI.p = nil
		c <- *cI
	}
	p, err := startProcess(frontEndPath, "npm", "start")
	if err != nil {
		cI.err = err
		cI.p = nil
		c <- *cI
	}
	cI.p = p
	cI.err = nil
	c <- *cI
}

type sslProxyCertFile struct {
	certPath string
	keyPath  string
}

func (s *sslProxyCertFile) runsslProxy() (*os.Process, error) {
	if s.certPath == "" && s.keyPath == "" {
		p, err := startProcess(sslProxy, "go", "run", "main.go", "-from", "localhost:4430", "-to", frontEndL, "-altnames", "localhost")
		if err != nil {
			return nil, err
		}
		return p, nil
	} else if s.certPath != "" && s.keyPath != "" {
		fmt.Println(s.certPath, s.keyPath)
		p, err := startProcess(sslProxy, "go", "run", "main.go", "-cert", s.certPath, "-key", s.keyPath, "-from", "localhost:4430", "-to", frontEndL, "-altnames", "localhost")
		if err != nil {
			return nil, err
		}
		return p, nil
	}
	return nil, nil
}

func runP(c chan chanItems, wg *sync.WaitGroup) {
	//CHECK IF THERE IS ANY CERT FILES IN PROXY DIREC.
	//IF NOT TRY TO RUN THE PROXY AND GENERATE NEW CERT FILES.
	//ELSE IF WE FOUND THESE FILES WE NEED TO ADD THEM AS FLAGS TO RUN THE PROXY.
	cI := new(chanItems)
	sslCF := new(sslProxyCertFile)
	defer wg.Done()
	proxyD, err := ioutil.ReadDir(sslProxy)
	if err != nil {
		cI.err = err
	}
	for _, f := range proxyD {
		if pemF := filepath.Ext(f.Name()); pemF == ".pem" {
			//RUN PROXY WITH CERT & KEY FLAGS
			if f.Name() == "cert.pem" {
				sslCF.certPath = f.Name()
			} else if f.Name() == "key.pem" {
				sslCF.keyPath = f.Name()
			}
		}
	}
	p, err := sslCF.runsslProxy()
	if err != nil {
		cI.err = err
	}
	cI.p = p
	c <- *cI
}

func initDatabase() {
	db := sqlite.InitDb()
	_, err := sqlite.CreateTables(db)
	checkError(err)
}

func run() {
	var Wg sync.WaitGroup
	_, err := genkey.KeysChecking()
	if err != nil {
		if err := os.Mkdir("./services/pwencrypter/keys", 0700); err != nil {
			log.Fatal(err)
		}
		_, err = genkey.KeysChecking()
		checkError(err)
	}
	initDatabase()
	Wg.Add(3)
	c1 := make(chan chanItems, 3)
	go startServer(c1, &Wg)
	go runFrontEnd(c1, &Wg)
	go runP(c1, &Wg)
	Wg.Wait()
	e1 := <-c1
	e2 := <-c1
	e3 := <-c1
	checkError(e1.err)
	checkError(e2.err)
	checkError(e3.err)
	e1.p.Wait()
	e2.p.Wait()
	e3.p.Wait()
}

func main() {
	run()
}
