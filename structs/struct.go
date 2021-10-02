package main

import "fmt"

type Person struct {
	name string
	age  int
	weight int
}

type Student struct {
	Person
	speciality string
	score int
}

func Older(p1, p2 Person) (Person, int) {
	if p1.age>p2.age{
		return p1, p1.age - p2.age
	}
	return p2, p2.age-p1.age
}

func main() {
	var tom Person
	//赋值初始化
	tom.name, tom.age, tom.weight = "Tom", 18, 165
	//字段名初始化
	bob := Person{age: 39, name:"Bob", weight: 180}
	//按照 struct 字段定义顺序初始化
	paul := Person{"Paul", 40, 178}

	tbOlder, tbDiff := Older(tom, bob)
	tbOlder, tpDiff := Older(tom, paul)
	tbOlder, bpDiff := Older(bob, paul)

	fmt.Printf("of %s and %s, %s is older by %d years\n", tom.name, bob.name, tbOlder.name, tbDiff)
	fmt.Printf("of %s and %s, %s is older by %d years\n", tom.name, paul.name, tbOlder.name, tpDiff)
	fmt.Printf("of %s and %s, %s is older by %d years\n", bob.name, paul.name, tbOlder.name, bpDiff)
	fmt.Println("-----------------------------------------")

	mark := Student{Person{"Mark", 25, 167}, "Computer Science", 98}
	fmt.Println("His name is", mark.name)
}
