package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

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
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}
