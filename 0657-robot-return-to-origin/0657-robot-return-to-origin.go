func judgeCircle(moves string) bool {
    x, y := 0, 0
    for _, move := range moves{
        switch move {
            case 'U':
                x++
            case 'D':
                x--
            case 'R':
                y++
            case 'L':
                y--
        }
    }
    return x == 0 && y == 0
}