package abck

import (
	"fmt"
	"time"
	"strings"
	"math"
	"regexp"
	"errors"
	"math/rand"
	"sort"
	"strconv"
	"reflect"
)

func randomNumberInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func randomNumberFloat() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Float64()
}

func GetCfDate() int64 {
	t := time.Now()
	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	return tUnixMilli
}

func D3() int64 {
	return GetCfDate() % 1e7
}

func rir(t int, a int, e int, n int) int {
	if t > a && t <= e {
		t += n % (e - a)
		if t > e {
			t = t - e + a
		}
	}
	return t
}

func Od(t string, a string) string {
	var e []string
	n := len(a);
	if (n > 0) {
			for o := 0; o < len(t); o++ {
					m := []rune(t)[o]
					r := string(t[o])
					i := a[(o % n)];
					m = rune(rir(int(m), 47, 57, int(i)))
					if m != []rune(t)[o] {
						r = string(m)
					}
					e = append(e, r)
			}
			if len(e) > 0 {
				return strings.Join(e[:], "")
			}
	}
	return t
}

func Ab(t string) int {
	a := 0
	for e := 0; e < len(t); e++ {
			n := []rune(t)[e]
			if n < 128 {
				a += int(n)
			}
	}
	return a
}

func Uar(userAgent string) string {
	return regexp.MustCompile(`/\\|"/g`).ReplaceAllString(userAgent, "")
}

func Fas(browser string) (int, error) {
	if browser == "chrome" {
		return 30261693, nil
	} else if browser == "firefox" {
		return 26067385, nil
	} else {
		return 0, errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
}

func mrjp(flt float64) float64 {
	return flt
}

func mrinf(flt float64) bool {
	return math.IsInf(flt, 0)
}

func invoke(fn interface{}, args ...float64) {
	v := reflect.ValueOf(fn)
	rargs := make([]reflect.Value, len(args))
	for i, a := range args {
			rargs[i] = reflect.ValueOf(a)
	}
	v.Call(rargs)
}

func Getmr() string {
	t := ""
	var e []interface{}
  e = append(e, math.Abs, math.Acos, math.Asin, math.Atanh, math.Cbrt, math.Exp, mrjp, math.Sqrt, mrinf, math.IsNaN, mrjp, mrjp, mrjp, mrjp)
	for n := 0; n < len(e); n++ {
		var o []float64
		var m float64
		m = 0.0
		r := time.Now()
		c := 0
		for i := 0; i < 1000 && m < .6; i++ {
				b := time.Now()
				for d := 0; d < 4e3; d++ {
					invoke(e[n], 3.14)
				}
				k := time.Now()
				o = append(o, 1000 * math.Round(float64(k.UnixNano() / 1000000) - float64(b.UnixNano() / 1000000)))
				m = float64(k.UnixNano() / 1000000) - float64(r.UnixNano() / 1000000)
				s := o
			sort.Float64s(s);
			c = int(s[int(math.Floor(float64(len(s) / 2)))]) / 5
			t = fmt.Sprintf("%v%v,", t, c)
		}
	}
	return t
}

func Sed(browser string) string {
	return "0,0,0,0,1,0,0"
}

func Np(browser string) (string, error) {
	if browser == "chrome" {
		return "10321144241322243122", nil
	} else if browser == "firefox" {
		return "11133333331333333333", nil
	} else {
		return "", errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
}

func calDis(t []int) float64 {
	a := t[0] - t[1]
	e := t[2] - t[3]
	n := t[4] - t[5]
	return math.Floor(math.Sqrt(float64(a * a + e * e + n * n)))
}

func Jrs(t int) ([]float64, error) {
	a := int(math.Floor(100000 * randomNumberFloat() + 10000))
	e := strconv.Itoa(t * a)
	m := len(e) >= 18
	n := 0
	var o []int
	for ; len(o) < 6; {
		num, err := strconv.Atoi(e[n:(n + 2)])
		if err != nil {
			return []float64{}, err
		}
		o = append(o, num)
		if m {
			n = n + 3
		} else {
			n = n + 2
		}
	}
	return []float64{float64(a), calDis(o)}, nil
}

func cc(t int) func(int, int) int {
	a := t % 4;
	if 2 == a {
		a = 3
	}
	e := 42 + a
	var n func(int, int) int
	n = func(t int, a int) int {
		return 0
	}
	if (42 == e) {
		n = func(t int, a int) int {
			return t * a
		}
	} else if (43 == e) {
		n = func(t int, a int) int {
			return t + a
		}
	} else {
		n = func(t int, a int) int {
			return t - a
		}
	}
	return n
}

func O9(d3 int64) int {
	a := int(d3)
	for n := 0; n < 5; n++ {
		o := math.Mod(math.Floor(float64(d3) / math.Pow(10, float64(n))), 10)
		m := o + 1;
		op := cc(int(o))
		a = op(int(a), int(m))
	}
	return a * 3
}

func GetType(t string) int {
	t = strings.ToLower(t)
	if "text" == t || "search" == t || "url" == t || "email" == t || "tel" == t || "number" == t {
		return 0
	} else if "password" == t {
		return 1
	} else {
		return 2
	}
}

func Z1(startTimestamp int64) int64 {
	return startTimestamp / 2016 * 2016
}

func floatToString(flt float64, precision int) string {
	return strconv.FormatFloat(flt, 'f', precision, 64)
}

type ScreenData struct {
	availWidth int
	availHeight int
	width int
	height int
	innerWidth int
	innerHeight int
	outerWidth int
}

type BrowserData struct {
	lang string
	userAgent string
	screenData ScreenData
}

func Gd(browser string, browserData BrowserData, startTimestamp int64, d3 int64) (string, error) {
	userAgent := Uar(browserData.userAgent);
	d := randomNumberFloat()
	xagg := 0
	psub := 0
	plen := 0
	bd := ""
	if browser == "chrome" {
		xagg = 12147
		psub = 20030107
		plen = 3
		bd = ",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:0,sc:0,wrc:1,isc:0,vib:1,bat:1,x11:0,x12:1"
	} else if browser == "firefox" {
		xagg = 11059
		psub = 20100101
		plen = 0
		bd = fmt.Sprintf(",cpen:0,i1:0,dm:0,cwen:0,non:1,opc:0,fc:1,sc:0,wrc:1,isc:%v,vib:1,bat:0,x11:0,x12:1", randomNumberInt(60, 100))
	} else {
		return "", errors.New("Unsupported browser specified (use \"chrome\" or \"firefox\")")
	}
	return fmt.Sprintf(
		"%v,uaend,%v,%v,%v,Gecko,%v,0,0,0,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,loc:",
		userAgent,
		xagg,
		psub,
		browserData.lang,
		plen,
		Z1(startTimestamp),
		d3,
		browserData.screenData.availWidth,
		browserData.screenData.availHeight,
		browserData.screenData.width, 
		browserData.screenData.height,
		browserData.screenData.innerWidth,
		browserData.screenData.innerHeight,
		browserData.screenData.outerWidth,
		bd,
		Ab(userAgent),
		floatToString(d, 16)[0:11] + floatToString(math.Floor(1000 * d / 2), -1),
		startTimestamp / 2), nil
}