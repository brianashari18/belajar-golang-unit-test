package helper

import (
	"fmt"
	"runtime"
	"testing"

	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Brian",
			request: "Brian",
		},
		{
			name:    "Anashari",
			request: "Anashari",
		},
		{
			name:    "Puyol",
			request: "Puyol",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Brian", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Brian")
		}
	})
	b.Run("Anashari", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Anashari")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hello Brian")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Brian",
			request:  "Brian",
			expected: "Hello Brian",
		},
		{
			name:     "Puyol",
			request:  "Puyol",
			expected: "Hello Puyol",
		},
		{
			name:     "Anashari",
			request:  "Anashari",
			expected: "Hello Anashari",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// sub test
func TestSubTest(t *testing.T) {
	t.Run("Brian", func(t *testing.T) {
		result := HelloWorld("Brian")
		require.Equal(t, "Hello Brian", result, "Result is not 'Hello Brian'")
	})
	t.Run("Anashari", func(t *testing.T) {
		result := HelloWorld("Anashari")
		require.Equal(t, "Hello Anashari", result, "Result is not 'Hello Anashari'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("SEBELUM UNIT TEST")

	m.Run()

	fmt.Println("SETELAH UNIT TEST")
}

// skip
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can run on Windows")
	}
}

// error and fail
func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Brian")
	if result != "Hello Brian" {
		// t.Fail()
		t.Error("Result is not 'Hello Brian'")
	}
	fmt.Println("TestHelloWorld done")
}

// fatal and failnow
func TestHelloWorldAnashari(t *testing.T) {
	result := HelloWorld("Anashari")
	if result != "Hello Anashari" {
		// t.FailNow()
		t.Fatal("Result is not 'Hello Anashari'")
	}
	fmt.Println("TestHelloWorldAnashari done")
}

// assert vs require
func TestHelloWorldPuyol(t *testing.T) {
	result := HelloWorld("Puyol")
	// assert.Equal(t, "Hello Puyol", result, "Result is not 'Hello Puyol'")
	require.Equal(t, "Hello Puyol", result, "Result is not 'Hello Puyol'")
	fmt.Println("TestHelloWorldPuyol with assert done")
}
