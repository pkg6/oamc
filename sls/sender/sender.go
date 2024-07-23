package sender

import (
	"github.com/pkg6/oamc/sls/listener"
	"golang.org/x/sync/errgroup"
	"log"
)

var definitionSenders = map[string]ISender{
	"console": &Console{},
}

type ISender interface {
	SetListener(listener *listener.Listener)
	SetConfig(config any)
	Output() error
}

type Sender struct {
	listener *listener.Listener
	is       []ISender
	ns       []string
	mis      map[string]ISender
}

func New(listener *listener.Listener) *Sender {
	s := Sender{
		listener: listener,
		mis:      make(map[string]ISender),
	}
	for name, config := range listener.Config.Sender {
		if isd, ok := definitionSenders[name]; ok {
			isd.SetListener(s.listener)
			isd.SetConfig(config)
			s.Add(name, isd)
		}
	}
	return &s
}

func (s *Sender) Add(name string, sender ISender) {
	s.is = append(s.is, sender)
	s.ns = append(s.ns, name)
	s.mis[name] = sender
}

func (s *Sender) Run() {
	var (
		g errgroup.Group
	)
	for _, name := range s.ns {
		if sd, ok := s.mis[name]; ok {
			g.Go(func() error {
				return sd.Output()
			})
		}
	}
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
