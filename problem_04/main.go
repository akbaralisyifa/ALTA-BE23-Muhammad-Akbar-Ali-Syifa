package main

import "fmt"

type Student struct {
    Name       string
    NameEncode string
    Score      int
}

type Chiper interface {
    Encode() string
    Decode() string
}

func (s *Student) Encode() string {
    var nameEncode string

    for _, char := range s.Name {
        shifted := char + 3
        nameEncode += string(shifted)
    }
    s.NameEncode = nameEncode
    return nameEncode
}

func (s *Student) Decode() string {
    var nameDecode string

    for _, char := range s.NameEncode {
        shifted := char - 3
        nameDecode += string(shifted)
    }

    return nameDecode
}

func main() {
	var menu int
    var a = Student{}

    var c Chiper = &a
	fmt.Println("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scanln(&menu);
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scanln(&a.Name)
		fmt.Print("\nInput Student's Name "+ a.Name+ " is: " + c.Encode())
	}else if menu == 2{
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scanln(&a.Name)
		fmt.Print("\nInput Student's Name "+ a.NameEncode+ " is: " + c.Decode())
	}else {
		fmt.Println("Wrong input name menu!")
	}

}