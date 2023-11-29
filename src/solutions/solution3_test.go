package solutions

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestMakeBirthday(t *testing.T) {
	type args struct {
		dogs []Dog
	}

	tests := []struct {
		name string
		args args
		want []Dog
	}{
		{
			name: "should mutate dogs slice with age + 1",
			args: args{
				[]Dog{
					{"O", 3},
					{"W", 4},
					{"A", 2},
				},
			},
			want: []Dog{
				{"O", 4},
				{"W", 5},
				{"A", 3},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := makeBirthday(test.args.dogs)
			assert.Equal(t, got, test.want)
		})
	}
}

func TestSortDogByAge(t *testing.T) {
	type args struct {
		dogs []Dog
	}

	tests := []struct {
		name string
		args args
		want []Dog
	}{
		{
			name: "should sort dogs slice with asc order",
			args: args{
				[]Dog{
					{"O", 3},
					{"W", 4},
					{"A", 2},
				},
			},
			want: []Dog{
				{"A", 2},
				{"O", 3},
				{"W", 4},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := sortDogByAge(test.args.dogs)
			assert.Equal(t, got, test.want)
		})
	}
}

var benchmankResult []Dog

func BenchmarkSortDogByName(b *testing.B) {
	dogs := generateTestData(1000)
	var result []Dog

	for i := 0; i < b.N; i++ {
		result = sortDogByName(dogs)
	}

	benchmankResult = result
}

func BenchmarkSortDogPtrByName(b *testing.B) {
	dogs := generateTestDataPtr(1000)

	for i := 0; i < b.N; i++ {
		sortDogPtrByName(dogs)
	}
}

func generateTestData(size int) []Dog {
	var testData []Dog
	names := []string{"Azor", "Puszek", "Reksio", "Pimpek", "Donald", "Włodek", "Kazimierz", "Burek", "Czarek", "Kuba"}

	for i := 0; i < size; i++ {
		testData = append(testData, Dog{names[rand.Intn(len(names))], uint32(rand.Intn(16))})
	}

	return testData
}

func generateTestDataPtr(size int) []*Dog {
	var testData []*Dog
	names := []string{"Azor", "Puszek", "Reksio", "Pimpek", "Donald", "Włodek", "Kazimierz", "Burek", "Czarek", "Kuba"}

	for i := 0; i < size; i++ {
		testData = append(testData, &Dog{names[rand.Intn(len(names))], uint32(rand.Intn(16))})
	}

	return testData
}
