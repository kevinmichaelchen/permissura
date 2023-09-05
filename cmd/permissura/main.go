package main

import "github.com/kevinmichaelchen/permissura/internal/cli"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.Main(cli.LDFlags{
		Version: version,
		Commit:  commit,
		Date:    date,
	})
}
