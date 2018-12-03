
このセクションで作成したソースでやっていることを記載する。

# [sample.go](./sample.go)

## structにmethodを定義する 

```go
type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
  v := Vertex{3, 4}
  v.Scale(100)
	fmt.Println(v.Area())
}
```


