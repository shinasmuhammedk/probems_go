package linkedlist

type Node struct{
    data int
    Next *Node
}

type linkedlist struct{
    head *Node
}

func (l *linkedlist) Reverse(){
    var prev *Node = nil
    curr := l.head
    
    for curr != nil{
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    l.head = prev
}

func (l *linkedlist) Insert(data int){
    newNode := &Node{data: data}
    
    if l.head == nil{
        l.head = newNode
        return
    }
    curr := l.head
    for curr != nil{
        curr = curr.Next
    }
    curr.Next = newNode
}