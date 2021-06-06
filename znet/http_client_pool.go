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
var mutexFree = sync.Mutex{}
var clients []*http.Client = nil
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
	mutexFree.Lock()
	clientCount--
	if len(clients) > clientCount && clientCount > 0 {
		clients[clientCount] = client
	}
	mutexFree.Unlock()
	if isWait {
		free <- client
	}
}

func init() {
	max := zconfig.Conf.GetInt("http.max_http_client")
	if max > 0 {
		maxLen = gconv.Int(max)
	}
	clients = make([]*http.Client, maxLen+2)
}
