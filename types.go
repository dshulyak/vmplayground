package vm

type Uint64 struct {
	value uint64
}

type List[T any] struct {
	head, tail *element[T]
}

func (l *List[T]) Add(value T) {
	e := &element[T]{value: value}
	if l.head == nil {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		l.tail = e
	}
}

type Iterator[T any] struct {
	current *element[T]
}

func (i *Iterator[T]) Next() (T, bool) {
	if i.current == nil {
		var empty T
		return empty, false
	}
	current := i.current
	i.current = current.next
	return current.value, true
}

type element[T any] struct {
	value T
	next  *element[T]
}

type Map[T comparable, V any] struct {
	order List[T]
	value map[T]V
}
