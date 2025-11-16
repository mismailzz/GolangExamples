package main


// To get functionality here - we just create the interface here because it doesnt exists earlier
// Interface will allow in a snicky way like if we made same signature of func in implementation 
// Then we can access here too - because then it will act like implemented struct on backend


// Defining Separate Interface, can be used later as one type
type StudentGetter interface {
	GetStudent (badgeNumber BadgeNumber) (*Student, error)
}

type TeacherGetter interface {
	GetTeacher (classID int) (*Teacher, error) 
}


// Can be combine into one
type CollegeStudentTeacher interface {
	StudentGetter
	TeacherGetter
} 


func SomeHandler(badgeNumber BadgeNumber, agent CollegeStudentTeacher) error {

	resp, err := agent.GetStudent(badgeNumber)
	if err != nil {
		return err 
	}

	_, err = agent.GetTeacher(resp.AssociatedClassTeacherID)
	if err != nil {
		return err 
	}

	return nil 
}

