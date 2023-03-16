package main

import "fmt"

type Component interface {
	Operation()
}

type Composite struct {
	children []Component
}

func (c *Composite) Operation() {
	fmt.Println("Composite operation")
	for _, child := range c.children {
		child.Operation()
	}
}

type Leaf struct {
	data string
}

func (l *Leaf) Operation() {
	fmt.Println("Leaf operation -", l.data)
}

func main() {
	root := &Composite{
		children: []Component{
			&Leaf{data: "file1.txt"},
			&Composite{
				children: []Component{
					&Leaf{data: "file2.txt"},
					&Composite{
						children: []Component{
							&Leaf{data: "file3.txt"},
						},
					},
				},
			},
		},
	}
	root.Operation()
}
