package repositories

import (
	"delta/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) Create(dashboard *models.Dashboard) error {
	return r.db.Create(dashboard).Error
}

func (r *DashboardRepository) GetByID(id uuid.UUID) (*models.Dashboard, error) {
	var dashboard models.Dashboard
	err := r.db.Preload("Panels").First(&dashboard, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("dashboard not found")
		}
		return nil, err
	}
	return &dashboard, nil
}

func (r *DashboardRepository) Update(dashboard *models.Dashboard) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(dashboard).Error
}

func (r *DashboardRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Dashboard{}, "id = ?", id).Error
}

func (r *DashboardRepository) UpdatePanel(dashboardID uuid.UUID, panel *models.Panel) error {
	return r.db.Model(&models.Dashboard{}).
		Where("id = ?", dashboardID).
		Association("Panels").
		Replace(panel)
}

func (r *DashboardRepository) GetAll() ([]models.Dashboard, error) {
	var dashboards []models.Dashboard
	err := r.db.Preload("Panels").Find(&dashboards).Error
	return dashboards, err
}
