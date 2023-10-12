package models

import (
	"time"
)

// Functions for all types, to make them HandlerObjects
type HandlerObject interface {
	GetId() string
	SetCreatedAt()
}

func (u *UserDB) GetId() string {
	return u.Id
}

func (a *AdminDB) GetId() string {
	return a.Id
}

func (b *BirdDB) GetId() string {
	return b.Id
}

func (p *PostDB) GetId() string {
	return p.Id
}

func (m *MediaDB) GetId() string {
	return m.Id
}

func (u *UserDB) SetCreatedAt() {
	u.CreatedAt = time.Now().Format(time.RFC3339)
}

func (a *AdminDB) SetCreatedAt() {
}

func (b *BirdDB) SetCreatedAt() {
}

func (p *PostDB) SetCreatedAt() {
	p.CreatedAt = time.Now().Format(time.RFC3339)
}

func (m *MediaDB) SetCreatedAt() {
}
