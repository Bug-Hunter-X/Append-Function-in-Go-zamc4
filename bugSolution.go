func main() {

        x := make([]int, 3, 10) // Pre-allocate capacity to 10

        fmt.Println(x)  // Output: [0 0 0]

        x = append(x, 10)

        fmt.Println(x)  // Output: [0 0 0 10]

        x = append(x, 11, 12, 13, 14)

        fmt.Println(x)  // Output: [0 0 0 10 11 12 13 14]

}        