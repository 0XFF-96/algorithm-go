package main

type record struct {
	num int
	min int
}

type MinStack struct {
	records []record
}

func Constructor() MinStack {
	return MinStack{ []record{} }
}

func (this *MinStack) Push(x int)  {
	if len(this.records) == 0 {
		this.records = append(this.records, record{x, x})
	} else {
		prevMin := this.records[len(this.records) - 1].min
		if x < prevMin {
			this.records = append(this.records, record{x, x})
		} else {
			this.records = append(this.records, record{x, prevMin})
		}
	}

}

func (this *MinStack) Pop()  {
	this.records = this.records[:len(this.records) - 1]
}

func (this *MinStack) Top() int {
	return this.records[len(this.records) - 1].num
}

func (this *MinStack) GetMin() int {
	return this.records[len(this.records) - 1].min
}

