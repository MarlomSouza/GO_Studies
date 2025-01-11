package internalmock

import (
	"emailn/internal/contract"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (r *CampaignServiceMock) Create(dto contract.NewCampaignDto) (string, error) {
	args := r.Called(dto)
	return args.String(0), args.Error(1)

}
func (r *CampaignServiceMock) Get() ([]contract.CampaignDto, error) {
	args := r.Called()
	return args.Get(0).([]contract.CampaignDto), args.Error(1)
}

func (r *CampaignServiceMock) GetById(id string) (*contract.CampaignDto, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignDto), args.Error(1)
}

func (r *CampaignServiceMock) Cancel(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
