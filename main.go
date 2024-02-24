package main

import (
	"github.com/pytool/ssh"
	"fmt"
	"flag"
	"os"
)
var (
  host string
	port     string
	username string
	password string
	remotedir string
	local string
)
func init() {
	flag.StringVar(&username, "u", "root", "username")
	flag.StringVar(&password, "pwd", "", "password")
	flag.StringVar(&host, "h", "127.0.0.1", "127.0.0.1")
	flag.StringVar(&port, "p", "22", "ssh port")
	flag.StringVar(&local, "l", "/local/", "local")
	flag.StringVar(&remotedir, "r", "/remotedir/", "remotedir")
	flag.Parse()
	if password == "" {
	    flag.Usage()
	    os.Exit(-1)
	}
}
func main() {
	client, err := ssh.NewClient(host, port, username, password)
	if err != nil {
	    fmt.Printf("Uptime: %s\n", err)
		panic(err)
	}
	defer client.Close()
	client.Upload(local, remotedir)
}
