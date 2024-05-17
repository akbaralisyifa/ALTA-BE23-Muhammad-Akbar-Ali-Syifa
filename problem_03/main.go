package main

import "fmt";

func main(){
	var name string;
	fmt.Println("Input your name: ");
	fmt.Scanln(&name);

	fmt.Println("Hello", name + "! Saya Golang bahasa yang menyenangkan.");

}