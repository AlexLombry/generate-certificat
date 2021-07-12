package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	courseName := "GOLANG COURSE"
	c, err := New("Golang", "Bob", "2021-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != courseName {
		t.Errorf("Course courseName is not valid. expected=%v, get=%v", courseName, c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2021-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseNameTooLong(t *testing.T) {
	courseName :=  "MY SUPER COURSE NAME IS FREAKING TOO LONG FOR US GODDAMMIT"
	_, err := New(courseName, "Bob", "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long courseName name (courseName=%s)", courseName)
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2021-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestNameNameTooLong(t *testing.T) {
	name :=  "MY SUPER NAME IS FREAKING TOO LONG FOR US GODDAMMIT"
	_, err := New("Golang", name, "2021-05-31")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (Name=%s)", name)
	}
}