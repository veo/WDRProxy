package main

import (
	"context"
	"fmt"
	"github.com/bogdanovich/dns_resolver"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"time"
)

type handle struct {
	reverseProxy string
}

func checkkeywords(data string)  (bool){
	for _, v := range keywords{
		if strings.Contains(strings.ToLower(data),v) == true{
			return true
		}else {
			return false
		}
	}
	return false
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	normallogFile, err1 := os.OpenFile("normal.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	highlogFile, err2 := os.OpenFile("dangerous.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err1 != nil {
		log.Fatalln("fail to create normal.log file!")
	}
	if err2 != nil {
		log.Fatalln("fail to create high.log file!")
	}
	post, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(strings.NewReader(string(post)))
	log.SetOutput(normallogFile)
	if string(post) != ""{
		log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent() + "  " + string(post))
	}else {
		log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	}
	query ,_ := url.PathUnescape(r.URL.String())
	postdata ,_ := url.PathUnescape(string(post))
	if checkkeywords(query)||checkkeywords(postdata){
		fmt.Fprintln(w,"stop hacking")
		log.SetOutput(highlogFile)
		if string(post) != ""{
			log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent() + "  " + string(post))
		}else {
			log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
		}

	}else {
		remote, err := url.Parse(this.reverseProxy)
		if err != nil {
			log.Fatalln(err)
		}
		dialer := &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}
		http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			remote := strings.Split(addr, ":")
			if cmd.ip == "" {
				resolver := dns_resolver.New([]string{"114.114.114.114", "114.114.115.115", "119.29.29.29", "223.5.5.5", "8.8.8.8", "208.67.222.222", "208.67.220.220"})
				resolver.RetryTimes = 5
				ip, err := resolver.LookupHost(remote[0])
				if err != nil {
					cmd.ip = remote[0]
				}else {
					cmd.ip = ip[0].String()
				}
			}
			addr = cmd.ip + ":" + remote[1]
			return dialer.DialContext(ctx, network, addr)
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		r.Host = remote.Host
		proxy.ServeHTTP(w, r)
	}
}
