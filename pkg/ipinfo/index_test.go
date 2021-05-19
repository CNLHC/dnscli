package ipinfo

import (
	"fmt"
	"testing"
)

func TestJSONIP(t *testing.T) {
	obj, e := JsonIPGet()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(obj)
	obj, e = JSONIPGetV6()
	fmt.Println(obj, e)

}
