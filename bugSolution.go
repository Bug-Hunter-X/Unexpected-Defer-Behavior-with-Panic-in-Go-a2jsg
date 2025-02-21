func main() {
  fmt.Println("This is a sample Go program.")
  defer func() {
    if r := recover(); r != nil {
      fmt.Printf("Recovered from panic: %v\n", r)
    }
    fmt.Println("This line will be printed after the function completes, even with panic.")
  }()
  panic("This will cause a panic!")
} 