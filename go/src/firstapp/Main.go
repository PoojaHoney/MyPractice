package main

import (
	// "fmt"
	// "io/ioutil"
	// "log"
	// "bytes"
	"fmt"
	// "strconv"
	// "math"
	// "net/http"
	// "strconv"
	// "reflect"
)

// var o int = 66

// var (
// 	std_id int = 786    //declare multiple variables at package level
// )

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// type Student struct {
// 	name        string `required max:"10"`
// 	age         int
// 	phoneNumber int
// 	parents     []string
// }

// type StudentClgDetails struct {
// 	Student
// 	id     int
// 	barnch string
// }

// func returnTrue() bool {
// 	fmt.Println("returning true")
// 	return true
// }

// func myPanic() {
// 	fmt.Println("In Panic Function")
// 	panic("Problem Occured")
// }

// type pointerStruct struct {
// 	name string
// 	age int
// 	dob string
// }

// func my_fun(param1, param2 string, param3 int) {
// 	fmt.Println(param1 + param2 + strconv.Itoa(param3))
// }

// func my_func_pointers(param1, param2 *string) {
// 	fmt.Println(param1, param2)
// 	fmt.Println(*param1, *param2)
// 	fmt.Println(&param1, &param2)
// 	*param1 = "shinchan"
// }

// func sum(msg string, values ...int) *int{ //slice the n number of parameters to the function
// 	result := 0
// 	for _, v :=range values {
// 		result += v
// 		// fmt.Println(inx)
// 	}
// 	fmt.Println(msg, result)
// 	result += 10
// 	return &result //golang holds the returning pointer variable in heap memory(shared memory) instead of clearing it
// }

// func sum(msg string, values ...int) (result int){ //syntactic sugar for reclaring the return variable will be in this func scope
// 	for _, v :=range values {
// 		result += v
// 		// fmt.Println(inx)
// 	}
// 	fmt.Println(msg, result)
// 	return  //golang holds the returning pointer variable in heap memory(shared memory) instead of clearing it
// }

// func divide(a, b float32) (float32, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// type greet struct {
// 	msg string
// }

// func greet1(msg string) greet {
// 	return greet{
// 		msg: msg,
// 	}
// }

// type greeter struct{
// 	greeting string
// 	msg string
// }

// func (r greeter) greet() { //(r greeter) this makes the function as method in a known context
// 	fmt.Println(r.greeting, r.msg)
// }

// type writer interface{ //defines behaviour not the data as in struct, we actualy store the method definitions
// 	write([]byte) (int, error)
// }

// type consoleWriter struct {}

// func (cw consoleWriter) write(data []byte)(int, error) {
// 	n, err := fmt.Println(string(data))
// 	return n, err
// }

// type incrementer interface {
// 	Increment() int
// }

// type increments int

// func (ic *increments) Increment() int{
// 	*ic ++
// 	return int(*ic)
// }

// type foobar interface {
// 	foo()
// 	bar()
// }

// type itemA struct{}

// func (a *itemA) foo() {
// 	fmt.Println("foo on A")
// }

// func (a *itemA) bar() {
// 	fmt.Println("bar on A")
// }

// type itemB struct{}

// func (a *itemB) foo() {
// 	fmt.Println("foo on B")
// }

// func (a *itemB) bar() {
// 	fmt.Println("bar on B")
// }

// func doFoo(item foobar) {
// 	item.foo()
// }

// func doBar(item foobar) {
// 	item.bar()
// }

// type toaster interface {
// 	toast()
// }

// type toasterstruct struct{}

// func (a *toasterstruct) toast() {
// 	fmt.Println("In Toast Method")
// }

// func doToast(t toaster) {
// 	fmt.Println("In Do Toast Method")
// 	t.toast()
// }
// func mayBeDoToast(i interface{}) {
// 	fmt.Println("In May Be Do Toast Method")
// 	if t, ok := i.(toaster); ok {
// 		t.toast()
// 	}
// }

func main() {

	//--------Interfaces--------mainly used for scalable
	//methods name = +er -> interface - writer -> method write

	var i interface{} = "0"

	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	}

	// t := new(toasterstruct)

	// t.toast()

	// doToast(t)
	// mayBeDoToast(t)

	//compose interface
	// doBar(&itemA{})
	// doBar(&itemB{})
	// doFoo(&itemA{})
	// doFoo(&itemB{})

	// ic1 := increments(9)
	// // var inc incrementer = &ic1
	// for i:=0; i<10; i++ {
	// 	fmt.Println(ic1.Increment())
	// }

	// ic1.Increment()
	// fmt.Println(ic1)

	// consolewrite := consoleWriter{}
	// consolewrite.write([]byte("Hello World!!"))

	//FUNCTIONS

	// g := greeter{
	// 	greeting: "Good Morning",
	// 	msg: "Hey babe!!",
	// }
	// g.greet()

	// a := greet1("Print me know")
	// fmt.Println(a)

	//anomus function
	// var divide func(float32, float32)(float32, error)
	// divide = func(a, b float32) (float32, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("This is an error")
	// 	} else {
	// 		return a / b, nil
	// 	}
	// }
	// a, err := divide(4.0, 2.0)
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(a)

	// a := func() {
	// 	fmt.Println("In Anonymus Function")
	// }
	// a()
	// a:= func () int{
	// 	fmt.Println("In Ananomus Function")
	// 	return 6
	// }()
	// fmt.Println(a)
	// for i := 0; i< 5; i++ {
	// 	func () {
	// 		fmt.Println(i)
	// 	}()
	// 		fmt.Println("Another Loop")
	// 	func (i int) { //this would be the correct where the compiler could not confuse
	// 		fmt.Println(i)
	// 	}(i)
	// }

	//multiple return values
	// d, err := divide(5.0, 9.0)
	// fmt.Println(d, err)

	//named returning parameters
	// sum := sum("Print The Sum:", 1,2,3,4,5)
	// fmt.Println(sum)

	//variadic parameters
	// sum := sum("Print The Sum:", 1,2,3,4,5)
	// fmt.Println(*sum)

	// a := "Print Me"
	// b := "Please"
	// my_func_pointers(&a, &b)

	// fmt.Println(a)
	// fmt.Println(b)

	// my_fun("Print me", " please", 7)

	//Pointers
	// *int -> pointer to the variable of int type
	// & -> addressing operator
	// * -> dereferencing operator
	// unsafe package - used for handling the addressing arthemtic operation - it ingnores the unexpected error

	//maps
	// a := map[string]int{"one": 1, "two": 2}
	// fmt.Printf("map a: %v\n", a)
	// b := a
	// fmt.Printf("map b: %v\n", b)

	// b["two"] = 3
	// fmt.Printf("map a: %v\n", a)
	// fmt.Printf("map b: %v\n", b)

	//arrays and slices
	// array := [3]int{1,2,4}
	// array1 := array
	// fmt.Printf("array: %v\n", array)
	// fmt.Printf("array1: %v\n", array1)

	// fmt.Printf("&array: %v\n", &array)
	// fmt.Printf("&array1: %v\n", &array1)

	// array1[2] = 90
	// fmt.Printf("array: %v\n", array)
	// fmt.Printf("array1: %v\n", array1)

	// slice := []int{5,6,7,8}
	// slice1 := slice
	// fmt.Printf("slice: %v\n", slice)
	// fmt.Printf("slice1: %v\n", slice1)

	// fmt.Printf("&slice: %v\n", &slice)
	// fmt.Printf("&slice1: %v\n", &slice1)

	// slice1[2] = 88
	// fmt.Printf("&slice: %v\n", &slice)
	// fmt.Printf("&slice1: %v\n", &slice1)

	//structures with pointers
	// a := pointerStruct{
	// 	name: "Ka",
	// 	age: 21,
	// 	dob: "19990909",
	// }
	// fmt.Printf("a: %v, %T\n", a, a)
	// var b *pointerStruct
	// b = &pointerStruct{
	// 	name: "poo",
	// 	age: 12,
	// }
	// fmt.Printf("b: %v , %T\n", b,b)
	// c := new(pointerStruct)
	// fmt.Printf("c: %v,%T\n", c,c)
	// c.age = 21
	// fmt.Printf("c: %v,%T\n", c,c)
	// fmt.Printf("c: %v,%T\n", c.age,c.age)
	// d := new(pointerStruct)
	// fmt.Printf("d: %v,%T\n", d,d)
	// (*d).age = 21
	// fmt.Printf("*d: %v,%T\n", *d,*d)
	// fmt.Printf("d age: %v,%T\n", d.age,d.age)

	// a := [3] int {1,2,4}
	// fmt.Printf("a: %v, %T", a, a)
	// a1 := &a[0]
	// a2 := &a[1]
	// a3 := &a[2]
	// fmt.Printf("ap: %v, %T\n", a1, a1)
	// fmt.Printf("ap: %v, %T\n", a2, a2)
	// fmt.Printf("ap: %v, %T\n", a3, a3)

	// a := 20
	// fmt.Printf("a: %v, %T\n", a, a)
	// b := 20
	// fmt.Printf("b: %v, %T\n", b, b)
	// fmt.Printf("a == b: %v, %T\n", a == b, a == b)
	// c := a
	// fmt.Printf("c: %v, %T\n", c, c)
	// ap := &a
	// fmt.Printf("ap: %v, %T\n", ap, ap)
	// bp := &b
	// fmt.Printf("ap: %v, %T\n", bp, bp)
	// cp := &c
	// fmt.Printf("ap: %v, %T\n", cp, cp)
	// fmt.Printf("ap == bp: %v, %T\n", ap == bp, ap == bp)
	// fmt.Printf("bp == cp: %v, %T\n", bp == cp, bp == cp)
	// fmt.Printf("cp == ap: %v, %T\n", cp == ap, cp == ap)
	// var ap1 *int = &a
	// fmt.Printf("ap1: %v, %T\n", ap1, ap1)
	// fmt.Printf("&ap1: %v, %T\n", &ap1, &ap1)
	// fmt.Printf("*ap1: %v, %T\n", *ap1, *ap1)
	// fmt.Printf("a == ap1: %v, %T\n", a == *ap1, a == *ap1)
	// fmt.Printf("ap == ap1: %v, %T\n", ap == ap1, ap == ap1)

	// var bp1 *int = &b
	// fmt.Printf("bp1: %v, %T\n", bp1, bp1)
	// fmt.Printf("&bp1: %v, %T\n", &bp1, &bp1)
	// fmt.Printf("*bp1: %v, %T\n", *bp1, *bp1)

	// *bp1 = 90
	// fmt.Printf("bp1: %v, %T\n", bp1, bp1)
	// fmt.Printf("&bp1: %v, %T\n", &bp1, &bp1)
	// fmt.Printf("*bp1: %v, %T\n", *bp1, *bp1)
	// fmt.Printf("b: %v, %T\n", b, b)
	// fmt.Printf("&b: %v, %T\n", &b, &b)

	// a = 99
	// fmt.Printf("ap1: %v, %T\n", ap1, ap1)
	// fmt.Printf("&ap1: %v, %T\n", &ap1, &ap1)
	// fmt.Printf("*ap1: %v, %T\n", *ap1, *ap1)
	// fmt.Printf("a: %v, %T\n", a, a)
	// fmt.Printf("&a: %v, %T\n", &a, &a)

	// fmt.Println("Start")
	// defer fmt.Println("End")
	// defer func() {
	// 	fmt.Println("In Defer Function")
	// 	if r := recover(); r != nil{
	// 		fmt.Println("In Recover Block")
	// 	}
	// }()
	// myPanic()
	// defer fmt.Println("After Panic")

	// var a float64 = 37
	// b := int(a / 10)
	// c := b * 95000
	// fmt.Println(c)
	// d := math.Mod(a, 10)
	// fmt.Println(d)

	//defer
	//panic
	//recover
	// fmt.Println("start")
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("Something is fishy")

	// defer fmt.Println("Does this occur before panic")
	// fmt.Println("start")
	// panic("Something is fishy")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// fmt.Println("start")
	// panic("something is bad very bad")
	// fmt.Println("end")

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// fmt.Printf("res: %v, %T\n", res, res)
	// fmt.Printf("res.Body: %v, %T\n", res.Body, res.Body)
	// robots, err := ioutil.ReadAll(res.Body)
	// fmt.Printf("robots: %v, %T\n", robots, robots)
	// fmt.Printf("resClose: %v, %T\n", res, res)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", robots)

	// fmt.Println("Start")
	// fmt.Println("Middle")
	// fmt.Println("End")
	// fmt.Println("*******************")
	// fmt.Println("Start")
	// defer fmt.Println("Middle")
	// fmt.Println("End")
	// fmt.Println("*******************")
	// defer fmt.Println("Start")
	// defer fmt.Println("Middle")
	// defer fmt.Println("End")

	//loop
	// for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
	// 	fmt.Println(i, j)
	// }
	// s := []int{12, 23, 34}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }
	// statePopulation := map[string]int{
	// 	"India":  324345658997,
	// 	"USA":    237634756,
	// 	"China":  7342864892759,
	// 	"Africa": 6899787878,
	// }
	// for k, v := range statePopulation {
	// 	fmt.Println(k, v)
	// 	if k == "India" {
	// 		v = 1
	// 	}
	// 	fmt.Println(k, v)
	// }

	//control statements
	// if returnTrue() {
	// 	fmt.Println("In Block")
	// }
	// if true {
	// 	fmt.Println("In IF block")
	// }
	// statePopulation := map[string]int{
	// 	"India":  324345658997,
	// 	"USA":    237634756,
	// 	"China":  7342864892759,
	// 	"Africa": 6899787878,
	// }
	// if pop, ok := statePopulation["India"]; ok {
	// 	fmt.Println(pop)
	// 	fmt.Println(ok)
	// }

	//tags
	// t := reflect.TypeOf(Student{})
	// field, _ := t.FieldByName("name")
	// fmt.Println(field.Tag)
	// f := Student{
	// 	name: "poiuytrewq",
	// }
	// fmt.Println(f)

	//composition nth but inheritance
	// aStudents := []StudentClgDetails{
	// 	{
	// 		id:     12,
	// 		barnch: "CSE",
	// 		Student: Student{
	// 			name:        "Pooja",
	// 			age:         23,
	// 			phoneNumber: 1234567890,
	// 			parents:     []string{"Apple", "Iron"},
	// 		},
	// 	}, {
	// 		id:     12,
	// 		barnch: "ECE",
	// 		Student: Student{
	// 			name:        "Cheppanu",
	// 			age:         23,
	// 			phoneNumber: 1234567890,
	// 			parents:     []string{"Iron", "Coal"},
	// 		},
	// 	},
	// }
	// fmt.Println(aStudents)
	// aStudents := StudentClgDetails{}
	// aStudents.name = "Pooja"
	// aStudents.barnch = "CSE"
	// fmt.Println(aStudents)

	//collection types - maps, structs
	// aStudents := Student{
	// 	name:        "Pooja",
	// 	age:         21,
	// 	phoneNumber: 9908887654,
	// 	parents:     []string{"Apple", "Banana"},
	// }
	// aStdList := [...]Student{
	// 	{
	// 		name:        "Pooja",
	// 		age:         21,
	// 		phoneNumber: 9908887654,
	// 		parents:     []string{"Apple", "Banana"},
	// 	}, {
	// 		name:        "Honey",
	// 		age:         23,
	// 		phoneNumber: 9907687654,
	// 		parents:     []string{"Mango", "Banana"},
	// 	},
	// }
	// sSignleStd := struct{name string}{name: "IDontKnow"}
	// fmt.Println(aStudents)
	// fmt.Println(aStudents.parents)
	// fmt.Println(aStdList)
	// fmt.Println(aStdList[0].parents)
	// fmt.Println(sSignleStd)
	// aAnotherStd := sSignleStd
	// aAnotherStd.name = "John"
	// fmt.Println(sSignleStd)
	// fmt.Println(aAnotherStd)

	// statePopulation := make(map[string]int, 3)
	// statePopulation = map[string]int{
	// 	"India":  324345658997,
	// 	"USA":    237634756,
	// 	"China":  7342864892759,
	// 	"Africa": 6899787878,
	// }
	// fmt.Println()
	// fmt.Printf("%v, %T\n", statePopulation, statePopulation)
	// fmt.Printf("%v\n", statePopulation["India"])
	// statePopulation["Goa"] = 86373737
	// fmt.Printf("%v, %T\n", statePopulation, statePopulation)
	// delete(statePopulation, "Goa")
	// fmt.Printf("%v, %T\n", statePopulation, statePopulation)
	// fmt.Printf("%v\n", statePopulation["Goa"])
	// pop, ok := statePopulation["Goa"]
	// fmt.Println(pop, ok)
	// pop1, ok1 := statePopulation["India"]
	// fmt.Println(pop1, ok1)
	// sp := statePopulation
	// fmt.Println(sp)
	// delete(sp, "India")
	// fmt.Println(sp)
	// fmt.Println(statePopulation)

	//colection types - arrays, slices

	// slice1 := []int{1, 2, 4, 3, 5, 6, 7, 8, 9, 0}
	// slice2 := slice1[1:]
	// fmt.Printf("%v\n", slice2)
	// slice3 := slice1[:len(slice1)-1]
	// fmt.Printf("%v\n", slice3)
	// slice4 := append(slice1[:2], slice1[4:]...)
	// fmt.Printf("%v\n", slice4)
	// fmt.Printf("%v\n", slice1)

	// arr1 := [10]int{1, 2, 4, 3, 5, 6, 7, 8, 9, 0}
	// arr2 := arr1[1:]
	// fmt.Printf("%v\n", arr2)
	// arr3 := arr1[:len(arr1)-1]
	// fmt.Printf("%v\n", arr3)
	// arr4 := append(arr1[:2], arr1[4:]...)
	// fmt.Printf("%v\n", arr4)
	// fmt.Printf("%v\n", arr1)

	// arr1 := []int{}
	// fmt.Println(arr1)
	// fmt.Printf("Length: %v\n", len(arr1))
	// fmt.Printf("Capacity: %v\n", cap(arr1))
	// arr1 = append(arr1, 1)
	// fmt.Printf("%v\n", arr1)
	// fmt.Printf("Length: %v\n", len(arr1))
	// fmt.Printf("Capacity: %v\n", cap(arr1))
	// arr1 = append(arr1, 2, 3, 4, 5)
	// fmt.Printf("%v\n", arr1)
	// fmt.Printf("Length: %v\n", len(arr1))
	// fmt.Printf("Capacity: %v\n", cap(arr1))
	// arr1 = append(arr1, []int{6, 8, 7, 9}...)
	// fmt.Printf("%v\n", arr1)
	// fmt.Printf("Length: %v\n", len(arr1))
	// fmt.Printf("Capacity: %v\n", cap(arr1))

	// arr1 := make([]int, 3, 100)
	// fmt.Println(arr1)
	// fmt.Printf("Length: %v\n", len(arr1))
	// fmt.Printf("Capacity: %v\n", cap(arr1))

	// arr1 := [...]int{2, 4, 6, 8, 10, 12, 14, 16}
	// arr2 := arr1[:]
	// arr3 := arr1[3:]
	// arr4 := arr1[:3]
	// arr5 := arr1[3:6]
	// fmt.Printf("arr1: %v\n", arr1)
	// fmt.Printf("arr2: %v\n", arr2)
	// fmt.Printf("arr1: %v\n", arr3)
	// fmt.Printf("arr2: %v\n", arr4)
	// fmt.Printf("arr1: %v\n", arr5)
	// slice1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	// slice2 := slice1[:]
	// slice3 := slice1[3:]
	// slice4 := slice1[:3]
	// slice5 := slice1[3:6]
	// fmt.Printf("slice1: %v\n", slice1)
	// fmt.Printf("slice2: %v\n", slice2)
	// fmt.Printf("slice3: %v\n", slice3)
	// fmt.Printf("slice4: %v\n", slice4)
	// fmt.Printf("slice5: %v\n", slice5)

	// a := []int{1, 3, 5}
	// b := a
	// b[1] = 9
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)

	// a := [...]int{1, 3, 5}
	// b := a
	// b[1] = 9
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", a)

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1,0,0}
	// identityMatrix[1] = [3]int{0,1,0}
	// identityMatrix[2] = [3]int{0,0,1}
	// fmt.Printf("%v", identityMatrix)

	// grades := [...]int{23, 435, 67, 85, 323, 22}
	// fmt.Printf("%v\n", grades)
	// fmt.Printf("%v", len(grades))

	// var index int = 2
	// fmt.Scanf("%v", &index);

	// grades := [3]int{98, 34, 21}
	// fmt.Printf("%v\n", grades)
	// fmt.Printf("%v\n", grades[0])
	// fmt.Printf("%v\n", index)

	//constants
	// const myConst int =  42
	// fmt.Printf("%v, %T\n",myConst, myConst)

	// fmt.Printf("%v\n", a)
	// fmt.Printf("%v\n", b)
	// fmt.Printf("%v\n", c)

	// fmt.Printf("%v\n", a  + 2)

	// var f float32 = 2344
	// fmt.Printf(`%.2f`, f)

	// runes
	r := "apple banana"
	b := []byte(r)
	fmt.Printf("%v, %T\n", b, b)
	c := string(b)
	fmt.Printf("%v, %T\n", c, c)
	b[3] = 'g'
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", b, b)

	// Text Types
	s := "I am Pooja"
	b1 := []byte(s)
	fmt.Printf("%v, %T\n", b1, b1)
	// s[3] = 'q'
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", s, s)
	// s := "I am Pooja"
	// s1 := " No"
	// fmt.Printf("%v, %T\n", s + s1, s + s1)
	// fmt.Printf("%v %T\n", s, s)
	// fmt.Printf("%v %T\n", s[2], s[2])
	// fmt.Printf("%v %T\n", string(s[2]), s[2])

	//Complex Number
	// var n complex64 = 1 + 2i
	// fmt.Printf("%v %T", n, n)
	// a := 8 + 2i
	// fmt.Printf("%v, %T\n", real(a), real(a))
	// fmt.Printf("%v, %T\n", imag(a), imag(a))
	// b := 3 + 5i
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// var n complex128 = complex(3, 89)
	// fmt.Printf("%v, %T\n", n,n)

	//Floating - no bitwise operators
	// var n float32 = 3.14
	// fmt.Printf("%g %T", n, n)
	// a := 8.4
	// b := 3.5
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	//Integers
	//Signed Integers
	// var n uint16 = 42
	// fmt.Printf("%v %T\n", n,n)

	// a := 8
	// b := 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(b / a)
	// fmt.Println(a % b)

	// var a int = 8
	// var b int8 = 10
	// fmt.Println(a + int(b))

	// fmt.Println(a | b)
	// fmt.Println(a & b)
	// fmt.Println(a ^ b)
	// fmt.Println(a &^ b)

	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	//Boolean Types
	// var n bool
	// fmt.Printf("%v %T\n", n, n)

	// var i int
	// fmt.Printf("%v %T\n", i, i)

	// var j string
	// fmt.Printf("%v %T\n", j, j)

	// var k float32
	// fmt.Printf("%v, %T\n", k, k)

	// var n bool = true
	// fmt.Printf("%v, %T", n, n)

	// n1 := 1 == 1
	// m1 := 1 == 2
	// fmt.Printf("%v, %T\n", n1, n1)
	// fmt.Printf("%v, %T", m1, m1)

	// Variables
	// var i int
	// i = 100
	// fmt.Println(i)

	// var j float32 = 98.65
	// fmt.Println(j)

	// k := 89.66
	// fmt.Println(k)

	// fmt.Printf("%v, %T \n", j, j)
	// fmt.Printf("%v, %T \n", k, k)

	// fmt.Printf("%v, %T \n", o, o)
	// var o int = 76
	// fmt.Printf("%v, %T \n", o, o)
	// o := 1
	// fmt.Printf("%v, %T", o, o)
	// fmt.Printf("%v, %T", std_id, std_id)

	// q := "pooja" //unused variables should not be exists
	// fmt.Printf("%v, %T", o, o)

	//Type Conversion
	// var i int = 42
	// fmt.Printf("%d, %T\n", i, i)
	// var j float32
	// j = float32(i)
	// fmt.Printf("%g, %T\n", j, j)

	// var k string
	// k = string(i)
	// fmt.Printf("%d %T\n", k, k)

	// var t string
	// t = strconv.Itoa(i)
	// fmt.Printf("%v %T\n", t, t)
	// d, err := strconv.Atoi("-42")
	// fmt.Printf("%v, %T, %v, %T", d, d, err, err)
	// d, err := strconv.ParseFloat("3.1415", 64)
	// fmt.Printf("%v, %T, %v, %T", d, d, err, err)
	// ------------------------------------------
	// Types of print statements in Go
	// fmt.Println("Good Morning")
	// fmt.Println("Hello Go")

	// fmt.Print("Checking")
	// fmt.Print("Print Statement")

	// fmt.Printf("CHecking What")
	// fmt.Printf("PrintF is going on")
}
