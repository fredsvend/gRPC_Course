package main;

import (
	"fmt"
	pb "grpccourse/DISYS_2"
)

func main() {
	course := pb.Course{CourseName: "Testing", CourseId: 2, CourseWorkload: 2, CourseStudentsatisfaction: 0}
	
	fmt.Println("CourseID: ", course.GetCourseId())
}


