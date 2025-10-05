package database

import (
	"wordle/internal/game"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(postgres.Open("postgres://postgres:password@localhost:5432/wordle"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	err = db.AutoMigrate(&game.Game{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	return &Database{db: db}
}

func (d *Database) CreateGame(game *game.Game) error {
	return d.db.Create(game).Error
}

func (d *Database) GetGame(id string) (*game.Game, error) {
	var game game.Game
	err := d.db.Where("id = ?", id).First(&game).Error
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (d *Database) UpdateGame(game *game.Game) error {
	return d.db.Save(game).Error
}

func (d *Database) GetGamesByUsername(username string) ([]game.Game, error) {
	var games []game.Game
	err := d.db.Where("username = ?", username).Find(&games).Error
	return games, err
}
