# vscode-go-inline-sql
Syntax highlighting for inline sql in go.

Raw string literals that start with `--sql\n` will be highlighted as sql.

```go
sql := `--sql
	SELECT * FROM users;
`
```
