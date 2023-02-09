package main

/*
This testing program can be converted into the actual project database
program requires gorm and sqlite3
opens/creates the "test.db" and works with the userProfile struct
this example inputs two users into the database, I dont believe that theyre any duplicates
	these entries are saved into the database
then searches for the two entries and outputs the id, this can be changed to name or any other variable
*/

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// user profile definition
type UserProfile struct {
	gorm.Model
	Name       string
	Password   string
	AdminLevel uint8
	Allergies  string
}

// sets up Sqlite3 database
func connnectDB(dbName string) *gorm.DB {
	dbName += ".db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("conncected to test.db")
	}
	return db
}

// function wraps all of the auto migration calls: for future use
func buildTables(db *gorm.DB) {
	db.AutoMigrate(&UserProfile{})
}

// adding users function: future use when more tables added
func addUser(addUser *UserProfile, db *gorm.DB) {
	db.Create(&addUser)
}

func testUserAdd(db *gorm.DB) {
	//adding users test code, works in command line
	numberOfEntries := 0
	fmt.Print("Enter Number of Entries: ")
	fmt.Scan(&numberOfEntries)

	insertedUsers := make([]UserProfile, numberOfEntries)
	for i := 0; i < numberOfEntries; i++ {
		fmt.Println("Enter User Data In the Following Format: UserName Password AdminLevel")
		fmt.Scan(&insertedUsers[i].Name, &insertedUsers[i].Password, &insertedUsers[i].AdminLevel)
		fmt.Println("Enter Any Allergies Seperated By a comma")
		fmt.Scan(&insertedUsers[i].Allergies)
		addUser(&insertedUsers[i], db)
	}
}

func testUserSearch(db *gorm.DB) {
	//Function is a framework for a future function to query database

	numberOfEntries := 0
	fmt.Print("Enter Number of Queries: ")
	fmt.Scan(&numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		var searchName string
		var foundUser UserProfile
		fmt.Print("Enter UserName To Search: ")
		fmt.Scan(&searchName)
		db.Where("Name = ?", searchName).First(&foundUser)

		fmt.Println("\nStored Information:\nUserName: " + foundUser.Name + "\nPassword: " + foundUser.Password)
		fmt.Println("Allergies: " + foundUser.Allergies + "\nCreatedAt: " + foundUser.CreatedAt.String())
		fmt.Println()
	}
}

func testSoftDelete(db *gorm.DB) {
	//framework for future functions used to softDelete
	softDeletes := 0
	fmt.Print("Enter Number of softDeletes: ")
	fmt.Scan(&softDeletes)
	for i := 0; i < softDeletes; i++ {
		var deleteUser UserProfile
		var dUserName string
		fmt.Print("Enter userName To Be Soft Deleted: ")
		fmt.Scan(&dUserName)
		db.Where("Name = ?", dUserName).First(&deleteUser)
		db.Delete(&deleteUser)
		deleteUser = UserProfile{}
		db.Unscoped().Where("Name = ?", dUserName).First(&deleteUser)
		fmt.Println("\n" + dUserName + " was soft deleted at: " + deleteUser.DeletedAt.Time.String())
	}
}

func testReturnSoftDelete(db *gorm.DB) {
	//Demonstration/Framework function used to show how the database works with returning softdeleted data

	returns := 0
	fmt.Print("Enter Number of Returns: ")
	fmt.Scan(&returns)
	for i := 0; i < returns; i++ {
		var returnedUser UserProfile
		var rUserName string

		fmt.Print("Enter userName To Be Returned: ")
		fmt.Scan(&rUserName)

		db.Unscoped().Where("Name = ?", rUserName).First(&returnedUser).Update("deleted_at", nil)

		returnedUser = UserProfile{}
		db.Where("Name = ?", rUserName).First(&returnedUser)

		fmt.Println("\n" + rUserName + " Has been returned to the database, printing stored data: ")
		fmt.Println("UserName: " + returnedUser.Name + "\nPassword: " + returnedUser.Password)
		fmt.Println("Allergies: " + returnedUser.Allergies + "\nCreatedAt: " + returnedUser.CreatedAt.String())
	}
}

func testUpdate(db *gorm.DB) {
	//function allows user to update specific fields from the commandline used to show how interaction with
	//database works/is a framework for future functions

	updates := 0
	fmt.Print("Enter Number of updates: ")
	fmt.Scan(&updates)
	for i := 0; i < updates; i++ {

		var uUserName string
		var uFieldName string
		var uValue string
		var updateUser UserProfile

		fmt.Print("\nEnter userName To Be Updated: ")
		fmt.Scan(&uUserName)
		fmt.Print("\nEnter Field To Be Updated: ")
		fmt.Scan(&uFieldName)
		fmt.Print("\nEnter Updated Value: ")
		fmt.Scan(&uValue)

		db.Where("Name = ?", uUserName).First(&updateUser).Update(uFieldName, uValue)
		fmt.Println("\nStored Information:\nUserName: " + updateUser.Name + "\nPassword: " + updateUser.Password)
		fmt.Println("Allergies: " + updateUser.Allergies + "\nCreatedAt: " + updateUser.CreatedAt.String())
		fmt.Println()
	}
}

func testHardDelete(db *gorm.DB) {
	//This function hard deletes values in the database
	//this can be used later as a framework to build future delete functions

	hardDeletes := 0
	fmt.Print("Enter Number of HardDeletes: ")
	fmt.Scan(&hardDeletes)

	for i := 0; i < hardDeletes; i++ {
		var hdUserName string
		var hardDeleteUser UserProfile
		fmt.Print("\nEnter userName of HardDelete: ")
		fmt.Scan(&hdUserName)

		db.Where("Name = ?", hdUserName).First(&hardDeleteUser)

		fmt.Println("\nUser Stored Information:\nUserName: " + hardDeleteUser.Name + "\nPassword: " + hardDeleteUser.Password)
		fmt.Println("Allergies: " + hardDeleteUser.Allergies + "\nCreatedAt: " + hardDeleteUser.CreatedAt.String() + "\n\nDeleting User")

		db.Unscoped().Delete(&hardDeleteUser)

		hardDeleteUser = UserProfile{}
		db.Where("Name = ?", hdUserName).First(&hardDeleteUser)

		if hardDeleteUser.ID == 0 {
			fmt.Println("User Was Successfully Deleted")
		} else {
			fmt.Println("User Was Unsuccessfully Deleted")
		}
	}
}

type Account struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func createAcc(ctx *gin.Context) {
	// user validation
	var req Account
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accCreationObject := Account{
		Name:     req.Name,
		Password: req.Password,
	}

	// need to update the add use function to accept diff parameters
	account, err := db.addUser(ctx, accCreationObject)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func main() {

	db := connnectDB("test")

	router := gin.Default()
	router.POST("/accounts", db.createAccount)
	router.Run("localhost:8080")

	// testing

	buildTables(db)

	testUserAdd(db)
	testUserSearch(db)
	testSoftDelete(db)
	testReturnSoftDelete(db)
	testUpdate(db)
	testHardDelete(db)
}
