package opstr

import "testing"

func TestSplitBySpace(t *testing.T) {
	str := `     UDP      000-001 sec     1.29M   160.93K`
	s := SplitBySpace(str,4)
	if s != "160.93K" {
		t.Errorf("应该输出160.93K，但是输出%s",s)
	}
}