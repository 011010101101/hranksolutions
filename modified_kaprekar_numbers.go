package main
import "fmt"
import "strconv"

func main() {
    var min,max int
    fmt.Scanf("%d\n%d", &min, &max)
    var found bool
    for i:=min;i<=max;i++ {
        s := strconv.Itoa(i*i)
        s1 := s[:len(s)/2]
        s2 := s[len(s)/2:]
        i1, _ := strconv.Atoi(s1)
        i2, _ := strconv.Atoi(s2)
        if i1+i2 == i {
            found = true
            fmt.Print(i, " ")
        }
    }
    if !found {
        fmt.Println("INVALID RANGE")
    }
}
