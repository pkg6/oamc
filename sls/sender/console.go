package sender

import (
	"fmt"
	"github.com/pkg6/oamc/sls/listener"
	"os"
)

type Console struct {
	l      *listener.Listener
	config any
}

func (c *Console) SetListener(listener *listener.Listener) {
	c.l = listener
}
func (c *Console) SetConfig(config any) {
	c.config = config
}
func (c *Console) Output() error {
	_, err := fmt.Fprint(os.Stdout, c.l.Strings())
	return err
}
