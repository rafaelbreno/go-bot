package repository

import (
	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/storage"
	"go.uber.org/zap"
)

// CommandRepo stores the functions
// related to Command actions.
type CommandRepo interface {
	Create(cmd entity.Command) (entity.Command, error)
	Update(cmd entity.Command) (entity.Command, error)
	Read(id string) (entity.Command, error)
	Delete(id string) error
}

// CommandRepoCtx represents
// the Commands's context.
type CommandRepoCtx struct {
	Storage *storage.Storage
	Logger  *zap.Logger
}

// Create insert a new command into
// the database.
func (cr CommandRepoCtx) Create(cmd entity.Command) (entity.Command, error) {
	if err := cr.Storage.SQL.Client.Create(&cmd).Error; err != nil {
		return entity.Command{}, err
	}
	return cmd, nil
}

// Read returns a Command from the
// given ID.
func (cr CommandRepoCtx) Read(id string) (entity.Command, error) {
	cmdFound := new(entity.Command)

	if err := cr.Storage.SQL.Client.First(&cmdFound, "id = ?", id).Error; err != nil {
		return entity.Command{}, err
	}

	return *cmdFound, nil
}

// Update - update a command with the given fields
// with the new values.
func (cr CommandRepoCtx) Update(cmd entity.Command) (entity.Command, error) {
	cmdNew := new(entity.Command)

	if err := cr.Storage.SQL.Client.First(&cmdNew, "id = ?", cmd.ID).Error; err != nil {
		return entity.Command{}, nil
	}

	cmdNew.UpdateFields(cmd)

	if err := cr.Storage.SQL.Client.Save(&cmdNew).Error; err != nil {
		return entity.Command{}, err
	}

	return *cmdNew, nil
}

// Delete removes a command from DB
// with given ID
func (cr CommandRepoCtx) Delete(id string) error {
	cmd := new(entity.Command)

	err := cr.Storage.SQL.Client.Where("id = ?", id).Delete(&cmd).Error

	return err
}
