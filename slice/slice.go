// slice is a lib package

package slice

import "fmt"

// if we want to use Inspect function outside the slice package, then we need to export it
// making the first letter as uppercase would export

func Inspect(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()

}
