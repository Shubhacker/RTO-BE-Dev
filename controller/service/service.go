package service

import (
	"context"
	"log"
	"net/http"
	"sync"

	database "github.com/Shubhacker/RTO-BE-Dev/Adapters/Database"
	"github.com/Shubhacker/RTO-BE-Dev/Adapters/constants"
	"github.com/Shubhacker/RTO-BE-Dev/Adapters/model"
	"github.com/Shubhacker/RTO-BE-Dev/Adapters/offense"
	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
	"github.com/google/uuid"
)

type Adapters struct {
}

func NewAdapter() *Adapters {
	return &Adapters{}
}

func (A Adapters) ReportSocialCommentsService(ctx context.Context, i *invoicer.ReportSocialCommentsRequest) (*invoicer.ReportSocialCommentsResponse, error) {
	var response invoicer.ReportSocialCommentsResponse

	Done, err := database.AddCommentOnSocial(i)
	if err != nil {
		log.Println(err.Error())
		response.Done = false
		return &response, err
	}

	response.Done = Done

	return &response, nil
}

func (A Adapters) ReactSocialPost(ctx context.Context, i *invoicer.ReactSocialRequest) (*invoicer.ReactSocialResponce, error) {
	var response invoicer.ReactSocialResponce

	// Check if data present in table (While creating data if gets error for unique column just call increment function)
	isRequestPresent, err := database.IsDataPresentInSocial(i.RequestId)
	if err != nil {
		if err.Error() == "no rows in result set" {
			go func(ipt *invoicer.ReactSocialRequest) {
				Done, err2 := database.CreateSocialData(ipt)
				if err2 != nil {
					log.Println(err2.Error())
				}
				response.Done = Done
			}(i)
		}
		// errCh <- err
		return &response, err
	}

	if isRequestPresent {
		Done, err := database.UpdateSocialData(i)
		if err != nil {
			return &response, err
		}
		response.Done = Done
	}

	return &response, nil
}

func (A Adapters) Social(ctx context.Context, i *invoicer.SocialRequest) (*invoicer.SocialData, error) {
	var response invoicer.SocialData
	var responseArray []*invoicer.SocialDataArray
	var RandomData []model.SocialData
	var RequestSlice []string
	var err error
	var likeMap map[string]model.ReportSocialData
	var wg sync.WaitGroup
	var mut sync.Mutex

	wg.Add(1)
	go func(requestBy string, wg *sync.WaitGroup) {
		RandomData, RequestSlice, err = database.FetchSocialFromDB(requestBy)
		if err != nil {
			log.Println(err.Error())
		}
		defer wg.Done()
	}(i.RequestBy, &wg)
	wg.Wait()

	wg.Add(1)
	go func(Slice []string, wg *sync.WaitGroup) {
		likeMap, err = database.FetchLikesDataFromDB(RequestSlice)
		defer wg.Done()
	}(RequestSlice, &wg)
	wg.Wait()

	for _, data := range RandomData {
		wg.Add(1)
		go func(like map[string]model.ReportSocialData, DBdata model.SocialData, wg *sync.WaitGroup, mut *sync.Mutex) {
			mut.Lock()
			var SingleData invoicer.SocialDataArray
			likeStruct := like[DBdata.RequestId]
			SingleData.ImageUrls = DBdata.ImageUrls
			SingleData.Location = DBdata.Location
			SingleData.Offense = DBdata.Offense
			SingleData.RTO_Approved = DBdata.RTOApproved
			SingleData.RequestId = DBdata.RequestId
			SingleData.SubnittedBy_RTO = DBdata.SubmittedByRto
			SingleData.TotalFine = DBdata.TotalFine
			SingleData.Likes = likeStruct.Likes
			SingleData.DislikeReport = likeStruct.DisasgreeReport
			SingleData.ValidReport = likeStruct.ValidReport
			SingleData.Dislikes = likeStruct.DisLikes
			responseArray = append(responseArray, &SingleData)
			mut.Unlock()
			defer wg.Done()
		}(likeMap, data, &wg, &mut)
		wg.Wait()
	}

	response.SocialData = responseArray

	return &response, nil
}

func (A Adapters) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponseMessage, error) {
	log.Println("Inside service")
	database.FetchInfoFromDB()
	var result *invoicer.CreateResponseMessage
	return result, nil
}

func (A Adapters) CreateReport(ctx context.Context, i *invoicer.CreateReportRequest) (*invoicer.CreateReportResponse, error) {
	var result invoicer.CreateReportResponse
	var DBModel model.CreateRequestDBModel
	DBModel.ByRTO = i.ByRto
	DBModel.Comment = i.Comment
	DBModel.ImageUrl = i.ImageUrls
	DBModel.Location = i.Location
	DBModel.OffenseList = i.OffenseList
	// Image processing if offense correct with image proof
	// Python script for image processing with ANPR with OpenCV
	DBModel.RTOApproved = false
	newId := uuid.New().String()
	// work on UUID pending
	DBModel.ReportId = newId
	DBModel.ReportedBy = i.ReportedBy
	DBModel.Social = true
	// calculating fine
	allOffenseMap := offense.GetFineWithOffense(DBModel.OffenseList)
	var totalFine int
	for _, fine := range allOffenseMap {
		totalFine += fine
	}
	// Check if offense included with bribe
	if allOffenseMap[constants.OfficerBribe] != 0 {
		totalFine = totalFine * 2
	}
	DBModel.TotalFine = int64(totalFine)
	DBModel.VehicleNumber = i.VehicleNumber
	log.Println("Inside service for Create Report API")
	go func(DBModel model.CreateRequestDBModel) {
		_, err := database.AddReportToDatabase(DBModel)
		if err != nil {
			result.Message = err.Error()
			result.MessageCode = http.StatusBadGateway
			// Implement queue mechanism if error reported from DB for retry
		}
	}(DBModel)

	result.Message = constants.ReportSubmitted
	result.MessageCode = http.StatusCreated
	return &result, nil
}
