package main
import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Score float32
	Sex string
}

func (s Monster) Print() {
	fmt.Println("----start----")
	fmt.Println(s)
	fmt.Println("----end ----")
}

func (s Monster) GetSum(a int, b int) int {
	return a + b
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(b interface{}) {
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)
	
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取结构体有几个字段
	num := rVal.NumField()
	fmt.Printf("struct has %d fileds\n", num)

	//遍历结构体字段
	for i := 0; i < num; i++ {
		fmt.Printf("filed %d: 值为%v\n", i, rVal.Field(i))

		//获取标签，使用reflect.Type获取
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("filed %d: tag为%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	rVal.Method(1).Call(nil)   //排序按照函数名字排序比较ASCII码
}

func main() {
	var a Monster = Monster{
		Name: "黄鼠狼",
		Age: 400,
		Score: 30.8,
	}
	TestStruct(a)

}