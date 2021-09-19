package singleton

import (
	"testing"
)

func Benchmark_GetInstanceFunc_01(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceFunc()
	}
}

func Benchmark_GetInstanceSync_02(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceSync()
	}
}

func Benchmark_GetInstanceMutex_03(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceMutex()
	}
}

func Benchmark_GetInstanceInit_04(t *testing.B) {
	for i := 0; i < t.N; i++ {
		_ = GetInstanceInit()
	}
}
