package form1

import (
	"time"

	c "github.com/WWanderer/afclone/models/common"
	"github.com/gofiber/fiber/v2/log"
)

type Context struct {
	Title          c.Control[string]     `json:"title,omitempty" bson:"title,omitempty" form:"title"`
	TitleEnglish   c.Control[string]     `json:"title_english,omitempty" bson:"title_english,omitempty" form:"title_english"`
	StartDate      c.Control[time.Time]  `json:"start_date,omitempty" bson:"start_date,omitempty" form:"start_date"`
	Duration       c.Control[int]        `json:"duration,omitempty" bson:"duration,omitempty" form:"duration"`
	EndDate        c.Control[time.Time]  `json:"end_date,omitempty" bson:"end_date,omitempty" form:"end_date"`
	Language       c.Control[c.Ccm2]     `json:"language,omitempty" bson:"language,omitempty" form:"language"`
	NationalAgency c.Control[c.Ccm2]     `json:"national_agency,omitempty" bson:"national_agency,omitempty" form:"national_agency"`
}

type ContextDTO struct {
	Title          string `form:"title"`
	TitleEnglish   string `form:"title_english"`
	StartDate      string `form:"start_date"`
	Duration       int    `form:"duration"`
	EndDate        string `form:"end_date"`
	Language       int    `form:"language"`
	NationalAgency int    `form:"national_agency"`
}

func (dto *ContextDTO) Map() *Context {
	context := new(Context)

	context.Title.Value = dto.Title
	context.TitleEnglish.Value = dto.TitleEnglish
	context.StartDate.Value = getDate(dto.StartDate)
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
	dto.StartDate = context.StartDate.Value.Format("2006-01-02")
	dto.Duration = context.Duration.Value
	dto.EndDate = context.EndDate.Value.Format("2006-01-02")
	dto.Language = int(context.Language.Value)
	dto.NationalAgency = int(context.NationalAgency.Value)

	return dto
}

func getDate(timeString string) (time.Time) {
	theTime, err := time.Parse("2006-01-02", timeString)
	if err != nil {
		log.Error("could not parse time: ", err)
		return time.Time{}
	}
	return theTime
}
