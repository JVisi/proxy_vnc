package proxy

import "testing"

func TestProxy(t *testing.T) {
	//create default session if required
	t.Skip("this isn't an automated test, just an entrypoint for debugging")

	proxy := &VncProxy{
		WsListeningURL:  "http://0.0.0.0:7778/", // empty = not listening on ws            // empty = no recording
		TCPListeningURL: ":5904",
		//RecordingDir:          "C:\\vncRec", // empty = no recording
		ProxyVncPassword: "1234", //empty = no auth
		SingleSession: &VncSession{
			TargetPassword: "123456",
			ID:             "dummySession",
			Status:         SessionStatusInit,
			Type:           SessionTypeRecordingProxy,
		}, // to be used when not using sessions
		UsingSessions: false, //false = single session - defined in the var above
	}

	proxy.StartListening()
}
