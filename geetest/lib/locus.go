package lib

import (
	"math"
	"math/rand"
	"time"
)

const charMaps string = "abcdefghijklmnopqrstuvwxyz0123456789"

type Random struct{}

func (r *Random) Range(a int, b int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(b-a) + a
}
func randPn(a int, b int) int {
	rand.Seed(time.Now().UnixNano())
	rk := rand.Intn(b)
	if rand.Intn(2) == 0 {
		rk = -rk
	}
	return rk + a
}
func randTrack(x int) [][]int {
	rands := Random{}
	_0xC4F9 := [][]int{
		{rands.Range(-200, 150), 29, 0},
		{-1, 0, rands.Range(500, 560)},
		{0, 1, 23},
		{-1, 0, 250},
		{1, 0, 21},
		{1, 0, 15},
		{2, 0, 24},
	}

	var a float64
	var move float64
	current := 0.0
	midX := float64(x) * 2.3 / 5
	t := 0.07
	v := 0.0
	for current < float64(x)-2 {
		if current < midX {
			a = 2
		} else {
			a = -4.0
		}
		v0 := v
		t = float64(rands.Range(3, 5)) / 10
		v = v0 + a*t
		move = v0*t + 1/2*a*t*t
		current += move
		_0xC4F9 = append(_0xC4F9, []int{int(math.Round(move)), 0, int(t * 30)})
	}
	_0xC4F9 = append(_0xC4F9, [][]int{
		{1, 0, 9},
		{2, 0, 80},
		{1, 0, 8},
		{1, 0, 16},
		{2, 0, 81},
		{1, 0, 16},
		{1, 0, 16},
		{0, 0, rands.Range(500, 900)},
	}...)
	return _0xC4F9
}
func (r *Random) AesKey() []byte {
	rand.Seed(time.Now().UnixNano())
	blk := make([]byte, 16)
	_, err := rand.Read(blk)
	if err != nil {
		panic(err)
	}
	return blk
}
func (r *Random) MoveSlide(xPos int, offset int) ([][]int, int, int) {
	var (
		TotalX  = 0
		TmsUnix = 0
	)

	var retTrack [][]int
	tempTrack := randTrack(xPos)
	tempX := 0
	retTrack = append(retTrack, tempTrack[0])
	for i := 1; i < len(tempTrack)-offset; i++ {
		if tempX < xPos-offset {
			track := tempTrack[i]
			track[2] = randPn(track[2], 2)
			if i == 0 {
				track[2] = 0
			}
			tempX += track[0]
			TmsUnix += track[2]
			retTrack = append(retTrack, track)
		}
	}
	TotalX = tempX
	for i := len(tempTrack) - offset; i < len(tempTrack); i++ {
		track := tempTrack[i]
		track[2] = randPn(track[2], 2)
		TotalX += track[0]
		TmsUnix += track[2]
		retTrack = append(retTrack, track)
	}
	return retTrack, TotalX, TmsUnix
}
