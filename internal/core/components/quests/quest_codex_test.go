package quests

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces/mocks"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

type createQuestTestCase struct {
	name         string
	input        *types.CreateQuestInput
	mockRepoErr  error
	mockGameData game.GameData
	expectedMsg  string
}

var createQuestTests = []createQuestTestCase{
	{
		name: "Successful Quest Creation with Party Level 5",
		input: &types.CreateQuestInput{
			Title:       "Test Quest",
			QuestType:   "Hunt",
			Description: "A test quest for testing",
			Source:      "Test Source",
			Level:       1,
		},
		mockRepoErr: nil,
		mockGameData: game.GameData{
			PartyLevel: 5,
			PartySize:  3,
		},
		expectedMsg: "should not return an error",
	},
	{
		name: "Quest Creation with Existing Quest ID",
		input: &types.CreateQuestInput{
			Title:       "Test Quest",
			QuestType:   "Hunt",
			Description: "A test quest for testing",
			Source:      "Test Source",
			Level:       1,
		},
		mockRepoErr: errors.New("quest id already exists"),
		mockGameData: game.GameData{
			PartyLevel: 5,
			PartySize:  3,
		},
		expectedMsg: "Should return a quest id already exists error",
	},
}

func TestCreateNewQuest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range createQuestTests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewMockQuestRepository(ctrl)

			// Set expectations for the repo
			mockRepo.EXPECT().AddNewQuest(gomock.Any(), gomock.Any()).Return(tt.mockRepoErr)

			qc := NewQuestCodex(mockRepo, &tt.mockGameData)

			// Call the function
			_, err := qc.CreateNewQuest(tt.input)

			// Assert
			if tt.mockRepoErr != nil {
				assert.NotEmpty(t, err, tt.expectedMsg)
			} else {
				assert.NoError(t, err, tt.expectedMsg)
			}
		})
	}
}
