package main

import "C"

import (
    "fmt"
    "github.com/buptbill220/c-call-test/lib"
 "sort"
 "sync"
 "math"
)
var count int
var mtx sync.Mutex
//export Add
func Add(a, b int) int { return a + b }
//export Cosine
func Cosine(x float64) float64 { return math.Cos(x) }
//export Sort
func Sort(vals []int) { sort.Ints(vals) }
//export Log
func Log(msg string) int {
  mtx.Lock()
  defer mtx.Unlock()
  fmt.Println(msg)
  count++
  return count
}

var l = &lib.ListTest{}
//export Foo
func Foo(a, b int) int {
    return a + b;
}

//export Bar
func Bar() {
    fmt.Println("I am bar, not foo!")
}

//export SetMap
func SetMap(key []string, val []int) {
    l.SetMap(key, val)
}

//export SetData
func SetData(data []int64) {
    l.SetData(data)
}

//export Print
func Print() {
    l.Print()
}
func main() {}
