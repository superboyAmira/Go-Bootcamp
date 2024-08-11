package repositories

import (
	"day06/internal/models"
	postgresql "day06/internal/storage/postgre"
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(Db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: Db,
	}
}

func (r *PostRepository) Create(model *models.Post,log *slog.Logger) (uuid *string, err error) {
	if err := postgresql.Ping(r.db); err != nil {
		log.Error("ping failed")
		return nil, err
	}
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		if err = d.AutoMigrate(); err != nil {
			log.Error("err with automigrate")
			return err
		}
		if result := d.Create(model); result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		log.Error("CREATE failed", "error", err.Error())
		return nil, err
	}
	log.Info("CREATED uuid", "uuid", model.Id)
	return &model.Id, nil
}

func (r *PostRepository) Get(uuid string, log *slog.Logger) (*models.Post, error) {
	if err := postgresql.Ping(r.db); err != nil {
		log.Error("ping failed")
		return nil, err
	}
	var model models.Post
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		var ok bool
		var res any
		if res, ok = d.Get(uuid); !ok {
			return errors.New("GET Error")
		}
		if model, ok = res.(models.Post); !ok {
			return errors.New("type assertion to models.Post failed")
		}
		return nil
	}); err != nil {
		log.Error("GET failed", "error", err.Error())
		return nil, err
	}
	log.Info("GETED uuid", "uuid", model.Id)
	return &model, nil
}
func (r *PostRepository) Update(model *models.Post, log *slog.Logger) error {
	if err := postgresql.Ping(r.db); err != nil {
		log.Error("ping failed")
		return err
	}
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		ok := d.First(model)
		if ok != nil {
			return ok.Error
		}
		ok = d.Model(model).Updates(models.Post{Decription: model.Decription, ShortDecription: model.ShortDecription})
		if ok != nil {
			return ok.Error
		}
		return nil
	}); err != nil {
		log.Error("UPDATE failed", "error", err.Error())
		return err
	}
	log.Info("UPDATED uuid", "uuid", model.Id)
	return nil
}

func (r *PostRepository) Delete(uuid string, log *slog.Logger) error {
	if err := postgresql.Ping(r.db); err != nil {
		log.Error("ping failed")
		return err
	}

	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		var model models.Post
		ok := d.First(&model, uuid)
		if ok != nil {
			return ok.Error
		}
		ok = d.Delete(model)
		if ok != nil {
			return ok.Error
		}
		return nil
	}); err != nil {
		log.Error("DELETE failed", "error", err.Error())
		return err
	}
	log.Info("DELETED uuid", "uuid", uuid)
	return nil
}