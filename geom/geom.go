package main

import "fmt"

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Box struct {
	width, height, depth float64
	color Color
}

type Color byte

type BoxList []Box

func (b Box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestsColor() Color {
	v := 0.0
	k := Color(WHITE)

	for _, b := range bl {
		if b.volume() > v {
			v = b.volume()
			k = b.color
		}
	}

	return k
}

func (bl BoxList) PaintItBlack() {
	for i, _ := range bl {
		bl[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
        Box{10, 10, 1, BLUE},
        Box{10, 30, 1, WHITE},
        Box{20, 20, 20, YELLOW},
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes) -1].color)
	fmt.Println("The biggest one is", boxes.BiggestsColor())

	boxes.PaintItBlack()

	fmt.Println("The color of the second one is", boxes[1].color)
	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestsColor())
}