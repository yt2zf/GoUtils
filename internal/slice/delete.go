package slice

import (
	"github.com/yt2zf/GoUtils/internal/errs"
)

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index >= length {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}

	res := src[index]
	for i := index; i < length-1; i++ {
		src[i] = src[i+1]
	}

	src = src[0 : length-1]
	return src, res, nil
}
