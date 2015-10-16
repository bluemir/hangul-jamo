package main

type Hangul struct {
	ChoSeong  byte
	JungSeong byte
	JongSeong byte
}
type HangulString []Hangul

func decompose(input string) (result []Hangul) {
	for _, ch := range input {
		result = append(result, Hangul{
			byte((ch - 0xAC00) / (28 * 21)),
			byte((ch - 0xAC00) / 28 % 21),
			byte((ch - 0xAC00) % 28),
		})
	}
	return
}
func (this Hangul) Compare(that Hangul) int {
	if this.ChoSeong > that.ChoSeong {
		return 1
	} else if this.ChoSeong < that.ChoSeong {
		return -1
	}
	if this.JungSeong > that.JungSeong {
		return 1
	} else if this.JungSeong < this.JungSeong {
		return -1
	}
	if this.JongSeong > that.JongSeong {
		return 1
	} else if this.JongSeong < that.JongSeong {
		return -1
	}
	return 0
}
func (this Hangul) toBytes() (result []byte) {
	result = append(result, this.ChoSeong, this.JungSeong, this.JongSeong)
	return
}

func (this HangulString) Compare(that HangulString) int {
	if c := len(this) - len(that); c != 0 {
		if c > 0 {
			return 1
		} else {
			return -1
		}
	}
	for i := 0; i < len(this); i++ {
		if c := this[i].Compare(that[i]); c != 0 {
			return c
		}
	}
	return 0
}
func (this HangulString) toBytes() (result []byte) {
	for _, hangul := range this {
		result = append(result, hangul.toBytes()...)
	}
	return
}
