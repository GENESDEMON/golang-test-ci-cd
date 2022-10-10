package main
import "github.com/gin-gonic/gin"

func SimpleTestRouter() *gin.Engine {
	router := gin.Default()
	return router
}

//Welcome message test
func TestWelcomeMessage((t *testing.T) {
    mockResponse := `{"message": "This a greeting from the test world!"}`
    r := SetUpRouter()
    r.GET("/", HomepageHandler)
    req, _ := http.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
}

//Create student test
func TestCreateStudent((t *testing.T) {
    r := SetUpRouter()
    r.POST("/createStudent", CreateStudent()
    studentId := xid.New().String()
    student := Student{
        ID: studentID,
        Name:       "Marilyn Monroe",
		Department: "Political Science",
		Level:      "Year 5",
    }
    jsonValue, _ := json.Marshal(student)
    req, _ := http.NewRequest("POST", "/createStudent", bytes.NewBuffer(jsonValue))

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)
}



func TestGetStudents(t *testing.T) {
    r := SetUpRouter()
    r.GET("/students", GetStudents)
    req, _ := http.NewRequest("GET", "/students", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    var students []Student
    json.Unmarshal(w.Body.Bytes(), &students)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.NotEmpty(t, students)
}



