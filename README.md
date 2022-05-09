### Golang sql query builder for postgres

qbuilder - is a simple sql string builder implementation tested with postgres.
All string arguments are included in the final result unchanged as is.

Feel free to make a PR

Simple SELECT example

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Select("*").
        From("users").
        Where("age > $1").
        ToSQL()
        
Output: SELECT * FROM "users" WHERE age > $1;
```

Example SELECT column list

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Select("id, age, email, first_name").
        From("users").
        Where("age > $1").
        ToSQL()

Output: SELECT id, age, email, first_name FROM "users" WHERE age > $1;
```

Example with postgres native function, pagination and sorting

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Select().   // Empty list interpreted as *
        From("users").
        Where("LOWER(first_name) ~ $1").
        OrderBy(qbuilder.Sort{
            "created_at": qbuilder.SortDirectionDESC,
            "first_name": qbuilder.SortDirectionASC,
        }).
        Limit(1).
        Offset(1).
        ToSQL()

Output: SELECT * FROM "users" WHERE LOWER(first_name) ~ $1 ORDER BY created_at DESC, first_name ASC;
```

Example INSERT

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Insert("users").
		Columns("email, first_name, last_name, created_at").
		Value("$1, $2, $3, NOW()").
		Returning("id").
		ToSQL()


Output: INSERT INTO "users" (email, first_name, last_name, created_at) VALUES ($1, $2, $3, NOW());
```

Example INSERT by 'db' tags, including anonymous embed structs fields 

```golang
import "github.com/slmder/qbuilder"

type User struct {
	Email           string     `db:"email"`
	FirstName       string     `db:"first_name"`
	LastName        string     `db:"last_name"`
	CreatedAt       time.Time  `db:"created_at"`
}
user := User{}
sql := qbuilder.Insert("users").RowE(user).ToSQL()

Output: INSERT INTO "users" (email, first_name, last_name, created_at) VALUES (:email, :first_name, :last_name, :created_at);
```

Example UPDATE 

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Update("users").
        SetMap(map[string]string{
            "first_name": "$1",
            "last_name":  "$2",
        }).
        Where("id = $1").
        ToSQL()

Output: UPDATE "users" SET first_name = $1, last_name = $2 WHERE id = $1;
```
Example UPDATE by 'db' tags

```golang
import "github.com/slmder/qbuilder"

type User struct {
	Email           string     `db:"email"`
	FirstName       string     `db:"first_name"`
	LastName        string     `db:"last_name"`
	CreatedAt       time.Time  `db:"created_at"`
}
user := User{}
sql := qbuilder.Update("users").
        SetMapE(user).
        Where("id = $1").
        ToSQL()

Output: UPDATE "users" SET email = :email, first_name = :first_name, last_name = :last_name, created_at = :created_at WHERE id = $1;
```

Example DELETE 

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.Delete("users").Where("id = :id").ToSQL()

Output: DELETE FROM "users" WHERE id = $1;
```

Complex sql with sub selects, union and CTE.
All builder have ability to use CTE.

```golang
import "github.com/slmder/qbuilder"

sql := qbuilder.
    Select("up.*").
    From("user_permissions", "up").
    WithRecursive("user_groups",
        qbuilder.Select("group_id").
            From("user_group").
            Where("user_id = $1").
            ToSQL()).
    With("user_permissions",
        qbuilder.Select("p.*").
            From("acl_permission", "p").
            InnerJoin("group_permission", "u", "p.id = u.permission_id").
            Wheref("u.group_id IN (%s)", qbuilder.Select("group_id").From("user_groups").ToSQL()).
            Union(qbuilder.Select("p.*").
                From("acl_permission", "p").
                InnerJoin("user_permissions", "up", "p.parent_id = up.id").
                ToSQL()).
            ToSQL()).
    ToSQL()

Output: WITH RECURSIVE user_groups AS (
            SELECT group_id 
                FROM user_group 
                WHERE (user_id = '$1')
            ),
            user_permissions AS (
            SELECT p.*
                FROM acl_permission AS p
                     INNER JOIN group_acl_permission AS u ON p.id = u.permission_id
                WHERE (u.group_id IN (SELECT group_id FROM user_groups))
                UNION
                (SELECT p.* FROM acl_permission AS p)
                UNION ALL
                (SELECT p.* FROM acl_permission AS p)
            )
       SELECT up.* FROM user_permissions AS up;
```
