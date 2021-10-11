package gcd

import (
	"fmt"
	"pks_sort/algo"
	"pks_sort/randUtil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	nums   [2]int
	result int
}

var benchMarkInput []int
var inputTestData = [...]testData{
	testData{nums: [2]int{4, 5}, result: 1},
	testData{nums: [2]int{90, 45}, result: 45},
	testData{nums: [2]int{190, 145}, result: 5},
	testData{nums: [2]int{353, 793}, result: 1},
	testData{nums: [2]int{3534527, 7939832}, result: 113},
}

func (t testData) String() string {
	return fmt.Sprintf("Test Data gcd(%d,%d) = %d", t.nums[0], t.nums[1], t.result)
}

func init() {
	benchMarkInput = append(benchMarkInput, randUtil.GenerateRandomInt(10000000000))
	benchMarkInput = append(benchMarkInput, randUtil.GenerateRandomInt(10000000000))
}

func TestGCDNaiveWith0(t *testing.T) {
	num := []int{1, 0}
	assert := assert.New(t)
	assert.Equal(1, algo.GCDNaive(num[0], num[1]),
		"Failed 0 Test")
	t.Log("Naive Assertion OK")
}

func TestGCDWith0(t *testing.T) {
	num := []int{1, 0}
	assert := assert.New(t)
	assert.Equal(1, algo.GCD(num[0], num[1]),
		"Failed 0 Test")
	t.Log("Naive Assertion OK")
}

func TestGCD(t *testing.T) {
	assert := assert.New(t)
	for _, data := range inputTestData {
		t.Logf("Testing Data :%v\n", data)
		assert.Equal(data.result, algo.GCD(data.nums[0], data.nums[1]))
	}
}

func BenchmarkGCDNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.GCDNaive(benchMarkInput[0], benchMarkInput[1])
	}
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.GCD(benchMarkInput[0], benchMarkInput[1])
	}
}
