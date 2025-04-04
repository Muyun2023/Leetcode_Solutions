type MinStack struct {
    top *StackNode
    min int
}
type StackNode struct {
	data int
	next *StackNode
    lastmin int
}
func Constructor() MinStack {
    return mystack
}
func (this *MinStack) Push(val int)  {
    if this.top == nil {
        newtop = &StackNode{data: val, next: this.top}
        this.min = val
    } else {
        newtop = &StackNode{data: val, next: this.top, lastmin: this.min}
    }
    this.top = newtop
    if this.top.data < this.min {
        this.min = this.top.data
    }
}
func (this *MinStack) Pop()  {
    if this.top.next == nil {
        this.top = nil
        return
    } 
    this.min = this.top.lastmin
    *this.top = *this.top.next
}
func (this *MinStack) Top() int {
    return this.top.data
}
func (this *MinStack) GetMin() int {
    return this.min; 
}

