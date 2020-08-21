package person

//go:generate mockgen -destination=../mock/male_mock.go -package=mock awesomeProject1/GoMockDemo/person Male

type Male interface {
	Get(id int64) error
}