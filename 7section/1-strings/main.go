package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)

	b := strings.Builder{}
	b.Write([]byte("Here is an example"))
	b.WriteString("\n")

	fmt.Println(b.String())

	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.Title(s1)) // deprecated
	s3 := "     test sss     "
	fmt.Println("s3:", s3, len(s3))
	s3 = strings.TrimSpace(s3)

	fmt.Println("s3 after trim:", s3, len(s3))

	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))
	fmt.Println(strings.Replace("test@test.com", "test", "john", 1)) // replace the first 1 "test" to "john"

	parts := strings.Split("jane@example.com", "@") // Split returns a slice
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jane example.com")
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ","))
}
