package main
import "fmt"
// DEFER
// The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
// Deferred functions are typically used to close database connections, file handlers and the like
const ( 
	logDeleted = "user deleted"
    logNotFound = "user not found"
	logAdmin = "admin deleted"
)
func logAndDelete(users map[string]userD, name string) (log string) {
	user, ok := users[name]
	defer delete(users, name)
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}
type userD struct {
	name string
	number int
	admin bool
}
func testDefer(users map[string]userD, name string){
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}
func DeferKeyword(){
users := map[string]userD{
	"john": {
		name: "john",
		number: 18965554631,
		admin: true,
	},
	"elon": {
		name: "elon",
		number: 198755556452,
		admin: true,
	},
	"paul": {
		name: "paul",
		number: 84848493930028,
		admin: false,
	},
}
fmt.Println(users)
testDefer(users, "john")
testDefer(users, "elon")
testDefer(users, "paul")
fmt.Println(users)
}