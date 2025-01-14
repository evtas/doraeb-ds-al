package main

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	Head  *Node
	Count int
}

func New() *LinkedList {
	return &LinkedList{
		Head:  new(Node),
		Count: 0,
	}
}

func (ll *LinkedList) Insert(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
	}
	newNode.next = ll.Head.next
	ll.Head.next = newNode
	ll.Count += 1
}

func (ll *LinkedList) Delete(val int) {
	pre := ll.Head
	cur := pre.next
	for cur != nil {
		if cur.val == val {
			pre.next = cur.next
			break
		}
		pre = cur
		cur = cur.next
	}
	ll.Count -= 1
}

func (ll *LinkedList) Update(old, new int) {
	cur := ll.Head.next
	for cur != nil {
		if cur.val == old {
			cur.val = new
			break
		}
		cur = cur.next
	}
}

// Reverse 单链表翻转
func (ll *LinkedList) Reverse() {
	if ll.Count <= 1 {
		return
	}
	cur := ll.Head.next
	ll.Head.next = nil
	for cur != nil {
		pre := cur
		cur = cur.next

		pre.next = ll.Head.next
		ll.Head.next = pre
	}
}

func (ll *LinkedList) Print() {
	cur := ll.Head.next
	for cur != nil {
		println(cur.val)
		cur = cur.next
	}
}

func main() {
	ll := New()
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Print()
	ll.Delete(3)
	ll.Update(2, 10)
	ll.Print()
}
