package gearbox

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleTST() {
	tst := &TST{}
	tst.Set("user", 1)
	fmt.Println(tst.Get("user").(int))
	fmt.Println(tst.Get("us"))
	fmt.Println(tst.Get("user1"))
	fmt.Println(tst.Get("not-existing"))

	tst.Set("account", 5)
	tst.Set("account", 6)
	fmt.Println(tst.Get("account").(int))

	tst.Set("acc@unt", 12)
	fmt.Println(tst.Get("acc@unt").(int))

	tst.Set("حساب", 15)
	fmt.Println(tst.Get("حساب").(int))
	tst.Set("", 14)
	fmt.Println(tst.Get(""))
	fmt.Println(tst.Count())
	// Output:
	// 1
	// <nil>
	// <nil>
	// <nil>
	// 6
	// 12
	// 15
	// <nil>
	// 5
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes() string {
	b := make([]byte, rand.Intn(100))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func BenchmarkTSTLookup(b *testing.B) {
	trie := &TST{}
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < rand.Intn(2000); n++ {
		trie.Set(RandStringBytes(), rand.Intn(10000))
	}
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		trie.Get("user")
	}
}