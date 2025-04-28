package main
import (
	"fmt"
	"reflect"
)
type Stu struct {
	Name string
	Age int
}
func reflectTest1(b interface{}) {
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	//2.获取到reflect.value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	
	n1 := 10 + rVal.Int()
	fmt.Println("n2 = ", n1)

	// 把rval转成interface{}
	iv := rVal.Interface()

	//将interface{}通过断言转成需要的类型
	n2 := iv.(int)
	fmt.Println(n2)
}

func reflectTest2(b interface{}) {
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println(rTyp)

	//2.获取到reflect.value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	
	//3.获取kind
	fmt.Println(rVal.Kind())
	fmt.Println(rTyp.Kind())
	
	// 把rval转成interface{}
	iv := rVal.Interface()

	//将interface{}通过断言转成需要的类型
	stu, ok := iv.(Stu)
	if ok {
		fmt.Printf("stu.name=%v, stu.Age=%v\n", stu.Name, stu.Age)
	}
	
}

//通过反射修改变量的值
func reflectTest3(b interface{}) {
	//2.获取到reflect.value
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
	rVal.Elem().SetInt(20)  //rval.Elem相当于c++解引用
	
	
}

func reflectTest4(b interface{}) {
	
	rVal := reflect.ValueOf(b)
	rTyp := reflect.TypeOf(b)


	fmt.Println(rVal.Kind())
	fmt.Println(rTyp.Kind())

	i := rVal.Interface()

	n := i.(float64)
	fmt.Printf("n type is %T, n is %v\n", n, n)
	
}

func main() {

	var num int = 100
	reflectTest1(num)


	stu1 := Stu {
		Name : "Tom",
		Age : 18,
	}
	reflectTest2(stu1)

	reflectTest3(&num)
	fmt.Println("num=", num)

	var v float64 = 1.2
	reflectTest4(v)
}