package qbuilder

import (
	"regexp"
	"testing"
)

func TestSelect(t *testing.T) {
	sel := SelectBuilder{}
	sql := sel.Select().
		From("\"user\"").
		Alias("u").
		InnerJoin("address", "a", "a.user_id = u.id").
		Where("a.street LIKE '%brod%'").
		AndWhere("a.house > 12").
		OrderBy(Sort{"a.flat": SortDirectionASC}).
		Having("a.flat > 12").
		GroupBy("a.flat").
		For(LockModeUpdateNowait).
		ToSQL()
	re := regexp.MustCompile("\\s{2,}")
	println(re.ReplaceAllString(sql, " "))
	sqlE := "SELECT * FROM \"user\" AS u INNER JOIN address AS a ON a.user_id = u.id WHERE ((a.street LIKE '%brod%') AND (a.house > 12)) GROUP BY a.flat HAVING a.flat > 12 ORDER BY a.flat ASC FOR UPDATE NOWAIT"
	if sqlE != re.ReplaceAllString(sql, " ") {
		t.Error("Select failed")
	}
}
