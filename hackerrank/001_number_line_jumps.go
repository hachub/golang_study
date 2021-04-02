package hackerrank

// https://www.hackerrank.com/challenges/kangaroo/problem

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}
	if v1 > x2-x1+v2 {
		return "NO"
	}
	return "YES"
}
