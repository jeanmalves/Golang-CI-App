package addition

import "testing"

func TestAdd(t *testing.T) {
    res := Add(5,5)

    if res != 10 {
        t.Errorf("Add(5,5) = %d, deve ser 10", res)
    }
}