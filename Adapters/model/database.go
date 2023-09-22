package model

type SocialData struct {
	RequestId      string
	ImageUrls      []string
	Location       string
	Offense        []string
	SubmittedByRto bool
	TotalFine      int64
	RTOApproved    bool
}

type ReportSocialData struct {
	RequestId       string
	Likes           int64
	DisLikes        int64
	ValidReport     int64
	DisasgreeReport int64
}

type ReportComments struct {
	Comment        string
	CommentAt      string
	CommentBy      string
	CommentLikes   int64
	CommentDislike int64
}

type FinalSocialData struct {
	RequestId          string
	ReportData         SocialData
	UsersFeedbackData  ReportSocialData
	ReportCommentsData []ReportComments
}
