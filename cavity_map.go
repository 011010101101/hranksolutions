package main
import "fmt"
import "bufio"
import "os"

func main() {
    var n int
    fmt.Scanf("%d", &n)
    r := bufio.NewReader(os.Stdin)
    var board [][]int
    for i:=0;i<n;i++ {
        var row []int
        for j:=0;j<n;j++ {
            v, _ := r.ReadByte()
            x := v - 48
            row = append(row, int(x))
        }
        r.ReadByte()
        board = append(board, row)
    }
    
    isCavity := make(map[[2]int]bool)
    for x:=1;x<n-1;x++ {
        for y:=1;y<n-1;y++ {
            value := board[y][x]
            n := board[y-1][x]
            s := board[y][x-1]
            e := board[y][x+1]
            w := board[y+1][x]
            //fmt.Println(value, n, s, e, w)
            if value > n && value > e &&
               value > w && value > s {
                   isCavity[[2]int{x,y}] = true
            }
        }
    }
    for y:=0;y<n;y++ {
        for x:=0;x<n;x++ {
            if isCavity[[2]int{x,y}] {
                fmt.Print("X")
            } else {
                fmt.Print(board[y][x])
            }
        }
        fmt.Println()
    }
}

func abs(x, y int) int {
    if x > y {
        return x-y
    } 
    return y-x
}
