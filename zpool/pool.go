package zpool

type ConnRes interface {
	Close() error
}
type Factory func() (ConnRes, error)

type Pool struct {
	conns   chan ConnRes
	factory Factory
}

func NewPool(factory Factory, cap int) *Pool {
	return &Pool{
		conns:   make(chan ConnRes, cap),
		factory: factory,
	}
}

func (p *Pool) new() (ConnRes, error) {
	return p.factory()
}

func (p *Pool) Get() (conn ConnRes) {
	select {
	case conn = <-p.conns:
		{
		}
	default:
		conn, _ = p.new()
	}
	return
}

func (p *Pool) Put(conn ConnRes) {
	select {
	case p.conns <- conn:
		{
		}
	default:
		conn.Close()
	}
}
