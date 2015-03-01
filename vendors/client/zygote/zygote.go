package zygote

import (
	"crypto/sha1"
	"fmt"
	"io"
	"strconv"
	"time"
	
	"jam/bazaar"
)

// Takes a few seconds to run. The point of this 'un is so that two 
// passwords that are the same don't look the same in the db.
func digest(password, salt string) string {
	h := sha1.New()
	var digest string
	for i := 1; i <= 500000; i++ {
		io.WriteString(h, password)
		io.WriteString(h, salt)
		digest = fmt.Sprintf("%x", h.Sum(nil))
		io.WriteString(h, digest)
	}
	return digest
}

func init() {
	var z VendorsClientZygote
	bazaar.Push(&z)
}

// type that represents a row of data in the client_zygote table
type Data struct {
	Name, Email, Password, Digest string
	ID, Created, Updated int64
}

// Adds a new client_zygote row to the table
// If the Data has a password field we'll automatically suss out the SHA-1
// digests for it.
func (d Data) Insert() error {
	now := time.Now().Unix()
	d.Digest = digest(d.Password, strconv.FormatInt(now, 10))
	_, e := bazaar.Exec(`
		INSERT INTO client_zygote (name, email, digest, created) 
		VALUES (?, ?, ?, FROM_UNIXTIME(?))
	`,  d.Name, d.Email, d.Digest, now)

	return e	
}

// Returns one zygote from the client_zygote table, as specified by the 
// param email. The reason we only return one is because an email should
// only be registered once.
func ByEmail(email string) *Data {
	return selectData("email", email)
}

// Returns one zygote from the client_zygote table, as specified by the
// param name. A name can not be shared across the db.
func ByName(name string) *Data {
	return selectData("name", name)
}

// Returns one zygote from the client_zygote table, as specified by the
// param id, which is the primary key.
func ByID(id uint64) *Data {
	return selectData("id", id)
}

// Grabs us a row, with the specified attr set to val
// If anything fails it returns nil
func selectData(attr string, val interface{}) *Data {
	rows, _ := bazaar.Query(`SELECT id, name, email, digest, UNIX_TIMESTAMP(created), UNIX_TIMESTAMP(updated) FROM client_zygote WHERE `+attr+`=?`, val)
	defer rows.Close()

	var d *Data = nil
	if rows.Next() {
		d = new(Data)
		e := rows.Scan(&d.ID, &d.Name, &d.Email, &d.Digest, &d.Created, &d.Updated)
		if e != nil { d = nil }
	}
	
	return d
}


// Vendor interface, registered in the bazaar
type VendorsClientZygote struct {
	bazaar.Vendor
}

// - Creates the db table 'client_zygote'
// - Adds indexes to its fields name and email
func (v *VendorsClientZygote) Create() error {
	_, e := bazaar.Exec(`
		CREATE TABLE IF NOT EXISTS client_zygote (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
			name VARCHAR(50) NOT NULL,
			email VARCHAR(50) NOT NULL,
			digest VARCHAR(40),
			created DATETIME,
			updated TIMESTAMP
		)
	`)
	if e == nil { bazaar.Exec(`ALTER TABLE client_zygote ADD INDEX (name)`) }
	if e == nil { bazaar.Exec(`ALTER TABLE client_zygote ADD INDEX (email)`) }
	return e
}

func (v *VendorsClientZygote) Destroy() error {
	_, e:= bazaar.Exec(`DROP TABLE client_zygote`)
	return e
}

func (v *VendorsClientZygote) Name() string {
	return "VendorsClientZygote"
}

func (v *VendorsClientZygote) Desc() string {
	return `Common client data.`
}
