package service

import (
	"fmt"
	"go_notifier/internal/db/repository"
	"go_notifier/internal/dto"
	"go_notifier/pkg/database"
	"time"
)

var userRepository repository.UserRepository
var campaignRepository repository.CampaignRepository

func SetCampaignRepository(repo repository.CampaignRepository) {
	campaignRepository = repo
}

func SetUserRepository(repo repository.UserRepository) {
	userRepository = repo
}

func init() {
	campaignRepository = &repository.CampaignRepositoryImpl{}
	userRepository = &repository.UserRepositoryImpl{}
}

func CreateUserCampaign(dto *dto.CampaignUser) (string, error) {
	user, err := userRepository.GetUserIDAndTimezoneByUUID(dto.UserUUID)
	if err != nil {
		return "", err
	}

	campaign, err := campaignRepository.GetCampgignIdAndTimeByUUID(dto.CampaignUUID)
	if err != nil {
		return "", err
	}

	userCampaignTime, err := calculateUserCampaignTime(campaign.Time, user.Timezone)

	insertStatement, err := database.DB.Mysql.Prepare("INSERT INTO user_campaign(campaign_id, user_id, time) VALUES (?,?, ?)")
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
func calculateUserCampaignTime(campaignTime, userTimezone string) (string, error) {
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
	userCampaignTime := userDateTimeUTC.Format("15:04:05")

	fmt.Println("Desired Time in User's Timezone:", userDateTime.Format("2006-01-02 15:04"))
	fmt.Println("Desired Time in UTC:", userDateTimeUTC.Format("2006-01-02 15:04 MST"))
	fmt.Println("Extracted Time from UTC:", userCampaignTime)

	return userCampaignTime, nil
}
