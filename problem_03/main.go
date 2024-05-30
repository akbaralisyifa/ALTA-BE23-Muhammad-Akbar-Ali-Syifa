package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	total := 0;

	for _, score := range s.score {
		total += score
	}

	if len(s.score) == 0 {
		return 0
	}

	return float64(total) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	min = s.score[0]
	name = s.name[0]

	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	max = s.score[0]
	name = s.name[0]

	for i, score := range s.score{
		if score > max {
			max = score
			name = s.name[i]
		}	
	}

	return max, name
}

func main() {
	var a = Student{}
	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input : " + string(i) + " Student's Name :")
		fmt.Scanln(&name)
		a.name = append(a.name, name)

		var score int;
		fmt.Println("Input : " + string(i) + " Student's Score :")
		fmt.Scanln(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAvarage Score Student is : ", a.Average())
	scoreMax, nameMax := a.Max();
	fmt.Println("Max Score Student is :  "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min();
	fmt.Println("Min Score Student is : "+nameMin+" (", scoreMin, ")")

}