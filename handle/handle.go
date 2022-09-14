package handle

import (
	"fmt"
	"gogin/datareq"
	"gogin/signup"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type signups struct {
	sign signup.Service
}

type handle struct {
	handle datareq.Service
}

func Newhandle(service datareq.Service) *handle {
	return &handle{service}
}
func Newsignup(test signup.Service) *signups {
	return &signups{test}
}

//	func (h *handle) Getall(c *gin.Context) {
//		test, err := h.handle.Getall()
//		test1 := []string{}
//		if err != nil {
//			for _, e := range err.(validator.ValidationErrors) {
//				test := fmt.Sprintf("error happen in field %s, due %s", e.Field(), e.ActualTag())
//				test1 = append(test1, test)
//			}
//			c.JSON(http.StatusBadRequest, test1)
//			return
//		}
//		var data []datareq.Dataresp
//
//		for _, r := range test {
//			resp := serving(r)
//			data = append(data, resp)
//		}
//		c.JSON(http.StatusOK, data)
//	}
func (h *handle) Getbyid(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.Atoi(id)
	getid, err := h.handle.Getbyid(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": serving(getid),
	})
}
func (h *handle) Create(c *gin.Context) {
	var data datareq.Userreq
	err := c.ShouldBindJSON(&data)
	// test1 is for catch the error
	test1 := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			test := fmt.Sprintf("error happen in field %s, due %s", e.Field(), e.ActualTag())
			test1 = append(test1, test)
		}
		c.JSON(http.StatusBadRequest, test1)
		return
	}
	create, err := h.handle.Create(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": serving(create),
	})
}
func (h *handle) Update(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.Atoi(id)
	var data datareq.Userreq
	err := c.ShouldBindJSON(&data)
	// test1 is for catch the error
	test1 := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			test := fmt.Sprintf("error happen in field %s, due %s", e.Field(), e.ActualTag())
			test1 = append(test1, test)
		}
		c.JSON(http.StatusBadRequest, test1)
		return
	}
	update, err := h.handle.Update(data, num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Data": serving(update),
	})
}
func (h *handle) Delete(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.Atoi(id)

	delete, err := h.handle.Delete(num)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Deleting Data": serving(delete),
	})
}

func (s *signups) Signup(c *gin.Context) {
	var data signup.Signup_Req
	err := c.ShouldBindJSON(&data)
	test1 := []string{}
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			test := fmt.Sprintf("error happen in field %s, due %s", e.Field(), e.ActualTag())
			test1 = append(test1, test)
		}
		c.JSON(http.StatusBadRequest, test1)
		return
	}
	
	test,err := s.sign.Signup(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"data": signupserve(test),
	})
}

func (s *signups) Login(c *gin.Context){
	var data signup.Loginreq
	err := c.ShouldBindJSON(&data)
		if err != nil {
			panic(err)
		}
	err = s.sign.Login(data)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}
	c.JSON(http.StatusOK, gin.H{
		"Status": true, 
	})
}

func serving(t datareq.DBdata) datareq.Dataresp {
	return datareq.Dataresp{
		ID:      t.ID,
		Name:    t.Name,
		Deposit: t.Deposit,
	}
}
func signupserve (t signup.Signup) signup.Response{
	return signup.Response{
		ID:       t.ID,
		Email:    t.Email,
		Username: t.Username,
	}
}