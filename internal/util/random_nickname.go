package util

import "math/rand"

type RandomNicknameGenerator struct {
	Prefix string
	Length int
}

func NewRandomNicknameGenerator(Prefix string, length int) *RandomNicknameGenerator {
	return &RandomNicknameGenerator{
		Prefix: Prefix,
		Length: length,
	}
}

func (g *RandomNicknameGenerator) Generate() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nickname := make([]byte, g.Length)
	for i := range nickname {
		nickname[i] = charset[rand.Intn(len(charset))]
	}
	return g.Prefix + string(nickname)
}

