package repository

import (
	"time"

	"gorm.io/gorm"
)

type JobsRepository struct {
	db *gorm.DB
}

type ApplicationStatus string

const (
	ApplicationStatusApplied       ApplicationStatus = "applied"
	ApplicationStatusIntervieweing ApplicationStatus = "interviewing"
	ApplicationStatusRejected      ApplicationStatus = "rejected"
	ApplicationStatusOffer         ApplicationStatus = "offer"
)

type Job struct {
	ID                uint              `gorm:"primaryKey" json:"id"`
	Title             string            `json:"title"`
	Company           string            `json:"company"`
	Url               string            `json:"url"`
	Location          string            `json:"location"`
	Type              string            `json:"type"`
	ApplicationStatus ApplicationStatus `gorm:"type:application_status_enum" json:"application_status"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

func (Job) TableName() string {
	return "jobs"
}

func NewJobsRepository(db *gorm.DB) *JobsRepository {
	return &JobsRepository{
		db: db,
	}
}

func (r *JobsRepository) CreateJob(job *Job) error {
	return r.db.Create(job).Error
}

func (r *JobsRepository) GetJob(id string) (*Job, error) {
	var job Job
	if err := r.db.First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobsRepository) GetJobs() ([]Job, error) {
	var jobs []Job
	if err := r.db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

func (r *JobsRepository) UpdateJobStatus(id string, status ApplicationStatus) error {
	result := r.db.Model(&Job{}).Where("id = ?", id).Update("application_status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *JobsRepository) DeleteJob(id string) error {
	result := r.db.Delete(&Job{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
