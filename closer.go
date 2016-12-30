package closer

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func Do(c io.Closer, msg ...error) {
	err := c.Close()
	if err == nil {
		return
	}
	for _, m := range msg {
		fmt.Fprintf(os.Stderr, "%+v\n", m)
	}
	fmt.Fprintf(os.Stderr, "%+v\n", err)
}

func WithStackStrace() error {
	return errors.New("Error closing in defer")
}
