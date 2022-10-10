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
	router.GET("/students/:id", GetStudentById())
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

//Get one student from the list
func GetStudentById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		for _, a := range students {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}
