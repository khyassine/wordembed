package word2vec

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/unixpickle/serializer"
)

func TestHierarchy(t *testing.T) {
	actual := BuildHierarchy(map[string]float64{
		"a": 0.501,
		"b": 0.25,
		"c": 0.125,
		"d": 0.124,
	})
	expected := Hierarchy{
		"a": []int{-1},
		"b": []int{1, -2},
		"c": []int{1, 2, -3},
		"d": []int{1, 2, 3},
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func TestHierarchyLarge(t *testing.T) {
	words := map[string]float64{}
	for i := 0; i < 10000; i++ {
		words[fmt.Sprintf("%d", i/5)] = float64(rand.Intn(10)+1) / 10
	}
	hier := BuildHierarchy(words)
	if len(hier) != len(words) {
		t.Errorf("expected len %d but got len %d", len(words), len(hier))
	}
}

func TestHierarchySerialize(t *testing.T) {
	expected := Hierarchy{
		"a": []int{-1},
		"b": []int{1, -2},
		"c": []int{1, 2, -3},
		"d": []int{1, 2, 3},
	}
	data, err := serializer.SerializeAny(expected)
	if err != nil {
		t.Fatal(err)
	}
	var actual Hierarchy
	if err := serializer.DeserializeAny(data, &actual); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v but got %v", expected, actual)
	}
}

func BenchmarkHierarchy(b *testing.B) {
	words := map[string]float64{}
	for i := 0; i < 5000; i++ {
		words[fmt.Sprintf("%d", i)] = rand.NormFloat64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildHierarchy(words)
	}
}
