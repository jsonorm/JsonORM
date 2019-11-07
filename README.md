# JSON ORM
JSON ORM is a json orm specification.

# idea

- table
- with
- relationship 
- condition
- join
- where
- orWhere and orwhere
- select
- offset
- limit
- order
- handle or func or get

```json
{
	"table": "user",
    "with": [
        {
            "table": "user_extend",
            "relationship": "one",
            "condition": "id,user_id",
            "with": [
                {
                    "table": "user_extend",
                    "relationship": "one",
                    "condition": "id,user_id"
                }
            ]
        },{
            "table": "user_login_log",
            "relationship": "many",
            "condition": "id,user_id"
        }
    ],
    "join": [
      ["user_auth","user.id","user_auth.user_id"],
      ["user_login_log","user.id","user_login_log.user_id"]
    ],
	"where": [
		["id", 1],
		["id", ">=", 1],
		["id", "=", 1],
		["id", "=<", 1],
		["id", "!=", 1],
		["id", "!>=", 1],
		["id", "!=<", 1],
		["id", "like", 1],
		["id", "in", 1],
		["id", "not in", 1]
	],
	"orWhere": [
		["id", 1],
		["id", ">=", 1],
		["id", "=", 1],
		["id", "=<", 1],
		["id", "!=", 1],
		["id", "!>=", 1],
		["id", "!=<", 1],
		["id", "like", 1],
		["id", "in", ["12",123]],
		["id", "not in", ["12",123]]
	],
	"select": ["id,name,sex,created_at","user_auth.*"],
	"offset": 0,
	"limit": 10,
	"orderBy": "id:desc;created_at:desc",
	"func": "get" // get first delete update create
}
```
