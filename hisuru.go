/*
  Copyright 2015 Adrian Stanescu. All rights reserved.
  Use of this source code is governed by the MIT License (MIT) that can be found in the LICENSE file.
 */

/*
  hisuru
  Go program that compares software versions in the x.y.z format

  Usage:
  x := "1"
  y := "1.0.1"
  z := "1.0"
  fmt.Println(hisuru.Compare(x, y)) // 1 = y
  fmt.Println(hisuru.Compare(x, z)) // 0 = equal
  fmt.Println(hisuru.Compare(x, a)) // -1 = x
 */

package hisuru

import (
//  "os"
  "strconv"
  "strings"
)

/*
  compare two versions in x.y.z form
  @param  {string} a     version string
  @param  {string} b     version string
  @return {int}          -1 = a is higher, 0 = equal, 1 = b is higher
 */

func Compare(a, b string) (ret int) {
  as        := strings.Split(a, ".")
  bs        := strings.Split(b, ".")
  loopMax   := len(bs)

  if len(as) > len(bs) {
    loopMax = len(as)
  }

  for i := 0; i < loopMax; i++ {
    var x, y string

    if len(as) > i {
      x = as[i]
    }

    if len(bs) > i {
      y = bs[i]
    }

    xi,_    := strconv.Atoi(x)
    yi,_    := strconv.Atoi(y)

    if xi > yi {
      ret = -1
    } else if xi < yi {
      ret = 1
    }

    if ret != 0 {
      break
    }
  }
  return
}
