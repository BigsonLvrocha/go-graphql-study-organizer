package main

import (
	"fmt"
	"strconv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/pkg/errors"
)

func gqlIDToUint(i graphql.ID) (uint, error) {
	r, err := strconv.ParseInt(string(i), 10, 32)
	if err != nil {
		return 0, errors.Wrap(err, "GqlIDToUint")
	}

	return uint(r), nil
}

func gqlIDP(id uint) *graphql.ID {
	r := graphql.ID(fmt.Sprint(id))
	return &r
}
