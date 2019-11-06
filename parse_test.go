package JsonORM

import "testing"

func TestParse(t *testing.T) {
	jsonORMStr := `
{
	"table": "user",
    "with": [
        {
            "table": "user_extend",
            "relationship": "one",
            "condition": "id,user_id"
        },{
            "table": "user_login_log",
            "relationship": "many",
            "condition": "id,user_id"
        }
    ],
	"where": [
		["sex", 1],
		["id", ">=", 1],
		["id", "=", 1],
		["id", "=<", 1],
		["id", "!=", 1],
		["id", "!>=", 1],
		["id", "!=<", 1],
		["nickname", "like", "%name%"],
		["id", "in", 1],
		["id", "not in", 1]
	],
	"orWhere": [
		["id", ">=", 1314]
	],
	"select": "id,name,sex,created_at",
	"offset": 0,
	"limit": 10,
	"orderBy": "id,desc",
	"handle": "get" 
}`
	err := Parse(jsonORMStr)
	if err != nil {
		t.Fatal(err)
	}
}
