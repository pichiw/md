package md

import "fmt"

var spanClasses = map[int]string{
	1:  "span-1",
	2:  "span-2",
	3:  "span-3",
	4:  "span-4",
	5:  "span-5",
	6:  "span-6",
	7:  "span-7",
	8:  "span-8",
	9:  "span-9",
	10: "span-10",
	11: "span-11",
	12: "span-12",
}

func Span(class string, span int) string {
	if span < 1 || span > 12 {
		panic("invalid span")
	}
	return fmt.Sprintf("%s--%s", class, spanClasses[span])
}
