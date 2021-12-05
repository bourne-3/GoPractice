package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var width = flag.Int("w", 256, "hash width (256 or 512)")  // name, value, usage

func h1() {
	flag.Parse()
	var function func(b []byte) []byte
	switch *width {
	case 256:
		// 给函数值 赋值
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width specified.")
	}
	//b, err := ioutil.ReadAll(os.Stdin)
	//if err != nil {
	//	log.Fatal(err)
	//}
	r := bufio.NewReader(os.Stdin)
	line, _, _ := r.ReadLine()

	fmt.Printf("%x\n", function([]byte(line)))  // 调用上面的函数
}

func h2() {
	res := ShaBitDiff([]byte("x"), []byte("y"))
	fmt.Printf(strconv.Itoa(res))
}


func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		switch {
		case i >= len(a):
			count += popCount(b[i])
		case i >= len(b):
			count += popCount(a[i])
		default:
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

// ShaBitDiff returns the number of bits that are different in the SHA256
// hashes of two buffers.
func ShaBitDiff(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)
	return bitDiff(shaA[:], shaB[:])
}
