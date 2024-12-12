package myfunc

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string
	Age  int
}

type Teacher struct {
	Name  string
	Level string
}

func Test1(c *gin.Context) {
	name := "这是传递字符串数据"
	c.HTML(200, "demo01/h1.html", name)
}

func Test2(c *gin.Context) {
	stu := Student{
		Name: "这是传递结构体的数据",
		Age:  0,
	}
	c.HTML(200, "demo02/h2.html", stu)
}
func Test3(c *gin.Context) {
	var arr [2]string
	arr[0] = "这是传递"
	arr[1] = "数组的数据"
	c.HTML(200, "demo03/h3.html", arr)
}

func Test4(c *gin.Context) {
	var arr [2]Student
	arr[0] = Student{
		Name: "鲜于昊辰",
		Age:  19,
	}
	arr[1] = Student{
		Name: "这是结构体数组的传递",
		Age:  1141514,
	}
	c.HTML(200, "demo04/h4.html", arr)
}
func Test5(c *gin.Context) {
	var a map[string]string = make(map[string]string, 2)
	a["前一段"] = "这是map"
	a["后一段"] = "的渲染"
	c.HTML(200, "demo05/h5.html", a)
}
func Test6(c *gin.Context) {
	var a map[string]Teacher = make(map[string]Teacher, 2)
	a["No1"] = Teacher{
		Name:  "张耀文",
		Level: "副教授",
	}
	a["No2"] = Teacher{
		Name:  "廖浩徳",
		Level: "教授",
	}
	c.HTML(200, "demo06/h6.html", a)
}
func Test7(c *gin.Context) {
	s := []string{"这是", "切片的", "渲染啊"}
	c.HTML(200, "demo07/h7.html", s)
}
