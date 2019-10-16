package znet

import (
	"github.com/gogf/gf/util/gconv"
	"github.com/zhang201702/zhang/zconfig"
	"net"
	"net/http"
	"sync"
	"time"
)

var clientTran = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   20 * time.Second,
		KeepAlive: 10 * time.Second,
	}).DialContext,
	MaxIdleConns:          4,
	IdleConnTimeout:       2 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

var maxLen int = 100
var clientCount = 0
var mutex = sync.Mutex{}
var clients = make([]*http.Client, maxLen)
var isWait = false
var free = make(chan *http.Client)

func fetchHttpClient() *http.Client {
	defer mutex.Unlock()
	client := clients[clientCount]
	if client == nil {
		client = &http.Client{
			Transport: clientTran,
		}
		clients[clientCount] = client
	}
	clientCount++

	return client
}
func GetHttpClient() *http.Client {
	mutex.Lock()
	if clientCount < maxLen {
		return fetchHttpClient()
	} else {
		isWait = true
		for {
			select {
			case <-free:
				isWait = false
				return fetchHttpClient()
			}
		}
	}
}

func FreeHttpClient(client *http.Client) {
	clientCount--
	clients[clientCount] = client
	if isWait {
		free <- client
	}
}

func init() {
	conInfo := zconfig.Default()
	if httpInfo, ok := (*conInfo)["http"]; ok {
		hi := httpInfo.(map[string]interface{})
		if max, ok2 := hi["max_http_client"]; ok2 {
			maxLen = gconv.Int(max)
		}
	}
}
