package counter

// import (
// 	"sync"
// )

// type Counter struct {
// 	m sync.Map
// }

// func NewCounter() *Counter {
// 	return &Counter{}
// }

// func (c *Counter) Load(key int) (string, bool) {
// 	val, ok := c.m.Load(key)
// 	if ok {
// 		return val.(string), true
// 	}
// 	return "", false
// }

// func (c *Counter) Store(key int, value string) {
// 	c.m.Store(key, value)
// }

// func (c *Counter) Range(f func(k int, v string) bool) {
// 	c.m.Range(func(key, value any) bool {
// 		return f(key.(int), value.(string))
// 	})
// }

// func (c *Counter) Delete() {
// 	c.m.Clear()
// }
