# monkey

- 書籍『Go言語でつくるインタプリタ』の演習。

## REPL起動

```go
go run main.go
Hello ktny! This is the Monkey programming language!
Feel free to type in commands
>> let a = 1 + 2
>> a
3
```

## テスト

```go
go test ./lexer ./parser ./evaluator
```
