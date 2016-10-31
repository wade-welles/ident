package idgen

import (
	"github.com/reflect/ident/source"

	"math/rand"
	"sync/atomic"
	"time"
)

var (
	rng = rand.New(rand.NewSource(time.Now().Unix()))
)

type Provider struct {
	Id uint16

	ch     chan struct{}
	ticker *time.Ticker
	seq    int32
}

func (p *Provider) resetSeq() {
	for {
		select {
		case <-p.ticker.C:
			atomic.StoreInt32(&p.seq, 0)
		case <-p.ch:
			return
		}
	}
}

func (p *Provider) Next() string {
	v := atomic.AddInt32(&p.seq, 1)
	return encode(p.Id, uint32(time.Now().Unix()), uint16(v), rng.Uint32())
}

func (p *Provider) Close() error {
	p.ticker.Stop()
	close(p.ch)

	return nil
}

func NewProvider(src source.Source) *Provider {
	p := &Provider{
		Id: source.MustID(src),

		ch:     make(chan struct{}),
		ticker: time.NewTicker(1 * time.Second),
	}
	go p.resetSeq()

	return p
}
