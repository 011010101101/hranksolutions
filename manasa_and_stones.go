package main
import "fmt"
import "sort"

func main() {
    var T int
    fmt.Scanf("%d", &T)
    for i:=0;i<T;i++ {
        var n,a,b int
        fmt.Scanf("%d\n%d\n%d", &n, &a, &b)
        //fmt.Println(n,a,b)
        search(n, a, b)
    }
}

func search(n, a, b int) {
    var search_ func(c []int, v, n int)
    m := make(map[int]bool)
    memo := make(map[[2]int]bool)
    search_ = func(c []int, v, n int) {
        if memo[[2]int{v,n}] {
            return
        }
        memo[[2]int{v,n}] = true
        c = append(c, v)
        if n == 1 {
            m[v] = true
            //fmt.Println(c)
            return
        }
        search_(c, v+a, n-1)
        search_(c, v+b, n-1)
    }
    search_([]int{}, 0, n)
    var v []int
    for k := range m {
        v = append(v, k)
    }
    sort.Ints(v)
    for _, v := range v {
        fmt.Print(v, " ")
    }
    fmt.Println()
}

func abs(x, y int) int {
    if x > y {
        return x-y
    } 
    return y-x
}
