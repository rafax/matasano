package utils

import "testing"

func TestHammingDistance(t *testing.T) {
	distance := HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	if distance != 37 {
		t.Errorf("Unexpected distance: %v != 37", distance)
	}
}

func TestFindMin_Standard(t *testing.T) {
	in := map[KeyEncoding]int{
		KeyEncoding{Key: "a"}: 123,
		KeyEncoding{Key: "b"}: -1,
		KeyEncoding{Key: "c"}: 122,
		KeyEncoding{Key: "d"}: 0,
	}
	k, v := findMin(in)
	if k.Key != "b" || v != -1 {
		t.Errorf("Unexpected minimum %v -> %v", k, v)
	}
}

func TestFindMin_One(t *testing.T) {
	in := map[KeyEncoding]int{KeyEncoding{Key: "one"}: 123}
	k, v := findMin(in)
	if k.Key != "one" || v != 123 {
		t.Errorf("Unexpected minimum %v -> %v", k, v)
	}
}

func TestXor(t *testing.T) {
	res := Xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	if res != "746865206b696420646f6e277420706c6179" {
		t.Errorf("Unexpected result: %v", res)
	}
}
