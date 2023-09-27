/// Advent of Code 2015 Day 1 Part 1

package main

import ("fmt"
	"os"
	"bufio"
)

func main() {
	fl :=  0; // floor level

	f, _ := os.Open("./aoc_2015-1.dat"); // Santa's directions,
	br := bufio.NewReader(f) // ... assume no Open error

	for {
		c, err := br.ReadByte()
		if err != nil { // EOF
			break
		}
		
		if (c == '(') {
			fl += 1;
		} else if (c == ')') {
			fl -= 1;
		}
	}
    	fmt.Printf("Santa ended up on floor: %d\n", fl);
}
