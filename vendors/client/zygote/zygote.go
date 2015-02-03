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
	var z zygote
	bazaar.Push(&z)
}

type zygote struct {
	bazaar.Vendor
}

type Data struct {
	Name, Email, Password, Digest string
	Created, Updated int64
}


// Adds a new client_zygote row to the table
// If the Data has a password field we'll automatically suss out the SHA-1
// digests for it.
func Insert(d Data) error {
	now := time.Now().Unix()
	digest := digest(d.Password, strconv.FormatInt(now, 10))

	 _, e := bazaar.Exec(`
		INSERT INTO client_zygote (name, email, digest, created) 
		VALUES (?, ?, ?, FROM_UNIXTIME(?))
	`,  d.Name, d.Email, digest, now)

	return e
}

// - Creates the db table 'client_zygote'
// - Adds indexes to its fields name and email
func (v *zygote) Create() error {
	_, e := bazaar.Exec(`
		CREATE TABLE IF NOT EXISTS client_zygote (
			id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, 
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
