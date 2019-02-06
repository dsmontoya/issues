package main

import (
	"fmt"
	"os"

	"github.com/dsmontoya/selenium"
	"github.com/dsmontoya/selenium/chrome"
)

func main() {
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	proxy := selenium.Proxy{
		Type:         selenium.Manual,
		SOCKS:        os.Getenv("PROXY_HOST"),
		SocksVersion: 5,
	}
	caps.AddProxy(proxy)
	caps.AddChrome(chrome.Capabilities{
		Args: []string{},
	})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://%s:%s/wd/hub", os.Getenv("SELENIUM_HOST"), os.Getenv("SELENIUM_PORT")))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
}
