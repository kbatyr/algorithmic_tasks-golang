package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = node
		return
	}
	ls := l.Head
	for ls.Next != nil {
		ls = ls.Next
	}
	ls.Next = &NodeL{Data: data}
}

// go-reloaded List- 1
func SortListInsert(l *NodeI, data_ref int) *NodeI {

	//создаем новый узел, помещаем в него data_ref
	n := &NodeI{Data: data_ref}

	// условие для 1го узла
	if data_ref < l.Data {
		n.Next = l
		l = n
		return l
	}

	// присваиваем нов.пер-ой значение l
	iteration := l

	// если след.узел меньше дата_реф, двигаемся к след узлу
	for iteration.Next != nil && iteration.Next.Data < data_ref {
		iteration = iteration.Next
	}
	// как только след узел будет больше дата_реф, вклиниваем наш новый узел посередине ->
	// меняя адрес левого узла, и присваивая нов.узлу адрес след.узла
	n.Next = iteration.Next
	iteration.Next = n

	// возвращаем измененный список с добавленным новым узлом
	return l
}

// go-reloaded List-2
func SortedListMerge(n2 *NodeI, n1 *NodeI) *NodeI {

	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}

	for n2 != nil {
		n1 = SortListInsert(n1, n2.Data)
		n2 = n2.Next
	}
	return n1
}

// go-reloaded List-3
func ListRemoveIf(l *List, data_ref interface{}) {
	i := l.Head

	if i == nil {
		return
	}

	for i.Next != nil {
		if i.Data == data_ref {
			*i = *i.Next
		} else if i.Next.Data == data_ref && i.Next.Next != nil {
			i.Next = i.Next.Next
		} else if i.Next.Data == data_ref && i.Next.Next == nil {
			i.Next = nil
		} else {
			i = i.Next
		}
	}
	if i.Data == data_ref {
		l.Head = nil
	}
}
