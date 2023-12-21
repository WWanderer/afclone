package form1

import (
	c "afclone/models/common"
	"fmt"
	"time"
)

type Context struct {
	Title          c.Control[string]
	TitleEnglish   c.Control[string]
	StartDate      c.Control[time.Time]
	Duration       c.Control[int]
	EndDate        c.Control[time.Time]
	Language       c.Control[c.Ccm2]
	NationalAgency c.Control[c.Ccm2]
}

type ContextDTO struct {
	Title          string `json:"title" form:"title"`
	TitleEnglish   string `json:"title_english" form:"title_english"`
	StartDate      string `json:"start_date" form:"start_date"`
	Duration       int    `json:"duration" form:"duration"`
	EndDate        string `json:"end_date" form:"end_date"`
	Language       int    `json:"language" form:"language"`
	NationalAgency int    `json:"national_agency" form:"national_agency"`
}

func (dto *ContextDTO) Map() *Context {
	context := new(Context)

	context.Title.Value = dto.Title
	context.TitleEnglish.Value = dto.TitleEnglish
	context.StartDate.Value = getDate(dto.StartDate)
	fmt.Println(dto.StartDate, context.StartDate.Value.Format("2006-01-02"))
	context.Duration.Value = dto.Duration
	context.EndDate.Value = context.StartDate.Value.AddDate(0, dto.Duration, 0)
	context.Language.Value = c.Ccm2(dto.Language)
	context.NationalAgency.Value = c.Ccm2(dto.NationalAgency)

	return context
}

func (context *Context) Map() *ContextDTO {
	dto := new(ContextDTO)

	dto.Title = context.Title.Value
	dto.TitleEnglish = context.TitleEnglish.Value
	dto.StartDate = context.StartDate.Value.Format("2002-01-02")
	dto.Duration = context.Duration.Value
	dto.EndDate = context.EndDate.Value.Format("2002-01-02")
	dto.Language = int(context.Language.Value)
	dto.NationalAgency = int(context.NationalAgency.Value)

	return dto
}

func getDate(timeString string) time.Time {
	theTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	return theTime
}
