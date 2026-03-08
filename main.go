package main

import (
	"fmt"
	"os"
)


func main() {
	var name string
	var choice string


	fmt.Println("ismingizni kiriting:")
	fmt.Scanln(&name)

	file, _ := os.OpenFile("ismlar.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if file == nil{
		fmt.Println("error")
		return
	}
	defer file.Close()

	file.WriteString(name + "\n")
	fmt.Println("ism royxatga qoshildi!")

	fmt.Println("jami royxatni korasizmi? (y/n)")
	fmt.Scanln(&choice)

	if choice == "y" {
		data, err := os.ReadFile("ismlar.txt")
		if err != nil {
			fmt.Println("error!")
			return
		}

		fmt.Println("jami ismlar royxati: ")
		fmt.Println(string(data))

	} else {
		fmt.Println("dastur tugadi!")
	} 

	// fmt.Printf("salom, %s", name)
	
}
