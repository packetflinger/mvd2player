package main

import (
	"encoding/json"
	"fmt"
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
	iferr(err)

	var config Config
	err = json.Unmarshal(configbody, &config)
	iferr(err)

	// copy the demo the right place
	demoname := fmt.Sprintf("%s%s%s%stempdemo.mvd2", config.Baseq2, sep, "demos", sep)
	demosrc, err := os.ReadFile(os.Args[1])
	iferr(err)

	err = os.WriteFile(demoname, demosrc, 0666)
	iferr(err)

	// make a temporary config and write it to baseq2 folder
	cfg := "alias loopdemo \"disconnect; mvdplay tempdemo; set nextserver loopdemo\"; loopdemo"
	cfgname := fmt.Sprintf("%s%stempmvd.cfg", config.Baseq2, sep)
	err = os.WriteFile(cfgname, []byte(cfg), 0666)
	iferr(err)

	// spawn a quake 2 process to start playing the demo
	cmd := exec.Command(config.Q2exe, "+exec", "tempmvd.cfg")
	_ = cmd.Run()

	// remove
	err = os.Remove(demoname)
	iferr(err)

	err = os.Remove(cfgname)
	iferr(err)
}

func iferr(e error) {
	if e != nil {
		panic(e)
	}
}
