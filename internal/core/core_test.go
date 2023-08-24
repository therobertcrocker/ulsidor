package core

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces/mocks"
)

type InitQuestComponentsTestCase struct {
	name         string
	mockRepoErr  error
	mockGameData *game.GameData
	expectedMsg  string
	wantErr      bool
}

var initQuestComponentsTests = []InitQuestComponentsTestCase{
	{
		name:        "Successful Quest Component Initialization",
		mockRepoErr: nil,
		mockGameData: &game.GameData{
			PartyLevel: 5,
			PartySize:  3,
		},
		expectedMsg: "should not return an error",
		wantErr:     false,
	},
	{
		name:        "Error Initializing Quest Repo",
		mockRepoErr: errors.New("error initializing quest repo"),
		mockGameData: &game.GameData{
			PartyLevel: 5,
			PartySize:  3,
		},
		wantErr:     true,
		expectedMsg: "should return a quest repo initialization error",
	},
}

func TestInitQuestComponents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range initQuestComponentsTests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorageManager(ctrl)
			mockStorage.EXPECT().LoadRepo("quests", gomock.Any()).Return(tt.mockRepoErr)

			coreInstance := NewCore(nil, mockStorage, tt.mockGameData)

			err := coreInstance.InitQuestComponents()

			// Assert

			if tt.wantErr {
				assert.NotEmpty(t, err, tt.expectedMsg)
			} else {
				assert.Empty(t, err, tt.expectedMsg)
			}

		})
	}
}
