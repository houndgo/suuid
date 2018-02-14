package suuid

import (
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/satori/go.uuid"
)

const (
	//DefaultAlphabet is the default alphabet used.
	DefaultAlphabet = "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

// SUUID is pg Namespaces Mount
type SUUID struct {
	alphabet *StringSet
}

// New is SUUID method
func New() *SUUID {
	suid := &SUUID{}
	suid.SetAlphabet(DefaultAlphabet)
	return suid
}

// NewWithAlphabet is SUUID method
func NewWithAlphabet(alphabet string) *SUUID {

	suuid := &SUUID{}
	if alphabet == "" {
		alphabet = DefaultAlphabet
	}
	suuid.SetAlphabet(alphabet)
	return suuid
}

// SetAlphabet is SUUID method
func (s *SUUID) SetAlphabet(alphabet string) {
	set := NewStringSet()
	for _, a := range alphabet {
		set.Add(string(a))
	}
	set.Sort()
	s.alphabet = set
}

func (s SUUID) String() string {
	return s.UUID("")
}

// UUID is SUUID method
func (s *SUUID) UUID(name string) string {
	var _uuid uuid.UUID
	if name == "" {
		_uuid, _ = uuid.NewV4()
	} else if strings.HasPrefix(name, "http") {
		_uuid = uuid.NewV5(uuid.NamespaceDNS, name)
	} else {
		_uuid = uuid.NewV5(uuid.NamespaceURL, name)
	}

	return s.Encode(_uuid)
}

//Encode is Encodes a UUID into a string (LSB first) according to the alphabet
// If leftmost (MSB) bits 0, string might be shorter
func (s *SUUID) Encode(uuid uuid.UUID) string {
	padLen := s.encodeLen(len(uuid.Bytes()))
	number := uuidToInt(uuid)
	return s.numToString(number, padLen)
}

// Decode is
func (s *SUUID) Decode(input string) (uuid.UUID, error) {
	_uuid, err := uuid.FromString(s.stringToNum(input))
	return _uuid, err
}

func (s *SUUID) encodeLen(numBytes int) int {
	factor := math.Log(float64(25)) / math.Log(float64(s.alphabet.Len()))
	length := math.Ceil(factor * float64(numBytes))
	return int(length)
}

//Covert a number to a string, using the given alphabet.
func (s *SUUID) numToString(number *big.Int, padToLen int) string {
	output := ""
	var digit *big.Int
	for number.Uint64() > 0 {
		number, digit = new(big.Int).DivMod(number, big.NewInt(int64(s.alphabet.Len())), new(big.Int))
		output += s.alphabet.ItemByIndex(int(digit.Int64()))
	}
	if padToLen > 0 {
		remainer := math.Max(float64(padToLen)-float64(len(output)), 0)
		output = output + strings.Repeat(s.alphabet.ItemByIndex(0), int(remainer))
	}

	return output
}

// Convert a string to a number(based uuid string),using the given alphabet.
func (s *SUUID) stringToNum(input string) string {
	n := big.NewInt(0)
	for i := len(input) - 1; i >= 0; i-- {
		n.Mul(n, big.NewInt(int64(s.alphabet.Len())))
		n.Add(n, big.NewInt(int64(s.alphabet.Index(string(input[i])))))
	}

	x := fmt.Sprintf("%x", n)
	x = x[0:8] + "-" + x[8:12] + "-" + x[12:16] + "-" + x[16:20] + "-" + x[20:32]
	return x
}

func uuidToInt(_uuid uuid.UUID) *big.Int {
	var i big.Int
	i.SetString(strings.Replace(_uuid.String(), "-", "", 4), 16)
	return &i
}
