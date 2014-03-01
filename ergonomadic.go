package main

import (
	"code.google.com/p/go.crypto/bcrypt"
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/jlatt/ergonomadic/irc"
	"log"
	"os"
)

func genPasswd(passwd string) {
	crypted, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	encoded := base64.StdEncoding.EncodeToString(crypted)
	fmt.Println(encoded)
}

func main() {
	conf := flag.String("conf", "ergonomadic.json", "ergonomadic config file")
	initdb := flag.Bool("initdb", false, "initialize database")
	passwd := flag.String("genpasswd", "", "bcrypt a password")
	flag.Parse()

	if *passwd != "" {
		genPasswd(*passwd)
		return
	}

	config, err := irc.LoadConfig(*conf)
	if err != nil {
		log.Fatal(err)
	}

	if *initdb {
		irc.InitDB(config.Database())
		return
	}

	// TODO move to data structures
	irc.DEBUG_NET = config.Debug["net"]
	irc.DEBUG_CLIENT = config.Debug["client"]
	irc.DEBUG_CHANNEL = config.Debug["channel"]
	irc.DEBUG_SERVER = config.Debug["server"]

	irc.NewServer(config).Run()
}
