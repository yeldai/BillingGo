package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func getInput(question string, r *bufio.Reader) (string, error){
	fmt.Print(question)
	input, err:= r.ReadString('\n')

	return strings.TrimSpace(input), err
}


func addBill() bill{

	reader:=bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _:= reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name ,_ := getInput("Create a new bill name:", reader)

	b:= createBill(name)
	fmt.Println("Created the bill-",b.name)
	return b

}

func questOptions( b bill){
	reader := bufio.NewReader(os.Stdin)
	opt ,_ := getInput("Choose option keyword which satisfies your need (a - add item, s - save bill, t- add tip): ", reader)
	
	switch opt{
	case "a":
	name ,_ := getInput("Item name: ", reader)
	price ,_ := getInput("Item price: ", reader)
	p, err := strconv.ParseFloat(price,64)
	if err !=nil{
		fmt.Println("The price must be a number")
		questOptions(b)
	}
	b.addItem(name,p)
	fmt.Println("item added ",name,price)
	questOptions(b)
	
	case "t":
	tip, _ := getInput("Enter tip amount:" ,reader)
	t, err := strconv.ParseFloat(tip,64)
	if err !=nil{
		fmt.Println("The tip must be a number")
		questOptions(b)
	}
	b.updateTip(t)
	questOptions(b)
	case "s":
		b.save()
		fmt.Println("you chose to save the bill- End of transactions ",b.name)
	default:
		fmt.Println("Please choose from the available options")
		questOptions(b)
	}


}

func main() {
	mybill:= addBill()
	questOptions(mybill)


}