package main

import "fmt"

func tugasPalindrome(input string) bool {

  for i := 0; i< len(input)/2; i++ {
    if input[i] != input[len(input)-i-1]{
      return false
    }
  }
	return true
}

func main() {
  fmt.Println("silahkan masukan teks/Angka Palindrome dibawah ini :")
  var str string
  fmt.Scan(&str)
	result := tugasPalindrome(str)
	if result == true {
    fmt.Println("True, itu adalah palindrome")
  } else {
    fmt.Println("False, itu bukan palindrome")
  }
}
