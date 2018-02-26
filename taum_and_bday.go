package main
import "fmt"

func main() {
    var T int
    fmt.Scanf("%d", &T)
    for i:=0;i<T;i++ {
        var B,W,X,Y,Z int
        fmt.Scanf("%d%d\n%d%d%d", &B,&W,&X,&Y,&Z)
        //fmt.Println(B,W,X,Y,Z)
        var t int
        switch {
            case X+Z < Y:
                t = B*X + W*(X+Z)
            case Y+Z < X:
                t = W*Y + B*(Y+Z)
            default:
                t = B*X+W*Y
        }
        fmt.Println(t)
    }
}
