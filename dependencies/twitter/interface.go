package twitter

type TwitterClient interface {
	Search(keyword string, maxID int64) ([]Tweet, error) 
}

type Tweet struct {
	ID int64
	Message string
}