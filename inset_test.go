package qbuilder

import (
	"testing"
)

func TestInsert(t *testing.T) {
	insert := InsertBuilder{}
	sql := insert.Insert("users").
		Row(map[string]string{
			"surname": "Sena",
		}).
		Returning("username").
		OnConflict().
		OnTarget("username").
		DoUpdate(map[string]string{
			"surname": "EXCLUDED.surname",
		}).
		ToSQL()
	if sql != "INSERT INTO users (surname) VALUES (Sena) ON CONFLICT (username) DO UPDATE SET \"surname\" = EXCLUDED.surname RETURNING username" {
		t.Error("insert failed")
	}
}
