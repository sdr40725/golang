package main

import (
	"fmt"
	"strings"
)

func getUserListSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}

	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}

	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}

	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("email = '%d'", opts.sexy))
	}

	return sql + " where " + strings.Join(where, " or ")
}

type searchOpts struct {
	username string
	email    string
	sexy     int
}

func main() {
	println(getUserListSQL(searchOpts{
		username: "test001",
	}))

	println(getUserListSQL(searchOpts{
		username: "test002",
		email:    "test002@gmail.com",
		sexy:     2,
	}))
}
