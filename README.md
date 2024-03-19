A lightweight RFB proxy that is capable of acting between a VNC server and a client, supporting tcp and websocket connections
* Supports all modern encodings & most useful pseudo-encodings
* Supports multiple VNC client connections & multi servers (chosen by sessionId)
* Supports being a "websockify" proxy (for web clients like NoVnc)
    
- Tested on tight encoding with:
    - Tightvnc (client + go client + java client + server)
    - NoVnc(web client) => use -wsPort to open a websocket
    - ChickenOfTheVnc(client)
    - VineVnc(server)
    - TigerVnc(client)
    - Qemu vnc(server) 


### Executables (see releases)
* proxy - the actual recording proxy, supports listening to tcp & ws ports and recording traffic to fbs files

## Usage:
    proxy -target="192.168.100.100:5901 -targPass=@@@@@ -tcpPort=5903 -wsPort=5905 -wsUrl="ws://localhost" -vncPass=@!@!@!