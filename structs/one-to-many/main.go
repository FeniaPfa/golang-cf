package main

import "fmt"

type Course struct {
	Title  string
	Videos []Video
}

type Video struct {
	Title  string
	Course Course
}

func main() {

	video1 := Video{Title: "Intro"}
	video2 := Video{Title: "Basics"}

	videos := []Video{ video1, video2}


	course := Course{Title: "Go Course", Videos: videos}

	video1.Course = course
	video2.Course = course

	fmt.Println(video1)
	fmt.Println(video2)
	fmt.Println(course)

	fmt.Println(video1.Course.Title)


	for _, video := range course.Videos{
		fmt.Println(video.Title)
	}

}