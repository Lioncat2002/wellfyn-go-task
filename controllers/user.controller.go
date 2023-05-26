package controllers

import (
	"main/models"
	"main/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wabarc/go-catbox"
)

// SignUp and Login format
type Data struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Other data format
type UpdateData struct {
	Name   string `json:"name"`
	Imgurl string `json:"imgurl"`
	Phno   string `json:"phno"`
}

// CurrentUser used for getting current user
func CurrentUser(c *gin.Context) {

	email, err := token.ExtractEmail(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

// a two phase process is better
// update the normal data first
// then update the resume
// UpdateResume used for uploading resume
func UpdateResume(c *gin.Context) {

	email, err := token.ExtractEmail(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := c.FormFile("resume") //the file name should be resume
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SaveUploadedFile(file, "tmp/"+file.Filename) //save file to disk since catbox doesn't support uploading file from memory
	//upload to catbox
	url, err := catbox.New(nil).Upload("tmp/" + file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.Resume = url
	_, err = models.UpdateUser(u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": u,
	})
}

// UpdateUser updates the user's data
func UpdateUser(c *gin.Context) {

	email, err := token.ExtractEmail(c) //this will check if the token is valid and return the email

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input UpdateData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.Name = input.Name
	u.Imgurl = input.Imgurl
	u.Phno = input.Phno
	_, err = models.UpdateUser(u) //update all the data expect resume

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": u, //send back the saved data
	})
}
func Login(c *gin.Context) {
	//login the user
	var input Data

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect.", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token}) //send JWT token

}

func SignUp(c *gin.Context) {

	var data Data
	//if the sent data is in valid format
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{}
	user.Email = data.Email
	user.Password = data.Password

	_, err := user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := models.LoginCheck(user.Email, user.Password) //check if user exists
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token, //send JWT token
	})
}
