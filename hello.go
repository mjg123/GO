package main

import "fmt"

type LLNode struct {
  val  int
  next *LLNode
}

type LList struct {
  head *LLNode
}

func main() {

  t := new(LLNode)
  t.val = 4

  fmt.Printf("%d\n", t.val)
}
