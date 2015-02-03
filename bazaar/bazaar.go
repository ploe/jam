// bazaar is the package that kind of presides over the MySQL connection 
// and it also allows us to think about the Vendors as a set.
//
// It implements a subtype Vendor which all Vendors should inherit from.
//
// It also sets up the connection from the command line parameters and a
// password that it asks for from stdin (so it isn't echoed on screen)
//
//The flags are:
//	'mysql-db'
//	'mysql-host'
//	'mysql-user'
package bazaar

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"github.com/howeyc/gopass"
	_"github.com/go-sql-driver/mysql"
)

type vendor interface {
	Create() error
	CreateBreak(error) bool
	Destroy() error
//	DestroyBreak() bool

	Name() string
	Desc() string
}

var banked map[string]vendor
func init() {
	banked = make(map[string]vendor)
}

// Push banks a vendor in the bazaar. It returns an error if it 
// fails to find a name for the vendor.
func Push(v vendor) error {
	var e error = nil
	if v.Name() != "" {
		banked[v.Name()] = v
	} else {
		e = errors.New("bazaar: No name returned by vendor")
	}
	return e
}

// Create calls all the Create methods on the banked vendors
//
// The error we get back is the one the vendor's CreateBreak function
// returned true for. Otherwise we get nil.
func Create() error {
	var e error = nil
	for _, v := range banked {
		e = v.Create()
		if v.CreateBreak(e) { break; }
	}
	return e
}

// Destroy calls the Destroy method on each banked vendor.
// It also drops the database specified in the flag 'mysql-db'
func Destroy() error {
	var e error = nil
	for _, v := range banked {
		e = v.Destroy()
		if e != nil { break; }
	}
		
	_, e = con.Exec("DROP DATABASE " + db)
	return e
}

const (
	OK = iota
	EOMYSQL_PING
	EOMYSQL_CREATE
	EOMYSQL_USE
)

var db, user, host, dsn string
var con *sql.DB;
func init() {
	flag.StringVar(&db, "mysql-db", "", "The 'MySQL Database' for the current sandbox")
	flag.StringVar(&user, "mysql-user", "", "The mysql-user to build the database instance with")
	flag.StringVar(&host, "mysql-host", "localhost:3306", "The mysql-host is the name of the server where MySQL lives. It can either be an IP or hostname followed by a port. (e.g. 127.0.01:3306)")
}

// Connects bazaar to the MySQL instance specified in the flags.
// The reason why this is separated from the init func is so we can 
// specify flags in other packages.
//
// This func requests a password from stdin and creates the db specified
// in the flag 'mysql-db'
func Connect() {
	fmt.Fprint(os.Stderr, "MySQL Password: ")	
	dsn = user + ":" + string(gopass.GetPasswd()) + "@tcp(" + host + ")/"

	var e error = nil
	con, e = OpenCox()
	if e != nil { 
		fmt.Fprintln(os.Stderr, "bazaar.Connect => MySQL " + e.Error())
		os.Exit(EOMYSQL_PING)
	}

	_, e = con.Exec("CREATE DATABASE IF NOT EXISTS " + db)
	if e != nil {
		fmt.Println("bazaar.Connect => MySQL " + e.Error())
		os.Exit(EOMYSQL_CREATE)
	}

	_, e = con.Exec("USE " + db)
	if e != nil {
		fmt.Println("bazaar.Connect => MySQL " + e.Error())
		os.Exit(EOMYSQL_USE)
	}
}

// OpenCox returns us a DB connection, using the details we have stored in
// the bazaar. It also returns an error if the Ping we do fails.
func OpenCox() (*sql.DB, error) {
	con, _ := sql.Open("mysql", dsn)
	e := con.Ping()
	return con, e
}

// A wrapper for the bazaar's Exec method
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return con.Exec(query, args...)
}

// A wrapper for the bazaar's Query method
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return con.Query(query, args...)
}

// Read-only value db, the command line param mysql-db
func DB() string {
	return db
}


// Read-only value user, the command line param mysql-user
func User() string {
	return user
}

// Read-only value host, the command line param mysql-host
func Host() string {
	return host
}

// A vendor is a struct representing a lib in the bazaar package.
//
// The Vendor type is to be used as the subtype for vendor packages.
//
// Any vendors you want to bank in the system must implement the same
// public methods as this type.
type Vendor struct {};

// Create is the constructor for the Vendor's database. By calling
// create it should be safe to assume that the required tables/resources
// will be allocated.
//
// The Vendor subtype will always return an error for this
func (v *Vendor) Create() error {
	return errors.New("bazaar: 'Create' method undefined")
}

// Destroy is the destructor for the Vendor's database. Calling this 
// should get rid of the tables and resources allocated by the vendor's
// Create method.
//
// The Vendor subtype will always return an error for this.
func (v *Vendor) Destroy() error {
	return errors.New("bazaar: 'Destroy' method undefined")
}

// Returns the Name of the vendor, duh!
//
// The Vendor subtype will always return a blank string.
func (v *Vendor) Name() string {
	return ""
}

// Returns a description, kind of like a dumb blurb describing the Vendor.
//
// The Vendor subtype will always return a blank string.
func (v *Vendor) Desc() string {
	return ""
}

// Called to see if bazaar.Create encountered an error from the vendor.
// If it returns true bazaar.Create bails. If false bazaar.Create keeps
// rolling.
//
// The Vendor subtype always returns true on any error. It will give up
// iterating over the banked vendors if it finds one.
func (v *Vendor) CreateBreak(e error) bool {
	return (e != nil)
}
