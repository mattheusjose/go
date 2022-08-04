package main

import "fmt"

type Queue struct {
	data   []int
	length int
	rear   int
	front  int
}

func Enqueue(queue *Queue, element int) {

	if IsFull(queue) {
		panic("Queue is full!")
	}

	position := queue.rear % queue.length
	queue.data[position] = element
	queue.rear += 1
}

func Dequeue(queue *Queue) {

	position := queue.front % queue.length
	queue.data[position] = 0
	queue.front += 1
}

func PrintQueue(queue *Queue) {

	for index := queue.front; index < queue.rear; index++ {
		position := index % queue.length
		element := queue.data[position]
		fmt.Print(fmt.Sprintf("[%v]", element))
	}
}

func IsFull(queue *Queue) bool {

	diff := queue.rear - queue.front
	isFull := diff == queue.length
	return isFull
}

func main() {

	length := 10
	queue := Queue{
		data:   make([]int, length),
		front:  0,
		length: length,
		rear:   0,
	}

	Enqueue(&queue, 10)
	Enqueue(&queue, 20)
	Dequeue(&queue)
	Dequeue(&queue)
	Enqueue(&queue, 30)
	Enqueue(&queue, 40)
	Enqueue(&queue, 50)
	Enqueue(&queue, 60)
	Dequeue(&queue)
	Dequeue(&queue)
	Enqueue(&queue, 70)
	Enqueue(&queue, 80)
	Enqueue(&queue, 90)
	Enqueue(&queue, 100)
	Enqueue(&queue, 100)

	PrintQueue(&queue)
}
