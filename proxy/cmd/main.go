package main

import (
	"flag"
	"os"

	"github.com/JVisi/proxy_vnc/logger"
	vncproxy "github.com/JVisi/proxy_vnc/proxy"
)

func main() {
	//create default session if required
	var tcpPort = flag.String("tcpPort", "", "tcp port")
	var wsPort = flag.String("wsPort", "", "websocket port")
	var wsUrl = flag.String("wsUrl", "", "websocket url:port (currently) with the prefix. (e.g.: ws://localhost:6961)")
	var vncPass = flag.String("vncPass", "", "password on incoming vnc connections to the proxy, defaults to no password")
	var targetVnc = flag.String("target", "", "target vnc server (host:port or /path/to/unix.socket)")
	var targetVncPass = flag.String("targPass", "", "target vnc password")
	var logLevel = flag.String("logLevel", "info", "change logging level")

	flag.Parse()
	logger.SetLogLevel(*logLevel)

	if *tcpPort == "" && *wsPort == "" {
		logger.Error("no listening port defined")
		flag.Usage()
		os.Exit(1)
	}

	if *targetVnc == "" {
		logger.Error("no target vnc server host/port or socket defined")
		flag.Usage()
		os.Exit(1)
	}

	if *vncPass == "" {
		logger.Warn("proxy will have no password")
	}

	tcpURL := ""
	if *tcpPort != "" {
		tcpURL = ":" + string(*tcpPort)
	}
	
	proxy := &vncproxy.VncProxy{
		WsListeningURL:   *wsUrl, // empty = not listening on ws
		TCPListeningURL:  tcpURL,
		ProxyVncPassword: *vncPass, //empty = no auth
		SingleSession: &vncproxy.VncSession{
			Target:         *targetVnc, //"localhost:5900
			TargetPassword: *targetVncPass, //"vncPass",
			ID:             "dummySession",
			Status:         vncproxy.SessionStatusInit,
			Type:           vncproxy.SessionTypeProxyPass,
		}, // to be used when not using sessions
		UsingSessions: false, //false = single session - defined in the var above
	}

	proxy.StartListening()
}
