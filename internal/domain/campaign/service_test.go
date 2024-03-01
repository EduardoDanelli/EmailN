package campaign

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"teste1@teste.com"},
	}
	repository = new(repositoryMock)
	service    = Service{Repository: repository}
)

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())

}

func Test_Create_SaveCampaign(t *testing.T) {

	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name || campaign.Content != newCampaign.Content || len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = repository

	service.Create(newCampaign)

	repository.AssertExpectations(t)

}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))
	service.Repository = repository

	_, err := service.Create(newCampaign)
	assert.Equal("error to save on database", err.Error())

}
