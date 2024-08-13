package repositories

import (
	"day06/internal/models"
	postgresql "day06/internal/storage/postgre"
	"fmt"

	"github.com/google/uuid"
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

func (r *PostRepository) Count() (*int64, error) {
	if err := postgresql.Ping(r.db); err != nil {
		return nil, err
	}
	var count int64
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		res := d.Model(models.Post{}).Count(&count)
		if res.Error != nil {
			return res.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &count, nil
}

func (r *PostRepository) Create(model *models.Post) (uuid *string, err error) {
	if err := postgresql.Ping(r.db); err != nil {
		return nil, err
	}
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		if err = d.AutoMigrate(); err != nil {
			return err
		}
		if result := d.Create(model); result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &model.Id, nil
}

func (r *PostRepository) Get(uuid string) (*models.Post, error) {
	if err := postgresql.Ping(r.db); err != nil {
		return nil, err
	}
	var model models.Post
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		res := d.First(&model, "id = ?", uuid)
		if res.Error != nil {
			return res.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &model, nil
}

func (r *PostRepository) GetAll(limit int, offset int) ([]models.Post, error) {
	if err := postgresql.Ping(r.db); err != nil {
		return nil, err
	}
	var models []models.Post
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		res := d.Limit(limit).Offset(offset).Find(&models)
		if res.Error != nil {
			return res.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return models, nil
}

func (r *PostRepository) Update(model *models.Post) error {
	if err := postgresql.Ping(r.db); err != nil {
		return err
	}
	if err := postgresql.TxSaveExecutor(r.db, func(d *gorm.DB) error {
		ok := d.First(model)
		if ok != nil {
			return ok.Error
		}
		ok = d.Model(model).Updates(models.Post{Description: model.Description, ShortDescription: model.ShortDescription})
		if ok != nil {
			return ok.Error
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Delete(uuid string) error {
	if err := postgresql.Ping(r.db); err != nil {
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
		return err
	}
	return nil
}

func (r *PostRepository) Init10Posts() error {
	var err error
	for i := 1; i < 11; i++ {
		model := models.Post{
			Id:               uuid.New().String(),
			ShortDescription: fmt.Sprintf("This is short description of post %d", i),
			Description:      fmt.Sprintf("* This is the full description of post %d.\n # It contains more detailed information. ``` Hello world ```", i),
		}
		_, err = r.Create(&model)
	}
	return err
}
