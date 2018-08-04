package main

import "fmt"

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
	PB = TB * 1024
)

const (
	one int = iota + 1
	two
	three
	// iota can be use for defining incrementing constant
	// starts at 0 and counts up
)

// const (
// 	KB ByteSize = 1 << (10 * (iota + 1))
// 	MB
// 	GB
// 	TB
// 	PB
// )

// The new type extends the float64
// We can use beyond what float64 has to offer
type ByteSize float64

// Check the appropriate byte type name based how the inputted byte
// return the appropriate name
func (b ByteSize) String() string {
	switch {
	case b >= PB:
		return "Too Big!"
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%dB", b)
}

func main() {
	fmt.Println(ByteSize(2048))               // 2.00KB
	fmt.Println(ByteSize(247247242732747.84)) // 224.87TB

	fmt.Println(three) // 3
	// NOTE - fmt package will use the String method we created if it exists on a type
}
