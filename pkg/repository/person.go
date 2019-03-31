package repository

import (
	"github.com/Rosaniline/gorm-ut/pkg/model"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Repository interface {
	Get(id uuid.UUID) (*model.Person, error)
	Create(id uuid.UUID, name string) error
}

type repo struct {
	DB *gorm.DB
}

func (p *repo) Create(id uuid.UUID, name string) error {
	person := &model.Person{
		ID:   id,
		Name: name,
	}

	return p.DB.Create(person).Error
}

func (p *repo) Get(id uuid.UUID) (*model.Person, error) {
	person := new(model.Person)

	err := p.DB.Where("id = ?", id).Find(person).Error

	return person, err
}

func CreateRepository(db *gorm.DB) Repository {
	return &repo{
		DB: db,
	}
}
