package main

import (
	"container/heap"
  "bufio"
  "os"
  "fmt"
  "strconv"
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

  if st.queue != nil {
	  st.queue.Prev = nil
  }
	heap.Remove(&(st.priori), head.Index)
	return head.Value
}


func (st MaxStack) Top() int {
	return st.queue.Value
}


func (st *MaxStack) PeekMax() int {
	return st.priori[0].Value
}


func (st *MaxStack) PopMax() int {
	max := heap.Pop(&(st.priori))
	node := max.(*DLinkNode)

	if node.Prev == nil {
		st.queue = node.Next
    if st.queue != nil {
      st.queue.Prev = nil
    }
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
  reader := bufio.NewReader(os.Stdin)

  L:
  for true {
    fmt.Println("What would you like to do?")
    selection, err := reader.ReadBytes('\n')

    selection = selection[:len(selection)-1]
    if err != nil {
      fmt.Println("Encountered an error, stopping")
      break
    }

    switch string(selection) {
      case "Push":
        fmt.Println("Enter value to push")
        value, err := reader.ReadBytes('\n')
        if err != nil {
          fmt.Println("Encountered an error, stopping")
          break L
        }
        value = value[:len(value)-1]
        toInt, err := strconv.Atoi(string(value))
        if err != nil {
          fmt.Println("Encountered an error, stopping")
          break L
        }

        maxStack.Push(toInt)
      case "Pop":
        fmt.Printf("You popped %d\n", maxStack.Pop())
      case "PopMax":
        fmt.Printf("MaxPopped %d\n", maxStack.PopMax())
      case "Top":
        fmt.Printf("Top %d\n", maxStack.Top())
      case "PeekMax":
        fmt.Printf("Max right now %d\n", maxStack.PeekMax())
      case "q":
        break L  
      default:
        fmt.Println("Gave a non known command, try again")
    }
  }

 
}
