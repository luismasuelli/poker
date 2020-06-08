package french

var faces = [64]string{
	"2c", "3c", "4c", "5c", "6c", "7c", "8c", "9c", "Tc", "Jc", "Qc", "Kc", "Ac",
	"2h", "3h", "4h", "5h", "6h", "7h", "8h", "9h", "Th", "Jh", "Qh", "Kh", "Ah",
	"2d", "3d", "4d", "5d", "6d", "7d", "8d", "9d", "Td", "Jd", "Qd", "Kd", "Ad",
	"2s", "3s", "4s", "5s", "6s", "7s", "8s", "9s", "Ts", "Js", "Qs", "Ks", "As",
	"*w", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!", "!!",
}

const (
	C2 Card = iota
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	CT
	CJ
	CQ
	CK
	CA
	H2
	H3
	H4
	H5
	H6
	H7
	H8
	H9
	HT
	HJ
	HQ
	HK
	HA
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	D9
	DT
	DJ
	DQ
	DK
	DA
	S2
	S3
	S4
	S5
	S6
	S7
	S8
	S9
	ST
	SJ
	SQ
	SK
	SA
	W_
)

// Defines a french card (1 out of 53, since the
// wildcards also count).
type Card uint8

func (card Card) String() string {
	return faces[card]
}

func (card Card) Face() string {
	return faces[card]
}

func (card Card) Set() uint64 {
	return (1 << 53) - 1
}
