package string

import "testing"

func Test(t *testing.T) {
  var test = []struct {
    s, want string
  } {
    {"Backward", "drawkcaB"},
    {"世界", "界世"},
    {"", ""},
  }
  for _, c := range test {
    got := Reverse(c.s)
    if got != c.want {
      t.Error("Reverse(%q) == %q, want %q", c.s, got, c.want)
    }
  }

}

