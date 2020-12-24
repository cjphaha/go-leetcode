package main

import "fmt"

type Quene struct {
	Front int
	End int
	Node [MAX_SIZE]*Tree
}

//判断队列是否为空
func EmptyQueue(q *Quene) bool{
	if q.Front == q.End{
		return true
	}else{
		return false
	}
}

//进入队列
func EnQueue(q *Quene,t *Tree) bool{
	if q.Front == MAX_SIZE - 1{
		return false
	}
	q.Front ++
	q.Node[q.Front] = t
	return true
}

//出队列
func DeQueue(q *Quene,t *Tree) bool{
	if q.Front == q.End{
		return false
	}
	q.End ++
	*t = *q.Node[q.End]
	return true
}

func InitQueue(q *Quene){
	q.Front = -1
	q.End = -1
}

func main1(){
	var q  Quene
	EmptyQueue(&q)
	tree := Tree{
		Value: 1,
		Left: nil,
		Right: nil,
	}
	tree2 := Tree{}
	EnQueue(&q,&tree)
	DeQueue(&q,&tree2)
	fmt.Println(tree2.Value)
}