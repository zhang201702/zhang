package zmq

type MessageHandler func(data ...interface{})
type messageDefine struct {
	Name    string
	Handler []MessageHandler
	message chan []interface{}
}

var mqMap = make(map[string]messageDefine)

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
func Trigger(messageName string, data ...interface{}) {
	if md, ok := mqMap[messageName]; ok {
		md.message <- data
	}
}

func (md *messageDefine) start() {
	go func() {
		for {
			data := <-md.message
			for i := range md.Handler {
				md.Handler[i](data...)
			}
		}
	}()
}
