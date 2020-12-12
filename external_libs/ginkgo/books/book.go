package books

import "fmt"

func init() {

}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) CategoryByLength() string {

	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

func DoSomething(ch chan string) {
	ch <- "Done!"
	fmt.Println("parallel")
}
