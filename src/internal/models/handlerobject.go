package models

import (
	"time"
)

// Functions for all types, to make them HandlerObjects
type HandlerObject interface {
	GetId() string
	SetCreatedAt()
}

func (u *User) GetId() string {
	return u.Id
}

func (a *Admin) GetId() string {
	return a.Id
}

func (b *Bird) GetId() string {
	return b.Id
}

func (p *Post) GetId() string {
	return p.Id
}

func (m *Media) GetId() string {
	return m.Id
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now().Format(time.RFC3339)
}

func (a *Admin) SetCreatedAt() {
}

func (b *Bird) SetCreatedAt() {
}

func (p *Post) SetCreatedAt() {
	p.CreatedAt = time.Now().Format(time.RFC3339)
}

func (m *Media) SetCreatedAt() {
}
