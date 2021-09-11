package main

import (  
    "fmt"
)

func addNumber(a int,b int) int {
    return a+b
}

func rectProps(length, width float64) (float64, float64, float64) {  
    var area = length * width
    var perimeter = (length + width) * 2
    var test = (length * width / 2)
    return area, perimeter,test
}

func main() {  
    // fmt.Println("Hello World")
    // name:= "Paradox"
    // var age int = 25
    // a, b := 145.8, 543.8
    // var c = math.Min(a,b)

    // fmt.Println("my name is", name,"a=>", a, "b=>", b, "min=>", c ," and my age is", age)

    // c1 := complex(5, 7)
    // c2 := 8 + 27i
    // cadd := c1 + c2
    // fmt.Println("C===>", c1)
    // fmt.Println("sum:", cadd)
    // cmul := c1 * c2
    // fmt.Println("product:", cmul)


    // first := "D"
    // last := "dox"
    // name := first +" "+ last
    // fmt.Println("My name is",name)


    // i := 55      //int
    // j := 67.8    //float64
    // sum := float64(i) + j //int + float64 not allowed
    // fmt.Println(sum)


    // var defaultName string = "Sam" //allowed
    // type myString string
    // var customName myString = "Sam" //allowed
    // customName = defaultName //not allowed



    // const a = 5
    // var intVar int = a
    // var int32Var int32 = a
    // var float64Var float64 = a
    // var complex64Var complex64 = a
    // fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)

    // var a = 40.2 % 8
    // fmt.Printf("a's type is %T and value is %v", a, a)

    // arr := [6]int{1,2,3,4,5,6}
    // for i,x:=range arr {
    //     fmt.Printf("a[%d] ==> %d \n",i,x)
    // }


    // sum of two number

    // a,b := 0,0
    // fmt.Printf('Enter First Number')
    // fmt.Scan(&a)
    // fmt.Printf('Enter Second Number')
    // fmt.Scan(&b)
    // var c int = addNumber(5,3)
    // var add, sub, mul, div = application(5,3)
    // fmt.Println(add)
    // fmt.Println(sub)
    // fmt.Println(div)
    area, _, test := rectProps(10.8, 5.6) // perimeter is discarded
    fmt.Printf("Area %f ===> %f ", area, test)
}


