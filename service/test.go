package service

import (
	"context"
)

type Test struct {
}

func (t Test) Test(ctx context.Context) (err error) {
	return nil
}
