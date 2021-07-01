package lib

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Params struct{}

func (q *Params) Userresponse(x int, challenge string) string {
	maxPos := len(challenge)
	r2 := challenge[maxPos-2 : maxPos]
	var n2 []int

	for i := 0; i < len(r2); i++ {
		o := int(r2[i])
		if o > 57 {
			n2 = append(n2, o-87)
		} else {
			n2 = append(n2, o-48)
		}
	}
	a := x + n2[0]*36 + n2[1]
	t := challenge[0 : maxPos-2]
	var s [][]byte
	s = append(s, []byte{}, []byte{}, []byte{}, []byte{}, []byte{})

	var m = 0
	i := 0
	u := map[int]bool{}
	for f := len(t); i < f; i++ {
		c := int(t[i])
		if !u[c] {
			u[c] = true
			s[m] = append(s[m], byte(c))
			m++
			if m == 5 {
				m = 0
			}
		}

	}
	l := a
	var v = 4
	var d = ""
	p := []int{1, 2, 5, 10, 50}

	for l > 0 {
		if l-p[v] >= 0 {
			h := rand.Intn(len(s[v]))
			d += string(s[v][h])
			l -= p[v]
		} else {
			s = s[:v]
			p = p[:v]
			v -= 1
		}
	}
	return d
}

func _e(trackArrays [][]int) [][]int {
	var t [][]int
	r2 := 0

	var (
		n2 = 0
		i  = 0
		o  = 0
	)
	for a := 0; a < len(trackArrays)-1; a++ {
		n2 = trackArrays[a+1][0] - trackArrays[a][0]
		i = trackArrays[a+1][1] - trackArrays[a][1]
		o = trackArrays[a+1][2] - trackArrays[a][2]
		if n2 == 0 && i == 0 && o == 0 {
			continue
		}
		if n2 == 0 && i == 0 {
			r2 += o
		} else {
			t = append(t, []int{n2, i, o + r2})
			r2 = 0
		}
	}

	if r2 != 0 {
		t = append(t, []int{n2, i, r2})
	}
	return t
}
func _n(trackArr []int) int {
	t := [][]int{{1, 0}, {2, 0}, {1, -1}, {1, 1}, {0, 1}, {0, -1}, {3, 0}, {2, -1}, {2, 1}}
	s := "stuvwxyz~"
	for m := 0; m < len(t); m++ {
		if trackArr[0] == t[m][0] && trackArr[1] == t[m][1] {
			return int(s[m])
		}
	}
	return 0
}
func _r(x int) string {
	t := "()*,-./0123456789:?@ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqr"
	length := len(t)
	m := ""
	i := int(math.Abs(float64(x)))
	o := i / length
	if o >= length {
		o = length - 1
	}
	if o != 0 {
		m = string(t[o])
	}
	i %= length
	a := ""
	if x < 0 {
		a += "!"
	}
	if m != "" {
		a += "$"
	}
	return a + m + string(t[i])
}
func (q *Params) EncTrack(trackArrays [][]int) string {
	var (
		i []string
		o []string
		a []string
	)

	//trackArrays = _e(trackArrays)
	for z := 0; z < len(trackArrays); z++ {
		arr := trackArrays[z]
		var t = _n(arr)
		if t == 0 {
			i = append(i, _r(arr[0]))
			o = append(o, _r(arr[1]))
		} else {
			o = append(o, string(t))
		}
		a = append(a, _r(arr[2]))
	}

	return strings.Join(i, "") + "!!" + strings.Join(o, "") + "!!" + strings.Join(a, "")
}

// Aa Params: EncTrack(trackArrays), c, s
func (q *Params) Aa(encTracks string, c []int, s string) string {
	n2 := 0
	i := 2
	a := encTracks
	l2 := len(encTracks)
	y := c[0]
	u := c[2]
	t := c[4]
	for true {
		var o string
		if n2+i <= len(s) {
			o = s[n2 : n2+i]
		} else {
			o = ""
		}

		if o == "" {
			break
		}
		n2 += i
		d, _ := strconv.ParseInt(o, 16, 64)
		f := string(d)
		g := int(d)
		l := (y*g*g + u*g + t) % l2
		a = a[:l] + f + a[l:]
	}

	return a
}
