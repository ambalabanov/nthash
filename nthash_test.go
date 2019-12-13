package nthash

import (
	"testing"
)

func TestGetHash(t *testing.T) {
	if GetHash("pass1!@#$") != "aba2fd935f15ccdd099a36bb777ba0f1" {
		t.Errorf("Hash was incorrect")
	}
}
func TestCheckHash(t *testing.T) {
	if !CheckHash("p@ssw0rd", "de26cce0356891a4a020e7c4957afc72") {
		t.Errorf("Hash doesn't match")
	}

}
