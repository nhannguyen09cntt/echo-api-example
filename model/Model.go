package models

import (
	"fmt"
	"errors"
	"strings"
	"database/sql"
	_"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
    DB = Init()
}

func getConnection() *sql.DB {
	return DB;
}

/*Create mysql connection*/
func Init() *sql.DB {
	db, err := sql.Open("mysql", "forum_vn:forum_vn_p@ssword@tcp(178.128.93.55:3306)/forum_vn")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	return db
}
/*end mysql connection*/


func getTable(table string) string {
	return table
}

func LeftJoinSqlBuild(tables []string, on []map[string]string, fields map[string][]string, p, listRows int, orderBy []string, groupBy []string, condition string) (sql string, err error) {
	if len(tables) < 2 || len(tables)-1 != len(on) {
		err = errors.New("The parameter is not standardized: the number of data tables of the joint query must be 2 or more, and the number of tables is one more than the on condition.")
		return
	}
	var (
		FieldSlice   []string
		StrOrderBy   string
		StrGroupBy   string
		StrCondition string
		joinKV       string
		join         = []string{tables[0]}
		usedTables   = []string{}
	)
	for table, field := range fields {
		for _, f := range field {
			FieldSlice = append(FieldSlice, strings.Trim(fmt.Sprintf("%v.%v", table, f), "."))
		}
	}
	for index, table := range tables {
		slice := strings.Split(strings.TrimSpace(table), " ")
		if len(slice) == 1 {
			slice = append(slice, slice[0])
		}
		usedTables = append(usedTables, slice[1])
		if index > 0 {
			on, joinKV = joinOn(slice[1], usedTables, on)
			join = append(join, "left join "+table+" on "+joinKV)
		}
	}
	if len(orderBy) > 0 {
		StrOrderBy = " order by " + strings.Join(orderBy, ",")
	}
	if len(condition) > 0 {
		StrCondition = " where " + condition
	}
	if len(groupBy) > 0 {
		StrGroupBy = " group by " + strings.Join(groupBy, ",")
	}

	sql = fmt.Sprintf("select %v from %v %v %v %v limit %v offset %v", strings.Join(FieldSlice, ","), strings.Join(join, " "), StrCondition, StrGroupBy, StrOrderBy, listRows, (p-1)*listRows)
	return
}

func joinOn(table string, usedTables []string, on []map[string]string) (newon []map[string]string, ret string) {
	table = table + "."
	lenon := len(on)
	for index, v := range on {
		for key, val := range v {
			if strings.HasPrefix(key, table) || strings.HasPrefix(val, table) {
				for _, used := range usedTables {
					if strings.HasPrefix(key, used) || strings.HasPrefix(val, used) {
						ret = key + "=" + val
						if index > 0 {
							newon = append(newon, on[0:index]...)
						}
						if index+1 <= lenon {
							newon = append(newon, on[(index+1):]...)
						}
						return
					}
				}
			}
		}
	}
	return
}


func JoinSqlBuild(tables []string, on []map[string]string, fields map[string][]string, p, listRows int, orderBy []string, groupBy []string, condition string) (sql string, err error) {
	if len(tables) < 2 || len(tables)-1 != len(on) {
		fmt.Println(len(tables))
		fmt.Println(len(on))
		err = errors.New("The parameter is not standardized: the number of data tables of the joint query must be 2 or more, and the number of tables is one more than the on condition.")
		return
	}
	var (
		FieldSlice   []string
		StrOrderBy   string
		StrGroupBy   string
		StrCondition string
		joinKV       string
		join         = []string{tables[0]}
		usedTables   = []string{}
	)
	for table, field := range fields {
		for _, f := range field {
			FieldSlice = append(FieldSlice, strings.Trim(fmt.Sprintf("%v.%v", table, f), "."))
		}
	}
	for index, table := range tables {
		slice := strings.Split(strings.TrimSpace(table), " ")
		if len(slice) == 1 {
			slice = append(slice, slice[0])
		}
		usedTables = append(usedTables, slice[1])
		if index > 0 {
			on, joinKV = joinOn(slice[1], usedTables, on)
			join = append(join, "join "+table+" on "+joinKV)
		}
	}
	if len(orderBy) > 0 {
		StrOrderBy = " order by " + strings.Join(orderBy, ",")
	}
	if len(condition) > 0 {
		StrCondition = " where " + condition
	}
	if len(groupBy) > 0 {
		StrGroupBy = " group by " + strings.Join(groupBy, ",")
	}

	sql = fmt.Sprintf("select %v from %v %v %v %v limit %v offset %v", strings.Join(FieldSlice, ","), strings.Join(join, " "), StrCondition, StrGroupBy, StrOrderBy, listRows, (p-1)*listRows)
	return
}
