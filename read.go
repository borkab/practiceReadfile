package main

//practicing after https://gobyexample.com/reading-files
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	file, err := os.Create("bee.txt")
	check(err)
	defer file.Close()
	fmt.Println("bee.txt was successfully created")

	_, err = file.WriteString("Here is the beehive but where are all the bees? Hiding away where nobody sees. Here they come flying out of their hive: one, two, three, four, five :)")
	check(err)
	fmt.Println("You have successfully written to yor bee.txt")

	//we slurp the entire contents of the file into the memory
	dat, err := os.ReadFile("bee.txt")
	check(err)
	fmt.Print(string(dat))

	//we open our file
	file, err = os.Open("bee.txt")
	check(err)

	//read some bytes from the beginning of the file
	b1 := make([]byte, 47)
	n1, err := file.Read(b1)
	check(err)
	fmt.Printf("\n%d bytes: %s\n", n1, string(b1[:n1]))

	//seek to a known location in the file and Read from there
	o2, err := file.Seek(48, 0)
	check(err)
	b2 := make([]byte, 30)
	n2, err := file.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	//the same as above but more robustly with the io.ReadAtLeast()
	o3, err := file.Seek(78, 0)
	check(err)
	b3 := make([]byte, 41)
	n3, err := io.ReadAtLeast(file, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	//there is no built-in rewind, but Seek() accomplishes this
	_, err = file.Seek(0, 0)
	check(err)

	//the bufio pkg implements a buffered reader
	r4 := bufio.NewReader(file)
	b4, err := r4.Peek(4)
	check(err)
	fmt.Printf("4 bytes: %s\n", string(b4))

	////read file into a string
	content, err := os.ReadFile("bee.txt")
	check(err)
	fmt.Println(string(content))

	//read file line by line

}
