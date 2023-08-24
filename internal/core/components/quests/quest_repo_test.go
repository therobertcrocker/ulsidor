package quests

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces/mocks"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/changelog"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

type initTestCase struct {
	name           string
	mockStorageErr error
	wantErr        bool
	expectedMsg    string
}

var initTests = []initTestCase{
	{
		name:           "Successful Quest Repo Initialization",
		mockStorageErr: nil,
		wantErr:        false,
		expectedMsg:    "should not return an error",
	},
	{
		name:           "Error Initializing Quest Repo",
		mockStorageErr: errors.New("error initializing quest repo"),
		wantErr:        true,
		expectedMsg:    "should return a quest repo initialization error",
	},
}

func TestInit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range initTests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorageManager(ctrl)
			mockStorage.EXPECT().LoadRepo("quests", gomock.Any()).Return(tt.mockStorageErr)

			qr := NewQuestRepo(mockStorage)

			err := qr.Init()

			if tt.wantErr {
				assert.NotEmpty(t, err, tt.expectedMsg)
			} else {
				assert.Empty(t, err, tt.expectedMsg)
			}
		})
	}
}

type addNewQuestTestCase struct {
	name           string
	inputQuest     *types.Quest
	logEntry       changelog.LogEntry
	mockStorageErr error
	expectedMsg    string
}

var addNewQuestTests = []addNewQuestTestCase{
	{
		name:       "Successful Quest Addition",
		inputQuest: &types.Quest{
			// Some quest data
		},
		logEntry:       changelog.LogEntry{}, // Some log entry data
		mockStorageErr: nil,
		expectedMsg:    "should not return an error",
	},
	{
		name:       "Error Adding Quest with Existing ID",
		inputQuest: &types.Quest{
			// Some quest data with an ID that already exists
		},
		logEntry:       changelog.LogEntry{}, // Some log entry data
		mockStorageErr: nil,
		expectedMsg:    "should return a quest id already exists error",
	},
	// ... other test cases
}

func TestAddNewQuest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tt := range addNewQuestTests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorageManager(ctrl)
			mockStorage.EXPECT().SaveRepo("quests", gomock.Any()).Return(tt.mockStorageErr)

			qr := NewQuestRepo(mockStorage)
			err := qr.AddNewQuest(tt.inputQuest, tt.logEntry)
			if tt.mockStorageErr != nil {
				assert.NotEmpty(t, err, tt.expectedMsg)
			} else {
				assert.Empty(t, err, tt.expectedMsg)
			}
		})
	}
}
