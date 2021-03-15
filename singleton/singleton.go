package main

import "sync"

func main() {
	
}


/*
	懒汉模式
*/
type singleton struct {}
var instance *singleton
func GetInstance() *singleton {
	if instance != nil {
		return instance
	}
	return &singleton{}
}



/*
	整个方法加锁
*/
var lock sync.Mutex
func GetInstance1() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil{
		instance = &singleton{}
	}
	return instance
}

/*
	创建时候加锁
*/
func GetInstance2() *singleton {
	if instance == nil {//非原子操作，此时可能已经创建了
		lock.Lock()
		instance = &singleton{}
		lock.Unlock()
	}
	return instance
}


/*
	双重校验锁
*/
func GetInstance3() *singleton {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &singleton{}
		}
		lock.Unlock()
	}
	return instance
}


/*
	原子操作：sync中找到了Once类型。它能保证某个操作仅且只执行一次
*/
var once sync.Once
func GetInstance4() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}













