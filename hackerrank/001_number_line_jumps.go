package hackerrank

// https://www.hackerrank.com/challenges/kangaroo/problem

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 <= v2 {
		return "NO"
	}
	if v1 > x2-x1+v2 {
		return "NO"
	}
	dx := x2*v1 - x1*v2
	dv := v1 - v2
	if dx%dv > 0 {
		return "NO"
	}
	k := dx / dv
	if (k-x1)%v1 > 0 {
		return "NO"
	}
	return "YES"
}
