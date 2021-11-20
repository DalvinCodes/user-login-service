package pkg

import "golang.org/x/crypto/bcrypt"

type PasswordRepo interface {
	GenerateHashedPassword(s string) ([]byte, error)
	Compare(p1, p2 string) error
	ConvertToString(hashedPassword []byte) string
}

type password struct{}

//NewPasswordService returns a password service
func NewPasswordService() PasswordRepo {
	return &password{}
}

// GenerateHashedPassword returns a hashed password and an error
func (p *password) GenerateHashedPassword(s string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), 12)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

// Compare compares two passwords
func (p *password) Compare(p1, p2 string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(p1), []byte(p2)); err != nil {
		return err
	}
	return nil
}

func (p *password) ConvertToString(hashedPassword []byte) string {
	pwd := string(hashedPassword)
	return pwd
}
