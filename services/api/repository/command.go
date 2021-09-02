package repository

import (
	"os"

	"github.com/rafaelbreno/go-bot/api/config"
	"github.com/rafaelbreno/go-bot/api/entity"
	"github.com/rafaelbreno/go-bot/api/internal"
	"github.com/rafaelbreno/go-bot/api/storage"
)

// CommandRepo stores the functions
// related to Command actions.
type CommandRepo interface {
	Create(cmd entity.Command) (entity.Command, error)
	Update(id string, cmd entity.Command) (entity.Command, error)
	Read(id string) (entity.Command, error)
	Delete(id string) error
}

// CommandRepoCtx represents
// the Command's context.
type CommandRepoCtx struct {
	Storage *storage.Storage
	Ctx     *internal.Context
}

const (
	add    = "ADD"
	update = "UPDATE"
	delete = "DELETE"
)

// Create insert a new command into
// the database.
func (cr CommandRepoCtx) Create(cmd entity.Command) (entity.Command, error) {
	if err := cr.Storage.SQL.Client.Create(&cmd).Error; err != nil {
		cr.Storage.Ctx.Logger.Error(err.Error())
		return entity.Command{}, err
	}

	jsonCommand := cmd.ToJSON()

	b, err := jsonCommand.ToJSONString()

	if err != nil {
		cr.Ctx.Logger.Error(err.Error())
		os.Exit(0)
	}

	cr.Storage.KafkaClient.Produce([]byte(add), b)

	return cmd, nil
}

// Read returns a Command from the
// given ID.
func (cr CommandRepoCtx) Read(id string) (entity.Command, error) {
	cmdFound := new(entity.Command)

	if err := cr.Storage.SQL.Client.First(&cmdFound, "id = ?", id).Error; err != nil {
		cr.Ctx.Logger.Error(err.Error())
		return entity.Command{}, err
	}

	return *cmdFound, nil
}

// Update - update a command with the given fields
// with the new values.
func (cr CommandRepoCtx) Update(id string, cmd entity.Command) (entity.Command, error) {
	cmdNew := new(entity.Command)

	if err := cr.Storage.SQL.Client.First(&cmdNew, "id = ?", id).Error; err != nil {
		cr.Ctx.Logger.Error(err.Error())
		return entity.Command{}, err
	}

	cmdNew.UpdateFields(cmd)

	if err := cr.Storage.SQL.Client.Save(&cmdNew).Error; err != nil {
		cr.Ctx.Logger.Error(err.Error())
		return entity.Command{}, err
	}

	return *cmdNew, nil
}

// Delete removes a command from DB
// with given ID
func (cr CommandRepoCtx) Delete(id string) error {
	cmd := new(entity.Command)

	err := cr.Storage.SQL.Client.Where("id = ?", id).Delete(&cmd).Error

	if err != nil {
		cr.Ctx.Logger.Error(err.Error())
	}

	return err
}

// NewCommandRepoCtx creates and return a
// configured CommandRepoCtx.
func NewCommandRepoCtx() CommandRepoCtx {
	return CommandRepoCtx{
		Storage: config.Storage,
		Ctx:     config.Ctx,
	}
}
