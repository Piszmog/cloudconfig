package cloudconfigclient

import (
	"fmt"
	"io"
)

func closeResource(r io.Closer) {
	if err := r.Close(); err != nil {
		fmt.Println(fmt.Errorf("cloudconfigclient: failed to close resource: %w", err))
	}
}
