package main

import "testing"

type MockCollegeStudentTeacher struct{}

func (m *MockCollegeStudentTeacher) GetStudent(badgeNumber BadgeNumber) (*Student, error) {
	return &Student{
		StudentBadgeNumber:       "BSCS1035",
		StudentName:              "Alex",
		Specialization:           "ComputerScience",
		AssociatedClassTeacherID: 101,
	}, nil
}

func (m *MockCollegeStudentTeacher) GetTeacher(classID int) (*Teacher, error) {
	return &Teacher{
		TeacherName:    "Langu",
		Specialization: "ComputerScience",
	}, nil
}

func TestHandler(t *testing.T) {

	agent := &MockCollegeStudentTeacher{}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "Get the User", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := SomeHandler(BadgeNumber("BSCS1035"), agent)
			if (err != nil) != tt.wantErr {
				t.Errorf("SomeHandler() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
