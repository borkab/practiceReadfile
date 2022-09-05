package main

import (
	"fmt"
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

}
