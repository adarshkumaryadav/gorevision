package main

import (
	"fmt"
	"math/rand"
)

// defining interface
type db interface {
	close()
	createUser(userNam string) (userId int)
	deleteUser(userId int)
}

// wanted to have an app that create and delete user from db
// so here db can change but the app logic should not change.

// let's start my application

// func (d db) StartApplication() {
// below is possible
/*func StartApplication(d db) {
	defer d.close()
	id := d.createUser("Adarsh")
	d.deleteUser(id)
}
*/
// ideally in prod we apply the same logic by introducing one struct for the app
type application struct {
	db db
}

func initializeApp(db db) *application {
	return &application{db: db}
}
func (a application) Run() {
	defer a.db.close()
	id := a.db.createUser("AADDDA")
	a.db.deleteUser(id)
}

// now logic for a type to be of interface type
type mySql struct{}

//let's implement all method

func (m mySql) close() {
	fmt.Println("Closing the mySql DB")
}
func (m mySql) createUser(name string) int {
	fmt.Println("Created user ", name)
	return rand.Int()
}
func (m mySql) deleteUser(id int) {
	fmt.Println("Deleting the entry for id", id)
}

// we can have another type of db let's say postsql
type postSql struct{}

//let's implement all method

func (m postSql) close() {
	fmt.Println("Closing the postSql DB")
}
func (m postSql) createUser(name string) int {
	fmt.Println("Created user ", name)
	return rand.Int()

}
func (m postSql) deleteUser(id int) {
	fmt.Println("Deleting the entry for id", id)
}

// now start with the main logic in main func

func main() {
	// new db instance
	// we can swap the db type here easily...
	newDb := postSql{}
	app := initializeApp(newDb)
	app.Run()
}
