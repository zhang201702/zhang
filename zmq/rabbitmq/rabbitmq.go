package rabbitmq

import (
	"errors"
	"github.com/streadway/amqp"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zlog"
	"time"
)

type ConfigInfo struct {
	Addr     string
	UserName string
	Password string
	Port     string
}

type MessageHandler struct {
	QueueName  string
	Exchange   string
	Key        string
	Handler    func(exchange string, body []byte)
	AutoDelete bool
	RetryCount int
}
type RabbitMQ struct {
	Url        string
	Conn       *amqp.Connection
	Handlers   []*MessageHandler
	RetryCount int
}

func Default() (*RabbitMQ, error) {
	info := zconfig.Get("rabbitMQ")
	if info == nil {
		return nil, errors.New("未找到默认的配置项[rabbitMQ]")
	}
	rabbitMQInfo := ConfigInfo{}
	zconfig.Conf.GetStruct("rabbitMQ", &rabbitMQInfo)
	//rabbitMQInfo := ConfigInfo{
	//	Addr: gInfo.GetString("addr"),
	//}
	url := z.String("amqp://", rabbitMQInfo.UserName, ":", rabbitMQInfo.Password, "@", rabbitMQInfo.Addr, ":", rabbitMQInfo.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		zlog.LogError(err, "RabbitMQ Default()")
		return nil, err
	}
	mq := RabbitMQ{
		Url:      url,
		Conn:     conn,
		Handlers: make([]*MessageHandler, 0),
	}
	go func() {
		ticker := time.NewTicker(time.Second * 30)
		for range ticker.C {
			if mq.Conn.IsClosed() {
				zlog.Log("receive.mq", "RabbitMQ.Connection is closed!!!!!!!!!!!!")
				mq.reload()
			}
		}
	}()
	return &mq, nil
}

func (mq *RabbitMQ) reload() {
	mq.RetryCount++
	if conn, err := amqp.Dial(mq.Url); err == nil {
		mq.Conn = conn
		for i := range mq.Handlers {
			mq.Handlers[i].doConsume(mq.Conn)
		}
	} else {
		zlog.LogError(err, "RabbitMQ reload")
	}
}

func (mh *MessageHandler) doConsume(conn *amqp.Connection) error {
	if conn.IsClosed() {
		return errors.New("conn is closed")
	}
	if mh.RetryCount > 3 {
		return errors.New("retry more than 3 times")
	}
	ch, err := conn.Channel()

	if err != nil {
		zlog.LogError(err.(error), "rabbitmq", mh.Exchange, "new Channel")
		return err
	}

	queueName := z.String("go.", mh.Exchange)
	if mh.Key != "" {
		queueName += z.String(".", mh.Key)
	}
	mh.QueueName = queueName
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if mh.Exchange != "" {
		err = ch.ExchangeDeclarePassive(
			mh.Exchange,   // name
			"fanout",      // type
			true,          // durable
			mh.AutoDelete, // auto-deleted
			false,         // internal
			false,         // no-wait
			nil,           // arguments
		)
		if err != nil {
			zlog.Error(err, "rabbitmq.doConsume.ExchangeDeclarePassive 异常", mh.Exchange)
			return err
		}
		err = ch.QueueBind(q.Name, mh.Key, mh.Exchange, false, nil)
	}

	msgs, err := ch.Consume(q.Name, q.Name,
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	zlog.Log("rabbitmq", mh.Exchange, "start listen!!! retry count", mh.RetryCount)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zlog.Log(err.(error), "rabbitmq", mh.Exchange, "message listen")
			}
			mh.RetryCount++
			mh.doConsume(conn)
		}()
		for d := range msgs {
			mh.exec(d.Body)
			d.Ack(false)
		}

	}()
	return err
}

func (mh *MessageHandler) exec(body []byte) {
	//start := time.Now().Unix()
	defer func() {
		//zlog.Log("rabbitmq", mh.Exchange, "start", start, "end", time.Now().Unix(), "spend", time.Now().Unix()-start)
		if err := recover(); err != nil {
			zlog.LogError(err.(error), "rabbitmq", mh.Exchange)
		}
	}()
	mh.Handler(mh.Exchange, body)

}

func (mq *RabbitMQ) Consume(exchange, key string, handler func(exchange string, body []byte)) error {
	mh := &MessageHandler{Exchange: exchange, Key: key, Handler: handler}
	mq.Handlers = append(mq.Handlers, mh)
	return mh.doConsume(mq.Conn)
}

func (mq *RabbitMQ) ConsumeAutoDelete(exchange, key string, handler func(exchange string, body []byte)) error {
	mh := &MessageHandler{Exchange: exchange, Key: key, Handler: handler, AutoDelete: true}
	mq.Handlers = append(mq.Handlers, mh)
	return mh.doConsume(mq.Conn)
}
