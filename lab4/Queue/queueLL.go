package queue

type QueueLL struct {
	ll LinkedList
}

type Node struct {
	Data        int
	NextNode    *Node
	PreviusNode *Node
}

type LinkedList struct {
	Head  *Node
	Tale  *Node
	Lenth int
}

func (ll *LinkedList) Push(num int) {
	if ll.Lenth == 0 {
		node := &Node{
			Data:        num,
			NextNode:    nil,
			PreviusNode: nil,
		}
		ll.Head = node
		ll.Tale = node
		ll.Lenth++
	} else if ll.Lenth == 1 {
		node := &Node{
			Data:        num,
			NextNode:    nil,
			PreviusNode: ll.Head,
		}
		ll.Tale = node
		ll.Head.NextNode = node
		ll.Lenth++
	} else {
		node := &Node{
			Data:        num,
			PreviusNode: ll.Tale,
			NextNode:    nil,
		}
		ll.Tale.NextNode = node
		ll.Tale = node
		ll.Lenth++
	}

}

func (ll *LinkedList) Get() int {
	if ll.Lenth < 1 {
		ll.Lenth--
		return 0
	}

	num := ll.Head.Data
	newHead := ll.Head.NextNode
	ll.Head = nil
	ll.Head = newHead
	ll.Lenth--
	return num
}

func (ql *QueueLL) Push(num int) {
	ql.ll.Push(num)
}

func (ql *QueueLL) Get() int {
	if ql.ll.Lenth < 0 {
		return 0
	}
	return ql.ll.Get()
}

func NewQueueLL() *QueueLL {
	return &QueueLL{
		ll: LinkedList{},
	}
}
