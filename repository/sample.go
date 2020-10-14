package repository

import (
	"github.com/RaptorGandalf/hello-go-api/model"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// SampleRepository - Interface for creating a SampleRepository struct to work with Sample models in the database
type SampleRepository interface {
	GetAll() (*[]model.Sample, error)
	Get(uuid.UUID) (*model.Sample, error)
	Create(sample *model.Sample) error
	Update(sample *model.Sample) error
	Delete(uuid.UUID) error
}

type sampleRepository struct {
	DB *gorm.DB
}

// GetSampleRepository - Returns a sample repository struct
func GetSampleRepository(db *gorm.DB) SampleRepository {
	return &sampleRepository{
		DB: db,
	}
}

// GetAll - Returns all Samples
func (r *sampleRepository) GetAll() (*[]model.Sample, error) {
	var samples []model.Sample

	err := r.DB.Find(&samples).Error

	return &samples, err
}

// Get - Returns a single Sample with the specified ID
func (r *sampleRepository) Get(id uuid.UUID) (*model.Sample, error) {
	var sample model.Sample

	err := r.DB.Where("id = ?", id).Take(&sample).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &sample, err
}

// Create - Creates a new Sample in the database
func (r *sampleRepository) Create(sample *model.Sample) error {
	sample.ID = uuid.New()

	return r.DB.Create(sample).Error
}

// Update - Updates an existing Sample in the database
func (r *sampleRepository) Update(sample *model.Sample) error {
	return r.DB.Save(sample).Error
}

// Delete - Deletes a Sample from the database
func (r *sampleRepository) Delete(id uuid.UUID) error {
	sample, err := r.Get(id)
	if err != nil {
		return err
	}

	return r.DB.Delete(&sample).Error
}
