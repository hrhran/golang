package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
    i := 0
    for i < len(students) {
        if students[0] == sandwiches[0] {
            sandwiches = sandwiches[1:]
            students = students[1:]
            i = 0
        } else {
            students = append(students, students[0])
            students = students[1:]
            i++
        }
    }
    return len(students)
}

func main() {
	a := []int{1,1,0,0}
	b := []int{0,1,0,1}
	fmt.Println(countStudents(a, b))
}
