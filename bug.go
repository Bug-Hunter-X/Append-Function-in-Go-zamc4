func main() {


        x := make([]int, 3, 5)

        fmt.Println(x)

        x = append(x, 10)

        fmt.Println(x)

        x = append(x, 11, 12, 13, 14)

        fmt.Println(x)

}