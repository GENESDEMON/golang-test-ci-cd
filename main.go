package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func main() {
	//setting up routes
	router := gin.Default()
	router.GET("/", WelcomeMessage)
	router.POST("/createStudent", CreateStudent())
	router.GET("/students", GetStudents())
	//router.PUT("/updatestudent/:id", UpdateAStudent())
	//router.DELETE("/deletestudent/:id", DeleteAStudent())
	router.Run()
}

//Setting up student struct
type Student struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Level      string `json:"level"`
}

//Students database
var students = []Student{
	{
		ID:         "10000xbcd3",
		Name:       "Alicia Winds",
		Department: "Political Science",
		Level:      "Year 3",
	},
}

//Welcome Message
func WelcomeMessage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hey boss!"})
}

//Create a new student account
func CreateStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newStudent Student
		if err := c.BindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		//Generte student ID
		newStudent.ID = xid.New().String()
		students = append(students, newStudent)
		c.JSON(http.StatusCreated, newStudent)
	}

}

//Get a list of Students
func GetStudents() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Fetch all students in the DB
		c.JSON(http.StatusOK, students)
	}
}

//Update a student Details
// func UpdateAStudent() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		Update a particular student
// 		id := c.Param("id")
// 		var newStudent Student
// 		if err := c.BindJSON(&newStudent); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"Status":  http.StatusBadRequest,
// 				"Message": "error",
// 				"Data":    map[string]interface{}{"data": err.Error()}})
// 			return
// 		}
// 		index := -1
// 		for i := 0; i < len(students); i++ {
// 			if students[i].ID == id {
// 				index = 1
// 			}
// 		}
// 		if index == -1 {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "Student profile not found",
// 			})
// 			return
// 		}
// 		students[index] = newStudent
// 		c.JSON(http.StatusOK, newStudent)
// 	}
// }

//Delete a student profile
// func DeleteAStudent() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		id := c.Param("id")
// 		index := -1
// 		for i := 0; i < len(students); i++ {
// 			if students[i].ID == id {
// 				index = 1
// 			}
// 		}
// 		if index == -1 {
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error": "Student profile not found",
// 			})
// 			return
// 		}
// 		students = append(students[:index], students[index+1:]...)
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Student profile deleted",
// 		})
// 	}
// }
