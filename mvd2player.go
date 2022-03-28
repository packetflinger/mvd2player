package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

type Config struct {
	Baseq2 string `json:"baseq2folder"`
	Q2exe  string `json:"q2binary"`
}

func main() {
	user, err := user.Current()
	iferr(err)

	sep := string(os.PathSeparator)
	configfile := fmt.Sprintf("%s%smvd2player.json", user.HomeDir, sep)

	configbody, err := os.ReadFile(configfile)
	if err != nil {
		fmt.Printf("Problems loading config file: %s\n", configfile)
		fmt.Println("Please make sure you create that file and add your config like:")
		fmt.Printf("{\n  \"baseq2folder\": \"c:/q2/baseq2\",\n")
		fmt.Printf("  \"q2binary\": \"c:/q2/q2pro.exe\"\n}\n")
		fmt.Printf("\n\nPress ENTER key to quit\n")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	var config Config
	err = json.Unmarshal(configbody, &config)
	iferr(err)

	// wasn't given demo file as argument
	if len(os.Args) < 2 {
		fmt.Printf("No .mvd2 file supplied as argument. Set this program as the default application for .mvd2 files.\n")
		fmt.Printf("\n\nPress ENTER key to quit\n")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	// copy the demo the right place
	demoname := fmt.Sprintf("%s%s%s%stempdemo.mvd2", config.Baseq2, sep, "demos", sep)

	// "copy" demo to a temp file in the right location
	demosrc, err := os.ReadFile(os.Args[1])
	iferr(err)
	err = os.WriteFile(demoname, demosrc, 0666)
	iferr(err)

	// make a temporary config and write it to baseq2 folder
	cfg := "alias loopdemo \"disconnect; mvdplay tempdemo; set nextserver loopdemo\"; loopdemo"
	cfgname := fmt.Sprintf("%s%stempmvd.cfg", config.Baseq2, sep)
	err = os.WriteFile(cfgname, []byte(cfg), 0666)
	iferr(err)

	// spawn a q2pro process to start playing the demo, block until completed
	cmd := exec.Command(config.Q2exe, "+exec", "tempmvd.cfg")
	_ = cmd.Run()

	// remove temp demo and config
	err = os.Remove(demoname)
	iferr(err)

	err = os.Remove(cfgname)
	iferr(err)
}

func iferr(e error) {
	if e != nil {
		log.Println(e)
		fmt.Printf("\n\nPress ENTER key to quit\n")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}
}
