package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email@e.com", "email2@e.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	// teste feito utilizando biblioteca externa Testify
	// assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

	// teste feito na mão
	// if campaign.ID != "1" {
	// 	t.Errorf("Expected 1")
	// } else if campaign.Name != name {
	// 	t.Errorf("Expected correct name")
	// } else if campaign.Content != content {
	// 	t.Errorf("Expected correct content")
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Errorf("Expected correct contacts")
	// }
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, erro := NewCampaign("", content, contacts)

	assert.Equal("name is required", erro.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, erro := NewCampaign(name, "", contacts)

	assert.Equal("content is required", erro.Error())
}

func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, erro := NewCampaign(name, content, []string{})

	assert.Equal("contacts is required", erro.Error())
}
