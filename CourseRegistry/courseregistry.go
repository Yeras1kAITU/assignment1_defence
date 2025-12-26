package CourseRegistry

import (
	"errors"
	"fmt"
)

// Student struct
type Student struct {
	ID      uint64
	Name    string
	Courses []string
}

// Registry struct
type Registry struct {
	Students map[uint64]Student
}

// NewRegistry creates a new Registry
func NewRegistry() *Registry {
	return &Registry{
		Students: make(map[uint64]Student),
	}
}

// AddStudent adds a new student to the registry
func (r *Registry) AddStudent(student Student) error {
	// Check if student ID already exists
	if _, exists := r.Students[student.ID]; exists {
		return errors.New("student ID already exists")
	}

	// Check if name is empty
	if student.Name == "" {
		return errors.New("student name cannot be empty")
	}

	// Initialize empty courses slice if nil
	if student.Courses == nil {
		student.Courses = []string{}
	}

	// Add student to registry
	r.Students[student.ID] = student
	return nil
}

// EnrollCourse enrolls a student in a course
func (r *Registry) EnrollCourse(studentID uint64, course string) error {
	// Check if student exists
	student, exists := r.Students[studentID]
	if !exists {
		return errors.New("student does not exist")
	}

	// Check if course name is empty
	if course == "" {
		return errors.New("course name cannot be empty")
	}

	// Check if student is already enrolled in the course
	for _, c := range student.Courses {
		if c == course {
			return errors.New("student is already enrolled in this course")
		}
	}

	// Add course to student's course list
	student.Courses = append(student.Courses, course)
	r.Students[studentID] = student
	return nil
}

// RemoveCourse removes a course from a student's enrollment
func (r *Registry) RemoveCourse(studentID uint64, course string) error {
	// Check if student exists
	student, exists := r.Students[studentID]
	if !exists {
		return errors.New("student does not exist")
	}

	// Find and remove the course
	found := false
	newCourses := []string{}
	for _, c := range student.Courses {
		if c == course {
			found = true
			continue
		}
		newCourses = append(newCourses, c)
	}

	if !found {
		return errors.New("course not found for this student")
	}

	// Update student's courses
	student.Courses = newCourses
	r.Students[studentID] = student
	return nil
}

// ListStudents returns all students as a slice
func (r *Registry) ListStudents() []Student {
	students := make([]Student, 0, len(r.Students))

	for _, student := range r.Students {
		students = append(students, student)
	}

	return students
}

// CoursesCount returns a map with course enrollment statistics
func (r *Registry) CoursesCount() map[string]int {
	courseCount := make(map[string]int)

	// Count enrollments for each course
	for _, student := range r.Students {
		for _, course := range student.Courses {
			courseCount[course]++
		}
	}

	return courseCount
}

// PrintStudents prints all students in the required format
func (r *Registry) PrintStudents() {
	students := r.ListStudents()

	if len(students) == 0 {
		fmt.Println("No students in registry")
		return
	}

	fmt.Println("\n=== Students in Registry ===")
	for _, student := range students {
		fmt.Printf("ID: %d | Name: %s | Courses: %v\n",
			student.ID, student.Name, student.Courses)
	}
}

// PrintCourseStatistics prints course enrollment statistics
func (r *Registry) PrintCourseStatistics() {
	courseStats := r.CoursesCount()

	if len(courseStats) == 0 {
		fmt.Println("No course enrollments")
		return
	}

	fmt.Println("\n=== Course Enrollment Statistics ===")
	for course, count := range courseStats {
		fmt.Printf("%s → %d\n", course, count)
	}
}

// RunCourseRegistry provides console interface
func RunCourseRegistry() {
	registry := NewRegistry()

	// Add initial test data
	registry.AddStudent(Student{ID: 1, Name: "Alice", Courses: []string{"Go", "Databases"}})
	registry.AddStudent(Student{ID: 2, Name: "Bob", Courses: []string{"Go"}})
	registry.AddStudent(Student{ID: 3, Name: "Charlie", Courses: []string{}})

	fmt.Println("Initial test data loaded:")
	registry.PrintStudents()
	registry.PrintCourseStatistics()

	for {
		fmt.Println("\n=== Course Registry System ===")
		fmt.Println("1. Add Student")
		fmt.Println("2. Enroll Course")
		fmt.Println("3. Remove Course")
		fmt.Println("4. List Students")
		fmt.Println("5. Course Statistics")
		fmt.Println("6. Exit")
		fmt.Print("Select option (1-6): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id uint64
			var name string

			fmt.Print("Enter student ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter student name: ")
			fmt.Scan(&name)

			student := Student{
				ID:      id,
				Name:    name,
				Courses: []string{},
			}

			if err := registry.AddStudent(student); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Student added successfully!")
			}

		case 2:
			var studentID uint64
			var course string

			fmt.Print("Enter student ID: ")
			fmt.Scan(&studentID)
			fmt.Print("Enter course name: ")
			fmt.Scan(&course)

			if err := registry.EnrollCourse(studentID, course); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Course enrolled successfully!")
			}

		case 3:
			var studentID uint64
			var course string

			fmt.Print("Enter student ID: ")
			fmt.Scan(&studentID)
			fmt.Print("Enter course name: ")
			fmt.Scan(&course)

			if err := registry.RemoveCourse(studentID, course); err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				fmt.Println("Course removed successfully!")
			}

		case 4:
			registry.PrintStudents()

		case 5:
			registry.PrintCourseStatistics()

		case 6:
			fmt.Println("Exiting Course Registry System...")
			return

		default:
			fmt.Println("Invalid choice! Please select 1-6.")
		}
	}
}
