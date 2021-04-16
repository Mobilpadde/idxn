package idxn

const (
	_alphanumeric      = "abcdefghijklmnopqrstuvwxyz1234567890"
	_alphanumericUpper = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	_alpha             = "abcdefghijklmnopqrstuvwxyz"
	_alphaUpper        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func Alphanumeric(i int) string {
	return Custom(i, _alphanumeric)
}

func AlphanumericWithUpper(i int) string {
	return Custom(i, _alphanumericUpper)
}

func Alpha(i int) string {
	return Custom(i, _alpha)
}

func AlphaWithUpper(i int) string {
	return Custom(i, _alphaUpper)
}

func Custom(i int, alpha string) string {
	var idx string

	l := len(alpha)
	mod := i % l
	div := i / l

	idx += string(alpha[mod])
	for j := div - 1; j >= 0; j-- {
		idx += Custom(div-j-1, alpha)
	}

	return idx
}
