package util

import "math/rand"

type RandomNicknameGenerator struct {
	Length int
}

func NewRandomNicknameGenerator(length int) *RandomNicknameGenerator {
	return &RandomNicknameGenerator{
		Length: length,
	}
}

func (g *RandomNicknameGenerator) Generate() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	nickname := make([]byte, g.Length)
	for i := range nickname {
		nickname[i] = charset[rand.Intn(len(charset))]
	}
	return string(nickname)
}
