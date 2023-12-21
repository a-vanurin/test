package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // f, err := os.Open("./D/d.txt")
    // if err != nil {
    //     return
    // }
    // defer f.Close()
    // in := bufio.NewReader(f)

    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var t int
    fmt.Fscan(in, &t)
    for i := 0; i < t; i++ {
        var s string
        fmt.Fscan(in, &s)

        ll := []string{""}
        var x, y int

        for _, cmd := range s {
            switch cmd {
            case 'L':
                if x > 0 {
                    x--
                }
            case 'R':
                if x < len(ll[y]) {
                    x++
                }
            case 'U':
                if y > 0 {
                    y--
                    x = min(x, len(ll[y]))
                }
            case 'D':
                if y < len(ll)-1 {
                    y++
                    x = min(x, len(ll[y]))
                }
            case 'B':
                x = 0
            case 'E':
                x = len(ll[y])
            case 'N':
                enter := ll[y][x:]
                ll[y] = ll[y][:x]
                ll = append(ll[:y+1], append([]string{enter}, ll[y+1:]...)...)
                x = 0
                y++
            default:
                ll[y] = ll[y][:x] + string(cmd) + ll[y][x:]
                x++
            }
        }

        for _, l := range ll {
            fmt.Fprintln(out, l)
        }
        fmt.Fprintln(out, "-")
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
