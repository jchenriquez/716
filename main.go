package main

import (
	"container/heap"
	"fmt"
)

type DLinkNode struct {
	Value int
	Index int
	Prev *DLinkNode
	Next *DLinkNode
}

type PriorityQueue []*DLinkNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq PriorityQueue) Less (i, j int) bool {
	return pq[i].Value >= pq[j].Value
}

func (pq *PriorityQueue) Push (n interface{}) {
	dLinkNode := n.(*DLinkNode)
	dLinkNode.Index = len(*pq)
	*pq = append(*pq, n.(*DLinkNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	ret := (*pq)[len(*pq)-1]
	(*pq)[len(*pq)-1] = nil
	ret.Index = -1
	*pq = (*pq)[:len(*pq)-1]

	return ret
}

type MaxStack struct {
	queue *DLinkNode
	priori PriorityQueue
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	priorityQ := make(PriorityQueue, 0, 100)
	heap.Init(&(priorityQ))

	return MaxStack{nil, priorityQ}
}


func (st *MaxStack) Push(x int)  {
	node := &DLinkNode{Value:x, Prev:nil, Next:nil}

	if st.queue != nil {
		node.Next = st.queue
		st.queue.Prev = node
		st.queue = node
	} else {
		st.queue = node
	}

	heap.Push(&(st.priori), node)
}


func (st *MaxStack) Pop() int {
	head := st.queue
	st.queue = head.Next
	st.queue.Prev = nil
	heap.Remove(&(st.priori), head.Index)
	return head.Value
}


func (st MaxStack) Top() int {
	return st.queue.Value
}


func (st *MaxStack) PeekMax() int {
	return st.priori[len(st.priori)-1].Value
}


func (st *MaxStack) PopMax() int {
	max := heap.Pop(&(st.priori))
	node := max.(*DLinkNode)

	if node.Prev == nil {
		st.queue = node.Next
	} else {
		prev := node.Prev
		prev.Next = node.Next

		if node.Next != nil {
			node.Next.Prev = prev
		}
	}

	return node.Value
}


/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.PeekMax();
 * param_5 := obj.PopMax();
 */

func main() {

	maxStack := Constructor()

	maxStack.Push(5)
	maxStack.Push(1)
	maxStack.Push(5)

	fmt.Printf("Top %d\n", maxStack.Top())
	fmt.Printf("Popped max %d\n", maxStack.PopMax())
	fmt.Printf("Top %d\n", maxStack.Top())
	fmt.Printf("Peeking max %d\n", maxStack.PeekMax())
	fmt.Printf("Popped %d\n", maxStack.Pop())
	fmt.Printf("Top %d\n", maxStack.Top())

}
