package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Grade []int
}

func addStudent(students []Student, newStudent Student) []Student {
	students = append(students, newStudent)
	return students
}

func updateGrades(students []Student, name string, newGrades []int) []Student {
	for i, val := range students {
		if val.Name == name {
			students[i].Grade = append(students[i].Grade, newGrades...)
			//fmt.Println(students[i].Grade)
		}
	}
	return students
}

//func calculateAverage(students []Student, name string) float64 {
//	var averageGrade float64
//	for _, val := range students {
//		if val.Name == name {
//			for _, grade := range val.Grade {
//				averageGrade += float64(grade)
//
//			}
//			averageGrade = averageGrade / float64(len(val.Grade))
//		}
//	}
//	return averageGrade
//}

func (s Student) Calculate() float64 {
	var a float64
	for _, val := range s.Grade {
		a += float64(val)
	}
	a = a / float64(len(s.Grade))
	return a
}

func printStudents(students []Student) {
	for _, val := range students {
		fmt.Printf("Name: %s, Grades: %d\n", val.Name, val.Grade)
	}
}

func main() {

	var students []Student

	students = addStudent(students, Student{"Jane", []int{50, 60, 70}})
	students = addStudent(students, Student{"Kate", []int{55, 63, 75}})
	fmt.Println(students)

	students = updateGrades(students, "Jane", []int{85, 75, 65})
	fmt.Println(students)

	//name := "Kate"
	//grade := calculateAverage(students, name)
	//fmt.Println("Average grade of ", name, grade)

	//var student Student = Student{"Ka", []int{2, 4, 5}}
	//x := student.Calculate()
	//fmt.Println(x)

	students[0].Calculate()

	printStudents(students)

}
