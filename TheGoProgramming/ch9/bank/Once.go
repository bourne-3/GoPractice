package bank

import "sync"

var scores map[string]int
//var mu sync.Mutex
var rw sync.RWMutex

// 懒加载
var loadScoreOnce sync.Once

func loadScores() {
	scores = map[string]int{
		"bourne" :80,
		"lucy" : 58,
		"kim" : 69,
	}
}
// not concurrency-safe
func Score(name string) int {

	// 1 解决1
	//mu.Lock()  // 没有办法对该变量进行并发访问
	//defer mu.Unlock()

	// 2 解决2
		// 体现到了读的好处，可以并发的访问读的情况了
	rw.RLock()
	if scores != nil {
		// 使用读锁直接返回
		s := scores[name]
		rw.RUnlock()
		return s
	}
	// 不可以直接获得，读锁不能再使用 更换
	rw.RUnlock()

	mu.Lock()
	if scores == nil{
		loadScores()
	}

	s := scores[name]
	mu.Unlock()
	return s
}

// 解决3
func Score2(name string) int{
	// 懒加载
	loadScoreOnce.Do(loadScores)


	return scores[name]
}


// 如果在并发的情况下可能是这样的
//func loadScores2() {
//	scores = make(map[string]int)
//	scores["bourne"] = 80
//	// ...
//}

