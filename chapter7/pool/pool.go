package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m			sync.Mutex
	resources	chan io.Closer
	factory 	func() (io.Closer, error)
	closed		bool
}

var ErrPoolClosed = errors.New("풀이 닫혔습니다")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("풀의 크기가 너무 작습니다")
	}
	return &Pool{
		factory:	fn,
		resources:	make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("리소스 획득:", "공유된 리소스")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("리소스 획득:", "새로운 리소스")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("리소스 반환:", "리소스 큐에 반환")
	default:
		log.Println("리소스 반환:", "리소스 해제")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)

	for r:= range p.resources {
		r.Close()
	}
}