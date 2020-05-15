/**
用来内部消息处理的订阅和处理，通过go 通道来实现 ,相当于内部的消息总线
	异步处理消息消费，通过go 的并发处理
	没有持久化功能,程序关闭，所有的订阅和消费都会消失。
	没有消息处理超时处理，消息消费也没有做异常处理。
  没有超时或消费异常，再次消费。
*/
package zmq

type MessageHandler func(data ...interface{})
type messageDefine struct {
	Name    string
	Handler []MessageHandler
	message chan []interface{}
}

var mqMap = make(map[string]messageDefine)

/**
订阅消息处理
*/
func Register(messageName string, handler MessageHandler) {

	if md, ok := mqMap[messageName]; ok {
		md.Handler = append(md.Handler, handler)
	} else {
		md = messageDefine{
			message: make(chan []interface{}),
			Name:    messageName,
			Handler: []MessageHandler{},
		}
		md.start()
		mqMap[messageName] = md
		md.Handler = append(md.Handler, handler)
	}
}

/**
消息源，触发消息
*/
func Trigger(messageName string, data ...interface{}) {
	if md, ok := mqMap[messageName]; ok {
		md.message <- data
	}
}

/**
开始消息监听
*/
func (md *messageDefine) start() {
	go func() {
		for {
			data := <-md.message
			for i := range md.Handler {
				go md.Handler[i](data...)
			}
		}
	}()
}
