package linklist

import "fmt"

type LinklistElement struct {
	next *LinklistElement
	previous *LinklistElement
	Data interface{}
}

func (e *LinklistElement) IsEmpty() bool  {
	return e.Data == nil
}
func (e *LinklistElement) Top() LinklistElement  {
	return LinklistElement{Data: e.Data}
}

func (e *LinklistElement) AddBack(item interface{}) {
	if e.Data == nil {
		e.Data = item
		return
	}
	curr := e
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &LinklistElement{next: nil, previous: curr, Data: item}
}

func (e *LinklistElement) PrintElements()  {
	if e.Data == nil {
		fmt.Println("Empty Linklist!")
		return
	}
	curr := e
	fmt.Print("Elements in the linklist: ")
	for true {
		if (curr.next != nil) {
			fmt.Print(curr.Data)
			fmt.Print(" ")
			curr = curr.next
		} else {
			fmt.Print(curr.Data)
			fmt.Println()
			break
		}
	}
}
func (e *LinklistElement) PopTop() {
	if e.next == nil{
		e.Data= nil
	}else if e.next.next != nil {
		e.Data = e.next.Data
		e.next.next.previous = e
		e.next = e.next.next
	}  else {
		e.Data = e.next.Data
		e.next = nil
	}
}

func MakeLinklist() LinklistElement {
	return LinklistElement{}
}

