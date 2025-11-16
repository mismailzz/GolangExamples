// BACKEND APP

package main 

import (
	"errors"
	"fmt"
)

var (
	ErrStudentNotFound = errors.New("student not found")
	ErrTeacherNotFound = errors.New("teacher not found")
) 


type BadgeNumber string

// Data structs hold information
type Student struct {
	StudentBadgeNumber BadgeNumber
	StudentName	string
	Specialization string 
	AssociatedClassTeacherID int 
}

// Data structs hold information
type Teacher struct {
	TeacherName string
	Specialization string 
}

// Service structs hold behavior/logic
// A CollegeDB is the administrative office with all student records
// This hold the DB and it deliver the data from DB
type CollegeDB struct {
	StudentsDB map[BadgeNumber]Student
	TeacherDB map[int]Teacher
}


// act as a initialization or constructor for DB 
func New() *CollegeDB {
	return &CollegeDB {
		StudentsDB: map[BadgeNumber]Student{
			"BSCS1035": {"BSCS1035", "Alex", "ComputerScience", 101},
			"BSCS1036": {"BSCS1036", "Akram", "Math", 201},
		},
		TeacherDB: map[int]Teacher{
			101: {"Langu", "ComputerScience"},
			201: {"Stuni", "Math"},
		},
	}
}



// Think about this func. get the data from some backendDB
func (s *CollegeDB) GetStudent (badgeNumber BadgeNumber) (*Student, error) {

	resp, found := s.StudentsDB[badgeNumber]
	if !found {
		return nil, ErrStudentNotFound
	}

	return &resp, nil
}

func (t *CollegeDB) GetTeacher (classID int) (*Teacher, error) {

	resp, found := t.TeacherDB[classID]
	if !found {
		return nil, ErrTeacherNotFound
	}

	return &resp, nil
}



// Just Check - if everything works fine by creating main
func main(){

	fmt.Println("Main Start")
	college := New() //initialize the DB or like connecting to DB (in realWorld)

	badgeNumber := BadgeNumber("BSCS1035")
	response, err := college.GetStudent(badgeNumber)
	if err != nil {
		fmt.Errorf("Error:%v\n", err)
	}
	fmt.Printf("UserFound: %v\n", response.StudentName)
	fmt.Println("Main End")

}



