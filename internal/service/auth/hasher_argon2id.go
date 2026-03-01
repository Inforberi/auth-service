package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Argon2idHasher struct {
	// Параметры Argon2id
	Time    uint32 // t
	Memory  uint32 // m (KiB)
	Threads uint8  // p
	KeyLen  uint32 // bytes

	// Длина соли
	SaltLen uint32 // bytes

	// Защита от DoS: ограничить длину пароля
	MaxPasswordLen int
}

func (h Argon2idHasher) Hash(password string) (string, error) {
	if h.MaxPasswordLen > 0 && len(password) > h.MaxPasswordLen {
		return "", errors.New("password too long")
	}
	if h.SaltLen == 0 {
		h.SaltLen = 16
	}
	if h.KeyLen == 0 {
		h.KeyLen = 32
	}
	if h.Time == 0 {
		h.Time = 1
	}
	if h.Memory == 0 {
		h.Memory = 46 * 1024 // 46 MiB в KiB (пример из OWASP)
	}
	if h.Threads == 0 {
		h.Threads = 1
	}

	salt := make([]byte, h.SaltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	sum := argon2.IDKey([]byte(password), salt, h.Time, h.Memory, h.Threads, h.KeyLen)

	// PHC-подобная строка (удобно хранить в TEXT)
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	sumB64 := base64.RawStdEncoding.EncodeToString(sum)

	// v=19 — версия Argon2 (фиксирована для argon2.IDKey)
	return fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s", h.Memory, h.Time, h.Threads, saltB64, sumB64), nil
}

func (h Argon2idHasher) Compare(storedHash, password string) bool {
	mem, tim, thr, salt, expected, err := parseArgon2idHash(storedHash)
	if err != nil {
		return false
	}

	// важно: использовать параметры из сохранённого хеша
	got := argon2.IDKey([]byte(password), salt, tim, mem, thr, uint32(len(expected)))

	return subtle.ConstantTimeCompare(got, expected) == 1
}

func parseArgon2idHash(s string) (mem uint32, tim uint32, thr uint8, salt []byte, sum []byte, err error) {
	parts := strings.Split(s, "$")
	// parts: ["", "argon2id", "v=19", "m=..,t=..,p=..", "<salt>", "<hash>"]
	if len(parts) != 6 || parts[1] != "argon2id" || !strings.HasPrefix(parts[2], "v=") {
		return 0, 0, 0, nil, nil, ErrHashFormat
	}

	// параметры
	mem, tim, thr, err = parseParams(parts[3])
	if err != nil {
		return 0, 0, 0, nil, nil, ErrHashFormat
	}

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return 0, 0, 0, nil, nil, ErrHashFormat
	}
	sum, err = base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return 0, 0, 0, nil, nil, ErrHashFormat
	}

	return mem, tim, thr, salt, sum, nil
}

func parseParams(p string) (mem uint32, tim uint32, thr uint8, err error) {
	// "m=19456,t=2,p=1"
	var m, t, pp string

	for _, kv := range strings.Split(p, ",") {
		kv = strings.TrimSpace(kv)
		if strings.HasPrefix(kv, "m=") {
			m = strings.TrimPrefix(kv, "m=")
		} else if strings.HasPrefix(kv, "t=") {
			t = strings.TrimPrefix(kv, "t=")
		} else if strings.HasPrefix(kv, "p=") {
			pp = strings.TrimPrefix(kv, "p=")
		}
	}
	if m == "" || t == "" || pp == "" {
		return 0, 0, 0, ErrHashFormat
	}

	mv, err := strconv.ParseUint(m, 10, 32)
	if err != nil {
		return 0, 0, 0, ErrHashFormat
	}
	tv, err := strconv.ParseUint(t, 10, 32)
	if err != nil {
		return 0, 0, 0, ErrHashFormat
	}
	pv, err := strconv.ParseUint(pp, 10, 8)
	if err != nil {
		return 0, 0, 0, ErrHashFormat
	}

	return uint32(mv), uint32(tv), uint8(pv), nil
}
