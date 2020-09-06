package grade

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

type SchoolI interface {
	AddStudent(grade int, name string)error
	ListStudent(grade int)([]string, error)
	ListAllGradeStudents()(map[int][]string, error)
	SetStudent()
}

type SchoolA struct {
	students map[int][]string
}

func NewSchool(schoolFactory string)SchoolI{
	switch schoolFactory {
	case "SchoolA":
		return &SchoolA{}
	default:
		return &SchoolA{}
	}
}

func (school *SchoolA) SetStudent(){
	school.students = make(map[int][]string, 0)
}

func (school *SchoolA) AddStudent(grade int, name string)error{
	if grade < 1{
		return errors.New("Invalid request")
	}
	if name == ""{
		return errors.New("Invalid request")
	}
	if school.students == nil{
		return errors.New("Invalid object")
	}
	students := school.students[grade]

	isEmpty := checkIsStudentEmpty(students)

	if isEmpty {
		students = CreateList()
	}
	students = append(students, name)
	school.students[grade] = students
	return nil
}

func (school *SchoolA) ListStudent(grade int)([]string, error){
	err := validateListStudent(grade, school)
	if err != nil{
		return nil, err
	}
	students := school.students[grade]

	if students == nil{
		return []string{}, nil
	}
	return students, nil
}

func (school *SchoolA) ListAllGradeStudents()(map[int][]string, error){
	temp := school.students
	if temp == nil{
		return nil, errors.New("School student can't be nil")
	}
	grades := make([]int, 0)
	for k, _ := range school.students{
		grades = append(grades, k)
	}
	sort.Ints(grades)
	for _, v := range grades{
         students := temp[v]
         sort.Strings(students)
         temp[v] = students
	}
	return temp, nil
}

func checkIsStudentEmpty(students []string)bool{
	if students == nil{
		return true
	}
	return false
}

func CreateList()[]string{
	students := make([]string, 0)
	return students
}

func validateListStudent(grade int, school *SchoolA)error{
	if grade < 1{
		return errors.New("Invalid request")
	}
	if school.students == nil{
		return errors.New("Invalid object")
	}
	return nil
}

func main() {
	school := NewSchool("SchoolA")
	school.SetStudent()
	student := "jim"
	student1 := "velan"
	student2 := "Apple"
	student3 := "Banana"
	student4 := "Abac"
	err := school.AddStudent(2, student)
	if err != nil {
		log.Fatal("Error should be empty")
	}
	err = school.AddStudent(2, student1)
	if err != nil {
		log.Fatal("Error should be empty")
	}
	err = school.AddStudent(1, student2)
	if err != nil {
		log.Fatal("Error should be empty")
	}
	err = school.AddStudent(1, student3)
	if err != nil {
		log.Fatal("Error should be empty")
	}
	err = school.AddStudent(1, student4)
	if err != nil {
		log.Fatal("Error should be empty")
	}
	studs, err := school.ListAllGradeStudents()
	if err != nil{

	}
	fmt.Println(studs)

}
