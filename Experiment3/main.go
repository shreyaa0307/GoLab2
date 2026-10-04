package main

import "fmt"

func main() {

	// Slice operations
	students := []string{"Rahul", "Priya", "Aman"}

	fmt.Println("Initial Slice:", students)

	// Add
	students = append(students, "Neha")
	fmt.Println("After Adding Neha:", students)

	// Remove by index
	index := 1
	students = append(students[:index], students[index+1:]...)
	fmt.Println("After Removing Index 1:", students)

	// Update
	students[1] = "Riya"
	fmt.Println("After Updating Index 1:", students)

	// Map operations
	marks := map[string]int{
		"Math":    85,
		"Science": 90,
		"English": 78,
	}

	fmt.Println("\nInitial Map:", marks)

	// Insert
	marks["Computer"] = 95
	fmt.Println("After Inserting Computer:", marks)

	// Delete
	delete(marks, "English")
	fmt.Println("After Deleting English:", marks)

	// Lookup
	value, exists := marks["Math"]

	if exists {
		fmt.Println("Math Marks:", value)
	} else {
		fmt.Println("Math subject not found")
	}
}
