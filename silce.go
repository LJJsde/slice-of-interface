package slice

import (
	"errors"
	"reflect"
	"sync"
)

var (
	ErrTypeDismatch = errors.New("type of the variable is difference")
	ErrEmpty        = errors.New("the slice is empty")
)

type Slice struct {
	Data []interface{}
	size int
	sync.RWMutex
}

func (e *Slice) InitSlice() *Slice {
	q := new(Slice)
	q.Data = nil
	q.size = 0
	return q
}

//插入和检测数据类型
func (e *Slice) PushHeadAndDoFilter(parma []interface{}) error {
	if e.Data == nil {
		e.Data = append(e.Data, 0)
		index := 2
		copy(e.Data[index+1:], e.Data[index:])
		e.Data[index] = parma
	} else {
		DataTypeIn := reflect.TypeOf(e.Data[0:])
		DataTypeOut := reflect.TypeOf(parma)
		if DataTypeIn == DataTypeOut {
			e.Data = append(e.Data, 0)
			index := 2
			copy(e.Data[index+1:], e.Data[index:])
			e.Data[index] = parma
		} else {
			return ErrTypeDismatch
		}
	}
	e.size++
	return nil
}

func (e *Slice) Pop() error {

	if e.Data == nil {
		return ErrEmpty
	} else {
		e.RLock()
		defer e.RUnlock()
		e.Data = append(e.Data[:0], e.Data[1:]...)
		e.Data = e.Data[:len(e.Data)-1]
	}
	e.size--
	return nil
}

func (e *Slice) Top() interface{} {
	return e.Data[0:]
}

func (e *Slice) Length() int {
	return e.size
}
