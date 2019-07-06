package models

import "database/sql"

type Groups struct {
	ID              uint16 `json:"id"`
	Name            string `json:"name"`
	Provider        string `json:"provider"`
	ProviderID      string `json:"provider_id"`
	MemberNumber    string `json:"member_number"`
	Cover           string `json:"cover"`
	Identity        string `json:"identity"`
	DetaDescription string `json:"meta_description"`
}

func NewGroups() *Groups {
	return &Groups{}
}

func GetTableGroups() string {
	return getTable("groups")
}

func GetSQLList() string {
	return "SELECT id, name, provider, provider_id, member_number, cover, identity, meta_description FROM groups WHERE status = ?"
}

func GetSQLOne() string {
	return "SELECT id, name, provider, provider_id, member_number, cover, identity, meta_description FROM groups WHERE identity = ?"
}

func (this *Groups) List(listRows int, status ...int) (groups []Groups, rows int64, err error) {
	var selectStmt *sql.Stmt
	DB := getConnection()
	statusType := 1
	rows = 0

	if len(status) > 0 {
		statusType = status[0]
	}

	selectStmt, err = DB.Prepare(GetSQLList())
	results, err := selectStmt.Query(statusType)

	for results.Next() {
		group := new(Groups)
		err = results.Scan(&group.ID, &group.Name, &group.Provider, &group.ProviderID, &group.MemberNumber, &group.Cover, &group.Identity, &group.DetaDescription)
		groups = append(groups, *group)
		rows++
	}

	return groups, rows, err
}

func (this *Groups) One(identity string) (group Groups, err error) {
	var selectStmt *sql.Stmt
	DB := getConnection()
	selectStmt, err = DB.Prepare(GetSQLOne())
	err = selectStmt.QueryRow(identity).Scan(&group.ID, &group.Name, &group.Provider, &group.ProviderID, &group.MemberNumber, &group.Cover, &group.Identity, &group.DetaDescription)

	return group, err
}
