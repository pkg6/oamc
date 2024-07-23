package data

import "fmt"

type Data struct {
	Format string
	Data   []any
}

func (d *Data) String() string {
	return fmt.Sprintf(d.Format, d.Data...)
}
