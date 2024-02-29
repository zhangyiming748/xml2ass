package main

import (
	"github.com/zhangyiming748/xml2ass/constant"
	"github.com/zhangyiming748/xml2ass/conv"
)

func main() {
	constant.SetLogLevel("Debug")
	conv.GetXmls()
}
