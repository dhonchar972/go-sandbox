package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Setup!")
	res := m.Run()
	fmt.Println("Tear-down!")

	os.Exit(res)
}

func TestMultiple(t *testing.T) { // base signatur
	// var x, y, result = 2, 2, 4

	// realResult := Multiple(x, y)

	// if realResult != result {
	// 	t.Errorf("%d != %d", result, realResult)
	// }

	t.Run("first", func(t *testing.T) {
		t.Parallel()
		t.Log("Test one started!")

		var x, y, result = 2, 2, 4

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}
	})

	t.Run("second", func(t *testing.T) {
		t.Parallel()
		t.Log("Test two started!")

		var x, y, result = 746, 436, 325_256

		realResult := Multiple(x, y)

		if realResult != result {
			t.Errorf("%d != %d", result, realResult)
		}

		t.Run("negative", func(t *testing.T) {
			t.Log("Inner test started!")

			var x, y, result = -2, 4, -8

			realResult := Multiple(x, y)

			if realResult != result {
				t.Errorf("%d != %d", result, realResult)
			}
		})
	})

	t.Cleanup(func() {
		// some action
	})

}
