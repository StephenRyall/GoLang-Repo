package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := strconv.FormatInt(int64(123), 10)
	r := strings.Split(s, "")

	k := strconv.FormatInt(int64(456), 10)
	p := strings.Split(k, "")

	c := strconv.FormatInt(int64(789), 10)
	d := strings.Split(c, "")

	e := append(r, p...)
	o := append(e, d...)
	fmt.Println(o)
}
