package main

import "fmt"

func main() {

	var nums  = []int{2,3,4}
	var sum  = 0
	for _, i2 := range nums {
		sum += i2
	}
	fmt.Println("sum:",sum)

	for i, i2 := range nums {
		if i2 == 3 {
			fmt.Println("index:",i)
		}
	}

	kvs := map[string]string{"a":"apple","b":"ban"}

	for s2, s3 := range kvs {
		fmt.Printf("%s -> %s\n",s2,s3)
	}

	for s2 := range kvs {
		fmt.Println("key:",s2)
	}

	for i, i2 := range "go" {
		fmt.Println(i,i2 )
	}

}
