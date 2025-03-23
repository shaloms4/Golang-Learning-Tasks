package main

import "fmt"

func gradeCalculator(name string, subjectGrades map[string]float64) (map[string]string, float64) {
	nameToGrade := make(map[string]string)
	var slice []float64

	for subjectName, grade := range subjectGrades {
		if grade <= 0 || grade > 100 {
			continue
		}

		slice = append(slice, grade)

		if grade >= 90 {
			nameToGrade[subjectName] = "A+"
		} else if grade >= 85 {
			nameToGrade[subjectName] = "A"
		} else if grade >= 75 {
			nameToGrade[subjectName] = "B+"
		} else if grade >= 70 {
			nameToGrade[subjectName] = "B"
		} else if grade >= 65 {
			nameToGrade[subjectName] = "B-"
		} else if grade >= 60 {
			nameToGrade[subjectName] = "C+"
		} else if grade >= 55 {
			nameToGrade[subjectName] = "C"
		} else if grade >= 50 {
			nameToGrade[subjectName] = "C-"
		} else if grade >= 40 {
			nameToGrade[subjectName] = "D"
		} else {
			nameToGrade[subjectName] = "F"
		}
	}

	var total float64
	for _, grade := range slice {
		total += grade
	}

	var avgGrade float64
	if len(slice) > 0 {
		avgGrade = total / float64(len(slice))
	} else {
		avgGrade = 0
	}

	return nameToGrade, avgGrade
}

type student struct {
	name              string
	subject_and_grade map[string]float64
}

func (s student) averageGrade() float64 {
	var total float64
	for _, grade := range s.subject_and_grade {
		total += grade
	}
	return total / float64(len(s.subject_and_grade))
}

func main() {
	var name string
	var numberOfSubjects int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter number of subjects: ")
	fmt.Scanln(&numberOfSubjects)

	// Collect subject grades
	subjectGrades := make(map[string]float64)
	for i := 0; i < numberOfSubjects; i++ {
		var subjectName string
		var grade float64

		fmt.Printf("Enter subject name %d: ", i+1)
		fmt.Scanln(&subjectName)

		fmt.Printf("Enter grade for %s: ", subjectName)
		fmt.Scanln(&grade)

		subjectGrades[subjectName] = grade
	}

	gradeToLetter, avgGrade := gradeCalculator(name, subjectGrades)

	fmt.Printf("Name: %s\n", name)
	for subject, grade := range gradeToLetter {
		fmt.Printf("%s : %s\n", subject, grade)
	}
	fmt.Printf("Average grade: %.2f\n", avgGrade)
}
