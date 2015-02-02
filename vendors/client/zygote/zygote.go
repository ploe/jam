package zygote

import (
	"jam/bazaar"
)

func init() {
	var z zygote
	bazaar.Push(&z)
}

type zygote struct {
	bazaar.Vendor
}

// - Creates the db table 'client_zygote'
// - Adds indexes to its fields name and email
func (v *zygote) Create() error {
	_, e := bazaar.Exec(`
		CREATE TABLE IF NOT EXISTS client_zygote (
			id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
			name VARCHAR(42) NOT NULL,
			email VARCHAR(254) NOT NULL,
			created DATETIME,
			updated TIMESTAMP
		)
	`)
	if e == nil { bazaar.Exec(`ALTER TABLE client_zygote ADD INDEX (name)`) }
	if e == nil { bazaar.Exec(`ALTER TABLE client_zygote ADD INDEX (email)`) }
	return e
}

func (v *zygote) Destroy() error {
	_, e:= bazaar.Exec(`DROP TABLE client_zygote`)
	return e
}


func (v *zygote) Name() string {
	return "vendors/client/zygote"
}

func (v *zygote) Desc() string {
	return `Common client data.`
}
