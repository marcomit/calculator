package util

type Iterator[T any] struct {
	list  []T
	index int
}

func NewIterator[T any](list []T) *Iterator[T] {
	return &Iterator[T]{list: list, index: 0}
}

func (iter *Iterator[T]) HasNext() bool {
	return iter.index < len(iter.list)
}

func (iter *Iterator[T]) HasPrev() bool {
	return iter.index > 0
}

func (iter *Iterator[T]) Check() bool {
	return iter.index >= 0 && iter.index < len(iter.list)
}

func (iter *Iterator[T]) checkIndex(i int) {
	length := len(iter.list)
	if i < 0 || i >= length {
		panic("Iterator out of range")
	}
}

func (iter *Iterator[T]) GetNext() T {
	return iter.list[iter.index+1]
}

func (iter *Iterator[T]) GetPrev() T {
	return iter.list[iter.index-1]
}

func (iter *Iterator[T]) Next() T {
	iter.checkIndex(iter.index)
	value := iter.list[iter.index]
	iter.index++
	return value
}

func (iter *Iterator[T]) Prev() T {
	value := iter.Peek()
	iter.index--
	return value
}
func (iter *Iterator[T]) Peek() T {
	return iter.list[iter.index]
}
