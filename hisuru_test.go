/*
  Copyright 2015 Adrian Stanescu. All rights reserved.
  Use of this source code is governed by the MIT License (MIT) that can be found in the LICENSE file.
 */

package hisuru

import (
  "testing"
)

func TestCompare(t *testing.T) {

  x := "1"
  y := "1.0.1"
  z := "1.0"
  a := ".1.1"
  b := ""
  c := "."
  d := "-1"
  e := "1.0.2"

  if Compare(x, y) != 1 {
    t.Errorf("Compare(%q, %q) != %q", x, y, 1)
  }

  if Compare(x, z) != 0 {
    t.Errorf("Compare(%q, %q) != %q", x, z, 0)
  }

  if Compare(x, a) != -1 {
    t.Errorf("Compare(%q, %q) != %q", x, a, -1)
  }

  if Compare(x, b) != -1 {
    t.Errorf("Compare(%q, %q) != %q", x, b, -1)
  }

  if Compare(x, c) != -1 {
    t.Errorf("Compare(%q, %q) != %q", x, c, -1)
  }

  if Compare(x, d) != -1 {
    t.Errorf("Compare(%q, %q) != %q", x, d, -1)
  }

  if Compare(d, x) != 1 {
    t.Errorf("Compare(%q, %q) != %q", x, y, 1)
  }

  if Compare(y, e) != 1 {
    t.Errorf("Compare(%q, %q) != %q", x, y, 1)
  }

}
