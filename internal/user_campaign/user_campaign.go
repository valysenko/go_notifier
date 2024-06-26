package user_campaign

import (
	"fmt"
	"go_notifier/internal/campaign"
	"go_notifier/internal/common"
	"go_notifier/internal/user"
	"go_notifier/pkg/database"
	"time"
)

// https://go.dev/wiki/CodeReviewComments#interfaces

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=UserRepository --case snake
type UserRepository interface {
	GetUserIDAndTimezoneByUUID(uuid string) (*user.UserIdTimezone, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.0 --name=CampaignRepository --case snake
type CampaignRepository interface {
	GetCampgignIdAndTimeByUUID(uuid string) (*campaign.CampaignIdTime, error)
}

type UserCampaignService struct {
	db                 *database.AppDB
	campaignRepository CampaignRepository
	userRepository     UserRepository
}

func NewUserCampaignService(db *database.AppDB, userRepo UserRepository, campaignRepo CampaignRepository) *UserCampaignService {
	return &UserCampaignService{
		db:                 db,
		campaignRepository: campaignRepo,
		userRepository:     userRepo,
	}
}

func (s *UserCampaignService) CreateUserCampaign(dto *common.CampaignUserRequest) (string, error) {
	user, err := s.userRepository.GetUserIDAndTimezoneByUUID(dto.UserUUID)
	if err != nil {
		return "", err
	}

	campaign, err := s.campaignRepository.GetCampgignIdAndTimeByUUID(dto.CampaignUUID)
	if err != nil {
		return "", err
	}

	userCampaignTime, err := s.calculateUserCampaignTime(campaign.Time, user.Timezone)

	insertStatement, err := s.db.Mysql.Prepare("INSERT INTO user_campaign(campaign_id, user_id, time) VALUES (?,?, ?)")
	if err != nil {
		return "", err
	}
	defer insertStatement.Close()

	res, err := insertStatement.Exec(campaign.ID, user.ID, userCampaignTime)
	if err != nil {
		return "", err
	}

	_, err = res.LastInsertId()
	return userCampaignTime, err
}

/*
* user campaign time is a Campaign time converted to user timezone and then saved in DB as UTC
* it's needed for a job which runs in UTC timezone, and "10:00" in different timezones is a different UTC time
* f.e for the campaign time 10:00:
* - for users from Ukraine, it should be 8:00 UTC time - notifications for those users will be sent in 8:00 UTC
* - for users from Germany, it should be 9:00 UTC time - notifications for those users will be sent in 9:00 UTC
 */
func (s *UserCampaignService) calculateUserCampaignTime(campaignTime, userTimezone string) (string, error) {
	userLocation, err := time.LoadLocation(userTimezone)
	if err != nil {
		return "", err
	}

	layout := "2006-01-02 15:04:05"
	desiredTimeWithDate := "1970-01-01 " + campaignTime
	desiredTimeObj, err := time.Parse(layout, desiredTimeWithDate)
	if err != nil {
		return "", err
	}

	currentTime := time.Now().In(userLocation)
	userDateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
		desiredTimeObj.Hour(), desiredTimeObj.Minute(), desiredTimeObj.Second(), 0, userLocation)

	userDateTimeUTC := userDateTime.UTC()
	userDateTimeUTC = userDateTimeUTC.Truncate(time.Minute) // it will set seconds=00
	userCampaignTime := userDateTimeUTC.Format("15:04:05")

	fmt.Println("Desired Time in User's Timezone:", userDateTime.Format("2006-01-02 15:04"))
	fmt.Println("Desired Time in UTC:", userDateTimeUTC.Format("2006-01-02 15:04 MST"))
	fmt.Println("Extracted Time from UTC:", userCampaignTime)

	return userCampaignTime, nil
}
