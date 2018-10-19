package lib

import (
        "fmt"
)

type ListTest struct {
    data []int64
    md   map[string]int
    str  string
}

func (l *ListTest) Print() {
    fmt.Printf("data len %d, map len %d, str len %d\n", len(l.data), len(l.md), len(l.str))
    for _, d := range l.data {
        fmt.Printf("%d \t", d)
    }
    fmt.Println("\n======\n")
    for k, v := range l.md {
        fmt.Printf("%s : %d\n", k, v)
    }
    fmt.Println("\n=======\n")
    fmt.Printf("%s\n", l.str)
}

func (l *ListTest) SetData(data []int64) {
    l.data = make([]int64, len(data))
    copy(l.data, data)
}

func (l *ListTest) SetMap(key []string, val []int) {
    fmt.Printf("set map key len %d, val len %d\n", len(key), len(val))
    l.md = make(map[string]int, len(key))
    for idx, k := range key {
        l.md[k] = val[idx]
        fmt.Printf("key %s, val %d\n", k, val[idx])
    }
    l.str = key[0]
}
