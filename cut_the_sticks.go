package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    var x []int
    for i:=0;i<n;i++ {
        var v int
        fmt.Scanf("%d", &v)
        x = append(x,v)
    }
    fmt.Println(len(x))
    //fmt.Println(x)
    for {
        min := 9999999999
        for _, v := range x {
            if v > 0 && v < min {
                min = v
            }
        }
        var count int
        for i :=range x {
            x[i] -= min
            if x[i] > 0 {
                count += 1
            }
        }
        if count == 0 {
            break
        }
        //fmt.Println(x)
        fmt.Println(count)
    }
}
