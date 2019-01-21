package dbops

import "testing"

//init(dblogin, truncate tables) -> run tests -> clear data(truncate tables) 保证tests和其他tests存在循环依赖 或者代码干扰

func clearTables() {
	dbConn.Exec("truncate userstable")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate  comments")
	dbConn.Exec("truncate  sessions")
}

func TestMain(m *testing.M){
	clearTables()//初始化
	m.Run() //跑所有的test
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {

	err := AddUserCredential("avenssi", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T)  {
	pwd, err := GetUserCredential("avenssi")
	if pwd!="123" || err != nil {
		t.Errorf( "Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("avenssi", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T)  {
	pwd, err := GetUserCredential("avenssi")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}