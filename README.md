# hisuru
Go program that compares software versions in the x.y.z format

Usage:
```golang
  x := "1"
  y := "1.0.1"
  z := "1.0"
  fmt.Println(hisuru.Compare(x, y)) // 1 = y
  fmt.Println(hisuru.Compare(x, z)) // 0 = equal
  fmt.Println(hisuru.Compare(x, a)) // -1 = x
```
