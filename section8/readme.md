
# Public / Privateの扱い

```go
package sample
var Public string = "Public"
var Private string = "Private"
```

上記のpackageをimportした場合、
Publicは参照可能だが、Privateは参照不可となる。

# テストの実行方法

```bash
# "..."でそのディレクトリ配下のユニットテストコードを探査＆実行してくれる
go test ./...
# カバレッジつき
go test -v ./...
```

# godocの表示

```bash
godoc -http=:6060
```

で起動。  

`http://localhost:6060`でアクセス可能。  
