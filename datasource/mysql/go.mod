module github.com/llyb120/bingo/datasource/mysql

go 1.24.5

require (
	github.com/go-sql-driver/mysql v1.9.3
	github.com/llyb120/bingo/config v0.0.0
	github.com/llyb120/bingo/core v0.0.0
)

replace (
	github.com/llyb120/bingo/config => ../../config
	github.com/llyb120/bingo/core => ../../core
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/petermattis/goid v0.0.0-20250721140440-ea1c0173183e // indirect
)
