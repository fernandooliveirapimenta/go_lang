package main

import (
	"errors"
	"fmt"
	"math"
)

var users = map[string]string{
	"305":"Sue",
	"204":"Bob",
	"631":"Jake",
	"073":"Tracy",
}

//custo type
type id string

// struct
type user struct {
	name string
	age int
	balance float64
	member bool
}

type point struct {
	x int
	y int
}

func compare()(bool, bool){
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	point2 := struct {
		x int
		y int
	}{}

	point2.x = 10
	point2.y = 5

	point3 := point{10,10}

	return point1 == point2, point1 == point3
}

type name string
type location struct {
	x int
	y int
}
type size struct {
	width int
	height int
}
type dot struct {
	name
	location
	size
}

func getDots() []dot{
	var dot1 dot
	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width: 5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 101
	dot4.size.width = 87
	dot4.size.height = 43

	return []dot{dot1,dot2,dot3,dot4}
}

type fernando struct {
	nome string
}
func main() {

	// Deleting elements from a map
	fmt.Println(users)
	deleteUser("204")
	fmt.Println(users)

   // Deleting elements from slice arrays
   slicee := []string{"Sue","Bob","Jake", "Tracy","Fernando"}
   fmt.Println(slicee)
   slicee = append(slicee[:3], slicee[3+1:]...)
   fmt.Println(slicee)

   // Creating a simple custom type
   id1,id2,id3 := getIds()
   fmt.Println("id1==id2", id1==id3)
   fmt.Println("id2==id3", id2==id3)

   //Creating struct types and values
   ustruct := getUsers2()
   for i:=0;i<len(ustruct); i++ {
   	fmt.Printf("%v: %#v\n",i,ustruct[i])
   }

   // Comparing structs each other
   a,b:=compare()
   fmt.Println("point1==point2",a)
   fmt.Println("point1==point3",b)

   // Struct embedding and initialization
   dots := getDots()
   for i:=0;i<len(dots);i++{
   	fmt.Printf("dot%v:%#v\n",i+1,dots[i])
   }

   // Numeric type conversion

	fmt.Println()
	fmt.Println()
   fmt.Println(convert())

   //Type assertion
	fmt.Println()
	fmt.Println()
   res, _ := doubler(5)
   fmt.Println("5:",res)
   res, _ = doubler("yum")
	fmt.Println("yum:",res)
	_, err := doubler(true)
	fmt.Println("true:",err)

	// Type Switch
	res2, _ := doublerSwitch(-5)
	fmt.Println("-5 :", res2)
	res2, _ = doublerSwitch(5)
	fmt.Println("5 :", res2)
	res2, _ = doublerSwitch("yum")
	fmt.Println("yum :", res2)
	res2, _ = doublerSwitch(true)
	fmt.Println("true:", res)
	res2, _ = doublerSwitch(float32(3.14))
	fmt.Println("3.14:", res2)

	// Type checker
	fmt.Println()
	fmt.Println()
	types := []interface{}{
		3,
		3.64,
		"oi",
		true,
		fernando{"fefe"},
	}

	for _, v := range types {
		fmt.Println(typeChecker(v))
	}

}

func typeChecker(v interface{})string{

	switch v.(type) {

	case int,int8,int64,uint:
		return "int"
	case float32,float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case fernando:
		return "struct"
	default:
		return "nenhum"

	}
}

func doublerSwitch(v interface{})(string, error){

	switch t:=v.(type) {

	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue",nil
		}
		return "falsefalse",nil

	case float32,float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f*2),nil
		}
		return fmt.Sprint(t.(float32) * 2), nil

	case int:
		return fmt.Sprint(t*2),nil
	case int8:
		return fmt.Sprint(t*2),nil
	case int16:
		return fmt.Sprint(t*2),nil
	case int32:
		return fmt.Sprint(t*2),nil
	case int64:
		return fmt.Sprint(t*2),nil
	case uint:
		return fmt.Sprint(t*2),nil
	case uint8:
		return fmt.Sprint(t*2),nil
	case uint16:
		return fmt.Sprint(t*2),nil
	case uint32:
		return fmt.Sprint(t*2),nil
	case uint64:
		return fmt.Sprint(t*2),nil
	default:
		return "", errors.New("unsupported")

	}
}
func doubler(v interface{})(string, error){
	if i, ok := v.(int); ok {
		return fmt.Sprint(i * 2), nil
	}

	if s, ok := v.(string); ok {
		return s + s, nil
	}

	return "", errors.New("unsuported type passed")
}
func convert() string  {

	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	m := fmt.Sprintf("int8 = %v > in64 = %v\n", i8, int64(i8))
	m += fmt.Sprintf("int = %v > in8 = %v\n",i, int8(i))
	m += fmt.Sprintf("int8 = %v > float32 = %v\n", i8, float32(i8))
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))

	return m
}

func getUsers2() []user{
	u1 := user{
		name: "Tracy",
		age: 51,
		balance: 98.43,
		member: true,
	}

	u2 := user{
		age:19,
		name: "Nick",
	}

	u3 := user{
		"Bob",
		25,
		0,
		false,
	}

	var u4 user
	u4.name = "Sue"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09

	return []user{u1,u2,u3,u4}
}

func getIds()(id, id, id){
	var id1 id
	var id2 id = "123"
	var id3 id
	id3 = "123"
	return id1,id2,id3
}

func deleteUser(id string) {
	delete(users, id)
}
