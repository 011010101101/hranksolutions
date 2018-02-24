package main
import "fmt"

func main() {
    var h,m int
    fmt.Scanf("%d\n%d", &h,&m)
    hl := map[int]string{
        1: "one",
        2: "two",
        3: "three",
        4: "four",
        5: "five",
        6: "six",
        7: "seven",
        8: "eight",
        9: "nine",
        10: "ten",
        11: "eleven",
        12: "twelve",
        13: "thirteen",
        14: "fourteen",
        15: "fifteen",
        16: "sixteen",
        17: "seventeen",
        18: "eighteen",
        19: "ninteen",
        20: "twenty",
        21: "twenty one",
        22: "twenty two",
        23: "twenty three",
        24: "twenty four",
        25: "twenty five",
        26: "twenty six",
        27: "twenty seven",
        28: "twenty eight",
        29: "twenty nine",
        31: "thirty one",
        32: "thirty two",
        33: "thirty three",
        34: "thirty four",
        35: "thirty five",
        36: "thirty six",
        37: "thirty seven",
        38: "thirty eight",
        39: "thirty nine",
        41: "forty one",
        42: "forty two",
        43: "forty three",
        44: "forty four",
        45: "forty five",
        46: "forty six",
        47: "forty seven",
        48: "forty eight",
        49: "forty nine",
        51: "fifty one",
        52: "fifty two",
        53: "fifty three",
        54: "fifty four",
        55: "fifty five",
        56: "fifty six",
        57: "fifty seven",
        58: "fifty eight",
        59: "fifty nine",
    }
    switch {
        case m == 0:
        fmt.Print(hl[h]," o' clock")
        case m == 30:
        fmt.Print("half past ",hl[h])
        case m == 45:
        fmt.Print("quarter to ",hl[(h+1)%12])
        case m == 15 :
        fmt.Print("quarter past ",hl[h])
        case m == 1 :
        fmt.Print("one minute past ",hl[h])
        case m < 30:
        fmt.Print(hl[m]," minutes past ",hl[h])
        default:
        fmt.Print(hl[60-m]," minutes to ",hl[(h+1)%12])
    }
}
