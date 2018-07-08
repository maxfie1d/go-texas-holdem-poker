package main

type poker struct {
	width int
	height int
}

func (r *poker) area() int {
	return r.width * r.height
}
