package database

import (
	"context"
	"log"

	model "github.com/Shubhacker/RTO-BE-Dev/Adapters/model"
	"github.com/Shubhacker/RTO-BE-Dev/invoicer"
	"github.com/jmoiron/sqlx"
)

type data struct {
	id       string
	image    string
	location int64
	comment  string
}

func FetchInfoFromDB() {
	var DBdata data
	query := `select report_id, image_urls, "location", "comment" from report.public_report;`
	rows, err := DbConn.Db.Query(context.Background(), query)
	if err != nil {
		log.Println("Issue (FetchInfoFromDB) : ", err.Error())
	}
	for rows.Next() {
		rows.Scan(
			&DBdata.id,
			&DBdata.image,
			&DBdata.location,
			&DBdata.comment,
		)
	}

}

func AddReportToDatabase(DBModel model.CreateRequestDBModel) (string, error) {
	var reportId string
	var inputArgs []interface{}
	sqlQuery := `
	INSERT INTO report.public_report (report_id,image_urls,"location",offense,is_submitted_by_rto,social,total_fine,rto_approved,"comment",reported_by,submitted_at,is_active,deactive_report,vehicle_number) 
	VALUES
	(?,`
	// Image ,

	inputArgs = append(inputArgs, DBModel.ReportId)
	// sqlQuery += `'[`
	// for key, image := range DBModel.ImageUrl {
	// 	var question string
	// 	log.Println(key, "<- Image")
	// 	question = `, ?`
	// 	if key == 0 {
	// 		question = `? `
	// 	}

	// 	sqlQuery += question
	// 	log.Println(image)
	// 	inputArgs = append(inputArgs, image)
	// }
	sqlQuery += `?, ? ,`
	inputArgs = append(inputArgs, DBModel.ImageUrl)
	inputArgs = append(inputArgs, DBModel.Location)
	// Offense ,
	// sqlQuery += `'[`
	// for key, offense := range DBModel.OffenseList {
	// 	var question string
	// 	log.Println(key, "<- offense")
	// 	question = `, ?`
	// 	if key == 0 {
	// 		question = `? `
	// 	}

	// 	sqlQuery += question
	// 	inputArgs = append(inputArgs, offense)
	// }
	sqlQuery += `?, ? ,? ,? ,? ,? ,? ,current_timestamp ,true ,false ,? ) returning report_id ;`

	inputArgs = append(inputArgs, DBModel.OffenseList)
	inputArgs = append(inputArgs, DBModel.ByRTO)
	inputArgs = append(inputArgs, DBModel.Social)
	inputArgs = append(inputArgs, DBModel.TotalFine)
	inputArgs = append(inputArgs, DBModel.RTOApproved)
	inputArgs = append(inputArgs, DBModel.Comment)
	inputArgs = append(inputArgs, DBModel.ReportedBy)
	inputArgs = append(inputArgs, DBModel.VehicleNumber)

	sqlQuery = sqlx.Rebind(sqlx.DOLLAR, sqlQuery)

	err := DbConn.Db.QueryRow(context.Background(), sqlQuery, inputArgs...).Scan(&reportId)
	if err != nil {
		log.Println("Issue (AddReportToDatabase) :", err.Error())
		return "", err
	}

	return "Reported TO RTO : " + reportId, nil
}

func FetchSocialFromDB(requestBy string) ([]model.SocialData, []string, error) {
	var response []model.SocialData
	var requestIdSlice []string
	query := `select report_id ,image_urls ,"location" ,offense ,is_submitted_by_rto ,total_fine ,rto_approved from report.public_report order by RANDOM() limit 30;`

	rows, err := DbConn.Db.Query(context.Background(), query)
	if err != nil {
		log.Println("Issue (FetchSocialFromDB) : ", err.Error())
		return response, requestIdSlice, err
	}
	for rows.Next() {
		var DBdata model.SocialData
		rows.Scan(
			&DBdata.RequestId,
			&DBdata.ImageUrls,
			&DBdata.Location,
			&DBdata.Offense,
			&DBdata.SubmittedByRto,
			&DBdata.TotalFine,
			&DBdata.RTOApproved,
		)
		response = append(response, DBdata)
		requestIdSlice = append(requestIdSlice, DBdata.RequestId)
	}

	return response, requestIdSlice, nil
}

func FetchLikesDataFromDB(reportIds []string) (map[string]model.ReportSocialData, error) {
	likesMap := make(map[string]model.ReportSocialData)
	var inputArgs []interface{}
	query := `select report_id , total_likes , total_dislike , valid_report , disagree from report.social_data sd where report_id in (`

	for key, reportId := range reportIds {
		var extendQuery string
		extendQuery += `, ?`
		if key < 1 {
			extendQuery = `?`
		}
		query += extendQuery
		inputArgs = append(inputArgs, reportId)
	}

	query += `)`
	query = sqlx.Rebind(sqlx.DOLLAR, query)
	rows, err := DbConn.Db.Query(context.Background(), query, inputArgs...)
	if err != nil {
		log.Println("Issue (FetchLikesDataFromDB) : ", err.Error())
		return likesMap, err
	}
	for rows.Next() {
		var DBdata model.ReportSocialData
		rows.Scan(
			&DBdata.RequestId,
			&DBdata.Likes,
			&DBdata.DisLikes,
			&DBdata.ValidReport,
			&DBdata.DisasgreeReport,
		)
		likesMap[DBdata.RequestId] = DBdata
	}

	return likesMap, nil
}

func IsDataPresentInSocial(requestId string) (bool, error) {
	var IsRequestPresent int

	query := `select 1 from report.social_data where report_id = $1`

	err := DbConn.Db.QueryRow(context.Background(), query, requestId).Scan(&IsRequestPresent)
	if err != nil {
		log.Println("Issue (IsDataPresentInSocial) : ", err.Error())
		return false, err
	}
	if IsRequestPresent == 0 {
		return false, nil
	}

	return true, nil
}

func UpdateSocialData(i *invoicer.ReactSocialRequest) (bool, error) {
	var multipleUpdate bool
	query := `update report.social_data set `

	if i.Like {
		if multipleUpdate {
			query += `, `
		}
		query += `total_likes = total_likes +1 `
		multipleUpdate = true
	}

	if i.DisLike {
		if multipleUpdate {
			query += `, `
		}
		query += `total_dislike = total_dislike +1 `
		multipleUpdate = true
	}

	if i.DisAgree {
		if multipleUpdate {
			query += `, `
		}
		query += `disagree = disagree +1 `
		multipleUpdate = true
	}

	if i.ValidReport {
		if multipleUpdate {
			query += `, `
		}
		query += `valid_report = valid_report +1 `
		multipleUpdate = true
	}
	query += ` where report_id = $1`

	_ = DbConn.Db.QueryRow(context.Background(), query, i.RequestId)

	return true, nil
}

func CreateSocialData(i *invoicer.ReactSocialRequest) (bool, error) {
	var inputArgs []interface{}
	var reportId string
	value := 0
	query := `insert into report.social_data (report_id, total_likes, total_dislike, valid_report, disagree) 
	values ($1, $2, $3, $4, $5) returning report_id`
	inputArgs = append(inputArgs, i.RequestId)
	if i.Like {
		value = 1
	}
	inputArgs = append(inputArgs, value)
	value = 0

	if i.DisLike {
		value = 1
	}
	inputArgs = append(inputArgs, value)
	value = 0

	if i.DisAgree {
		value = 1
	}
	inputArgs = append(inputArgs, value)
	value = 0

	if i.ValidReport {
		value = 1
	}
	inputArgs = append(inputArgs, value)
	value = 0
	err := DbConn.Db.QueryRow(context.Background(), query, inputArgs...).Scan(&reportId)
	if err != nil {
		log.Println("Issue (CreateSocialData) : ", err.Error())
		return false, err
	}

	return true, nil
}

func AddCommentOnSocial(i *invoicer.ReportSocialCommentsRequest) (bool, error) {
	var inputArgs []interface{}
	var reportId string
	query := `INSERT INTO report.report_comments (report_id,"comment",created_at,comment_by,comment_likes,comment_dislike) VALUES
	($1,$2,current_timestamp,$3,$4,$5) returning report_id`
	inputArgs = append(inputArgs, i.RequestId)
	inputArgs = append(inputArgs, i.Comment)
	inputArgs = append(inputArgs, i.CommentBy)
	inputArgs = append(inputArgs, 0)
	inputArgs = append(inputArgs, 0)

	err := DbConn.Db.QueryRow(context.Background(), query, inputArgs...).Scan(&reportId)
	if err != nil {
		log.Println("Issue (AddCommentOnSocial) : ", err.Error())
		return false, err
	}

	return true, nil
}
