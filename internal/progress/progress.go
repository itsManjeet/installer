package progress

type ProgressBar interface {
	Update(int, string)
}
