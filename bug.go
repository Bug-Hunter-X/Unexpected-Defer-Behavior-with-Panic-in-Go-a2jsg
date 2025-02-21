func main() {
  fmt.Println("This is a sample Go program.")
  defer fmt.Println("This line will be printed after the function completes.")
  panic("This will cause a panic!")
}