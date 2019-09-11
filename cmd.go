package main

import "flag"

type Cmd struct {
	bind string
	remote string
	ip string
}
var keywords  = []string{"select","insert","update","drop","delete","dumpfile","outfile","rename","floor","extractvalue","updatexml","name_const"}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.bind, "l", "0.0.0.0:8080", "listen on ip:port")
	flag.StringVar(&cmd.remote, "r", "https://www.xunblog.com", "reverse proxy addr")
	//flag.StringVar(&cmd.ip, "ip", "", "reverse proxy addr server ip")
	flag.Parse()
	return cmd
}
