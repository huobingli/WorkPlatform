package main
import (
  "fmt"
)
func main() {
  m1 := make([]string, 1)
  m1[0] = "test"
  fmt.Println("调用 func1 前 m1 值:", &m1[0])
  fmt.Println("调用的地址：", &m1[0])
  func1(m1)
  fmt.Println("调用 func1 后 m1 值:", &m1[0])
}
func func1 (a []string) {
  a[0] = "val1"
  fmt.Println("func1中:", &a)
  fmt.Println("func1中地址:", &a[0])
}