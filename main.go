package main

import (
	"3DBinPacking/balram/binpacking"
	"fmt"
	"log"
)

type goods [4]int

func (g goods) GetWidth() int {
	return g[1]
}

func (g goods) GetHeight() int {
	return g[2]
}

func (g goods) GetDepth() int {
	return g[3]
}

func (g goods) GetWeight() int {
	return 10
}

func main() {
	items := []binpacking.Item{
		goods{1, 20, 100, 30},
		goods{2, 100, 20, 30},
		goods{3, 20, 100, 30},
		goods{4, 100, 20, 30},
		goods{5, 100, 20, 30},
		goods{6, 100, 100, 30},
		goods{7, 100, 100, 30},
	}
	want := []binpacking.Box{binpacking.BoxSamples[0]}

	want[0].Items = []binpacking.BoxItem{
		{Item: items[5], RType: 0, Pos: [3]int{0, 0, 0}},
		{Item: items[6], RType: 0, Pos: [3]int{100, 0, 0}},
		{Item: items[0], RType: 0, Pos: [3]int{200, 0, 0}},
		{Item: items[1], RType: 0, Pos: [3]int{0, 100, 0}},
		{Item: items[2], RType: 1, Pos: [3]int{100, 100, 0}},
		{Item: items[3], RType: 2, Pos: [3]int{200, 100, 0}},
		{Item: items[4], RType: 0, Pos: [3]int{0, 120, 0}},
	}

	got, err := binpacking.Pack(items)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(printBoxes(got))

}

func printBoxes(boxes []binpacking.Box) (r string) {
	for i, box := range boxes {
		r += fmt.Sprintln("box", i, box.Width, box.Height, box.Depth, len(box.Items))
		for i, item := range box.Items {
			r += fmt.Sprintln("  ", i, item)
		}
	}

	return
}
