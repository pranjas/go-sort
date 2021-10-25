package lcm

import (
	"pks_sort/algo"
	"pks_sort/randUtil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
}

type testData struct {
	nums   [2]int
	result int
}

var benchMarkInput []int

var inputTestData = [...]testData{
	testData{nums: [2]int{4, 5}, result: 20},
	testData{nums: [2]int{3, 5}, result: 15},
	testData{nums: [2]int{30, 10}, result: 30},
	testData{nums: [2]int{127, 19}, result: 2413},
	testData{nums: [2]int{23120, 32098}, result: 371052880},
	testData{nums: [2]int{127, 19}, result: 2413},
}

func TestLCMNaive(t *testing.T) {
	assert := assert.New(t)
	for _, data := range inputTestData {
		t.Logf("Testing Data: %v\n", data)
		assert.Equal(data.result, algo.LCMNaive(data.nums[0], data.nums[1]))
	}
}

func TestLCM(t *testing.T) {
	assert := assert.New(t)
	for _, data := range inputTestData {
		t.Logf("Testing Data: %v\n", data)
		assert.Equal(data.result, algo.LCMFast(data.nums[0], data.nums[1]))
	}
}

func init() {
	benchMarkInput = append(benchMarkInput, randUtil.GenerateRandomInt(1000000))
	benchMarkInput = append(benchMarkInput, randUtil.GenerateRandomInt(1000000))
}

func BenchmarkLCMNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.LCMNaive(benchMarkInput[0], benchMarkInput[1])
	}
}

func BenchmarkLCMFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.LCMFast(benchMarkInput[0], benchMarkInput[1])
	}
}
