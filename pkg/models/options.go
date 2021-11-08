package models

type Options struct {
	Pipeline       bool
	URL            string
	URLsFile       string
	RawFile        string
	Format         string
	Output         string
	NoCacheBusting bool
}
