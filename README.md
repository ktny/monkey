# monkey

- 書籍『Go言語でつくるインタプリタ』の演習。

## 実装

下記を実現する字句解析器、構文解析器、評価器。

- 数値
- 真偽値
- 文字列
- 配列
- ハッシュ
- if式
- let文
- return文
- 関数
- マクロ

## REPL

```go
go run main.go
Hello ktny! This is the Monkey programming language!
Feel free to type in commands
>> let five = 5;
>> let ten = 10;
>> let add = fn(x, y) { x + y };
>> add(five, ten);
```

## テスト

```sh
$ go test ./...
?       github.com/ktny/monkey  [no test files]
ok      github.com/ktny/monkey/ast      (cached)
ok      github.com/ktny/monkey/evaluator        (cached)
ok      github.com/ktny/monkey/lexer    (cached)
ok      github.com/ktny/monkey/object   (cached)
ok      github.com/ktny/monkey/parser   (cached)
?       github.com/ktny/monkey/repl     [no test files]
?       github.com/ktny/monkey/token    [no test files]
```

## 感想

[『Go言語でつくるインタプリタ』を読んだ - Kattsu Sandbox](https://katsusand.dev/posts/review-go-interpreter/)
