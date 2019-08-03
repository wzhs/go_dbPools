package go_dbPools

import (
	"sync"
	"github.com/wzhs/go_dbPools/mysql"
)

var DbMan *DbManage
var oncemysql sync.Once

type DbManage struct {
	Db *mysql.SQLConnPool
}
func GetDBInstance() *DbManage {
	oncemysql.Do(func() {
		DbMan = NewDbManage()
	})
	return DbMan
}
var _mysqlhost string
var _databse string
var _username string
var _password string
var _charset string
var _mysqlmaxOpens int
var _mysqlmaxIdels int

func SetMysqlParas(host, database, user, password, charset string, maxOpenConns, maxIdleConns int){
	_mysqlhost=host
	_databse=database
	_username=user
	_password=password
	_charset=charset


	_mysqlmaxOpens=maxOpenConns
	_mysqlmaxIdels=maxIdleConns

}
//構造函數
func NewDbManage() *DbManage {
	db := &DbManage{
		//Db: InitMySQLPool("222.186.136.77:6710", "ksdb", "root", "SHIyuxing5@!", "utf8mb4", 0, 0),
		Db:mysql.InitMySQLPool(_mysqlhost,_databse,_username,_password,_charset,_mysqlmaxOpens,_mysqlmaxIdels),
	}//默认连接
	if len(_host)==0{
		db.Db=mysql.InitMySQLPool("222.186.136.77:6710", "ksdb", "root", "SHIyuxing5@!", "utf8mb4", 0, 0)
	}
	return db
}