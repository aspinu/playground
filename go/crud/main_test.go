package main

import (
	"database/sql"
	"testing"

	"github.com/dixonwille/wmenu"
)

func Test_handleFunc(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db   *sql.DB
		opts []wmenu.Opt
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleFunc(tt.db, tt.opts)
		})
	}
}
