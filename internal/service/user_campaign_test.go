package service

import (
	"go_notifier/internal/db/repository"
	"go_notifier/internal/dto"
	"go_notifier/internal/service/mocks"
	"go_notifier/pkg/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserCampaign(t *testing.T) {
	db, mock := database.InitMockDB(t)
	defer db.Close()

	dto := &dto.CampaignUser{
		CampaignUUID: "uuid1",
		UserUUID:     "uuid2",
	}

	mockCampaignRepo := mocks.NewCampaignRepository(t)
	mockUserRepo := mocks.NewUserRepository(t)
	SetCampaignRepository(mockCampaignRepo)
	SetUserRepository(mockUserRepo)

	successFunc := func(userTimezone, campaignTime, calculatedTimeForUser string) func(t *testing.T) {
		return func(t *testing.T) {
			var expectedUserId int = 5
			var expectedCampaignId int = 7

			userResp := &repository.UserIdTimezone{
				ID:       expectedUserId,
				Timezone: userTimezone,
			}
			mockUserRepo.On("GetUserIDAndTimezoneByUUID", dto.UserUUID).Once().Return(userResp, nil)

			campaignResp := &repository.CampaignIdTime{
				ID:   expectedCampaignId,
				Time: campaignTime,
			}
			mockCampaignRepo.On("GetCampgignIdAndTimeByUUID", dto.CampaignUUID).Once().Return(campaignResp, nil)

			mock.ExpectPrepare("INSERT INTO user_campaign").
				ExpectExec().
				WithArgs(campaignResp.ID, userResp.ID, calculatedTimeForUser).
				WillReturnResult(sqlmock.NewResult(1, 1))

			_, err := CreateUserCampaign(dto)
			assert.Nil(t, err)
		}
	}

	successCases := map[string]struct {
		userTimezone     string
		campaignTime     string
		expectedUserTime string
	}{
		"success with UTC timezone":        {userTimezone: "UTC", campaignTime: "10:11:00", expectedUserTime: "10:11:00"},
		"success with UTC timezone case 2": {userTimezone: "UTC", campaignTime: "15:30:40", expectedUserTime: "15:30:40"},
		"success with Kyiv timezone":       {userTimezone: "Europe/Kyiv", campaignTime: "10:30:00", expectedUserTime: "08:30:00"},
	}
	for name, testCase := range successCases {
		t.Run(name, successFunc(testCase.userTimezone, testCase.campaignTime, testCase.expectedUserTime))
	}
}
