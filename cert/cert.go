package cert

import (
	"fmt"
	"strings"
	"time"
)

type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

var MaxLenCourse = 20
var MaxLenName = 30

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {
	c, cErr := validateCourse(course)
	if cErr != nil {
		return nil, cErr
	}

	n, nErr := validateName(name)
	if nErr != nil {
		return nil, nErr
	}

	d, dErr := parseDate(date)
	if dErr != nil {
		return nil, dErr
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Customer Support Operations",
		LabelPresented:     "This Certificate is Presented To",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}

	return cert, nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, "course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateName(name string) (string, error) {
	c, err := validateStr(name, MaxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 || len(c) >= maxLen {
		return c, fmt.Errorf("Invalid string. got=%s, len=%d", c, len(c))
	}
	return c, nil
}
