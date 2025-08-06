package answer

type MyQueue struct {
	inStack  []int
	outStack []int
}

func ConstructorV1() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	this.move()
	top := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return top
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.outStack[len(this.outStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

func (this *MyQueue) move() {
	if len(this.outStack) == 0 {
		for i := len(this.inStack); i > 0; i-- {
			this.outStack = append(this.outStack, this.inStack[i-1])
			this.inStack = this.inStack[:i-1]
		}
	}
}
