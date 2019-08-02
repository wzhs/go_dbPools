package test

import (
	"fmt"
	"testing"
	"go_dbPools"
	"go_dbPools/redis"
)

func Test_InitPool(t *testing.T)  {

	go_dbPools.SetRedisParas("222.186.136.51:7910","shiyuxing521",0,-1,-1)
	_,err:=go_dbPools.GetRedisInstance().Pool.SetString("test2","234")
	if err!=nil{
		fmt.Println(err.Error())
	}
}

func Test_RedisPool(t *testing.T){
	pool:= redis.InitRedisPool("222.186.136.51:7910","shiyuxing521",0,-1,-1)

	_,err:=pool.SetString("test","123")
	if err!=nil{
		fmt.Println(err.Error())
	}
	result,err:=pool.GetString("test2")

	if err!=nil{
		fmt.Println("ERROR",err.Error())
	}else {
		fmt.Println("SUCCESS",result)
	}
	//Set test
	//for i := 1; i < 100; i++ {
	//	_,err:=pool.Do("SADD","IPS",fmt.Sprintf("192.168.0.%d",i))
	//	if err!=nil{
	//		fmt.Println(err.Error())
	//		break
	//	}
	//}
	//查询是否存在在集合中
	result2,err:=pool.Do("SISMEMBER","IPS","192.168.0.115")
	if err!=nil{
		fmt.Println("ERROR",err.Error())
	}else {
		fmt.Println("SUCCESS",result2.(int64))
	}
}
