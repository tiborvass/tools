// +build ignore

package K1

func f(x int, y ...interface{})                  {}
func g(a []interface{}, b int, c ...interface{}) {}

func example() {
	temp := 5
	args := []interface{}{3, false, "abc"}
	params := args

	f(temp)
	f(temp, 42)
	f(temp, 1, temp+temp, 2)
	f(temp, func() { f(7) })
	g(nil, 0, args)
	g(nil, 0, args...)
	f(temp, params)
	f(temp, params...)
}
