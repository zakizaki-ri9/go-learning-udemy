
このセクションで作成したソースでやっていることを記載する。

# [01-pointer.go](./01-pointer.go)

## structにmethodを定義する 

```go
type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Area())
}
```


