package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExampleCleanup(t *testing.T) {
	Convey("Something works properly", t, func() {
		So(1, ShouldEqual, 1)   // assert
		So(2*2, ShouldEqual, 4) // assert
		Convey("More test", func() {
			So(1, ShouldEqual, 1)   // assert
			So(2*2, ShouldEqual, 4) // assert
			Convey("Test A", func() {
				So(1, ShouldEqual, 1)   // assert
				So(2*2, ShouldEqual, 4) // assert
			})
			Convey("Test B", func() {
				So(1, ShouldEqual, 1)   // assert
				So(2*2, ShouldEqual, 4) // assert
			})
			Convey("Test C", func() {
				So(1, ShouldEqual, 1)   // assert
				So(2*2, ShouldEqual, 4) // assert
			})
			Convey("Test D", func() {
				panic("PANIC!!!!")
			})
		})
		Reset(func() {
			t.Log("Finish!")
		})
	})
}

// $GOPATH/bin/goconvey - web UI(web server) for tests!!!
