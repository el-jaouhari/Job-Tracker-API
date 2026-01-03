package service

import (
	"strings"

	"github.com/el-jaouhari/Job-Tracker-API/internal/repository"
	"gorm.io/gorm"
)

type Job struct {
	ID                uint   `json:"id,omitempty"`
	Title             string `json:"title"`
	Company           string `json:"company"`
	Url               string `json:"url"`
	Location          string `json:"location"`
	Type              string `json:"type"`
	ApplicationStatus string `json:"application_status"`
}

type JobsService struct {
	jobsRepository *repository.JobsRepository
}

func NewJobsService(jobsRepository *repository.JobsRepository) *JobsService {
	return &JobsService{
		jobsRepository: jobsRepository,
	}
}

func (s *JobsService) CreateJob(job *Job) error {
	// Validate required fields
	if strings.TrimSpace(job.Title) == "" {
		return &ValidationError{Field: "title", Message: "title is required"}
	}
	if strings.TrimSpace(job.Company) == "" {
		return &ValidationError{Field: "company", Message: "company is required"}
	}
	if strings.TrimSpace(job.Url) == "" {
		return &ValidationError{Field: "url", Message: "url is required"}
	}
	if strings.TrimSpace(job.Location) == "" {
		return &ValidationError{Field: "location", Message: "location is required"}
	}
	if strings.TrimSpace(job.Type) == "" {
		return &ValidationError{Field: "type", Message: "type is required"}
	}
	if strings.TrimSpace(job.ApplicationStatus) == "" {
		return &ValidationError{Field: "application_status", Message: "application_status is required"}
	}

	// Validate application status
	if !isValidStatus(job.ApplicationStatus) {
		return &StatusError{
			Status:      job.ApplicationStatus,
			ValidStatus: []string{"applied", "interviewing", "rejected", "offer"},
		}
	}

	return s.jobsRepository.CreateJob(&repository.Job{
		Title:             job.Title,
		Company:           job.Company,
		Url:               job.Url,
		Location:          job.Location,
		Type:              job.Type,
		ApplicationStatus: repository.ApplicationStatus(job.ApplicationStatus),
	})
}

func (s *JobsService) GetJob(id string) (*Job, error) {
	if strings.TrimSpace(id) == "" {
		return nil, ErrInvalidID
	}

	job, err := s.jobsRepository.GetJob(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrJobNotFound
		}
		return nil, err
	}
	return &Job{
		ID:                job.ID,
		Title:             job.Title,
		Company:           job.Company,
		Url:               job.Url,
		Location:          job.Location,
		Type:              job.Type,
		ApplicationStatus: string(job.ApplicationStatus),
	}, nil
}

func (s *JobsService) GetJobs() ([]Job, error) {
	jobs, err := s.jobsRepository.GetJobs()
	if err != nil {
		return nil, err
	}
	var jobsResponse []Job
	for _, job := range jobs {
		jobsResponse = append(jobsResponse, Job{
			ID:                job.ID,
			Title:             job.Title,
			Company:           job.Company,
			Url:               job.Url,
			Location:          job.Location,
			Type:              job.Type,
			ApplicationStatus: string(job.ApplicationStatus),
		})
	}
	return jobsResponse, nil
}

func (s *JobsService) UpdateJobStatus(id string, status string) error {
	if strings.TrimSpace(id) == "" {
		return ErrInvalidID
	}

	if strings.TrimSpace(status) == "" {
		return &ValidationError{Field: "status", Message: "status query parameter is required"}
	}

	if !isValidStatus(status) {
		return &StatusError{
			Status:      status,
			ValidStatus: []string{"applied", "interviewing", "rejected", "offer"},
		}
	}

	err := s.jobsRepository.UpdateJobStatus(id, repository.ApplicationStatus(status))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrJobNotFound
		}

		if strings.Contains(err.Error(), "invalid input value for enum") {
			return &StatusError{
				Status:      status,
				ValidStatus: []string{"applied", "interviewing", "rejected", "offer"},
			}
		}
		return err
	}

	return nil
}

func (s *JobsService) DeleteJob(id string) error {
	if strings.TrimSpace(id) == "" {
		return ErrInvalidID
	}

	err := s.jobsRepository.DeleteJob(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrJobNotFound
		}
		return err
	}

	return nil
}

func isValidStatus(status string) bool {
	validStatuses := []string{
		string(repository.ApplicationStatusApplied),
		string(repository.ApplicationStatusIntervieweing),
		string(repository.ApplicationStatusRejected),
		string(repository.ApplicationStatusOffer),
	}

	status = strings.ToLower(strings.TrimSpace(status))
	for _, valid := range validStatuses {
		if status == valid {
			return true
		}
	}
	return false
}
