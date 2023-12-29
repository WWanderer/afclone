package form1

import (
	"time"

	c "github.com/WWanderer/afclone/models/common"
)

type TravelGrant struct {
	DistanceBand                 c.Control[c.Ccm2]  `bson:"distance_band"`
	NumberOfEligibleParticipants c.Control[int]     `bson:"number_of_eligible_participants"`
	UnitCost                     c.Control[float64] `bson:"unit_cost"`
	Grant                        c.Control[float64] `bson:"grant"`
}

type TravelGrantDTO struct {
	DistanceBand                 int     `form:"distance_band"`
	NumberOfEligibleParticipants int     `form:"number_of_eligible_participants"`
	UnitCost                     float64 `form:"unit_cost"`
	Grant                        float64 `form:"grant"`
}

func (tgd *TravelGrantDTO) Map() (tg *TravelGrant) {
	tg = new(TravelGrant)
	tg.DistanceBand.Value = c.Ccm2(tgd.DistanceBand)
	tg.NumberOfEligibleParticipants.Value = tgd.DistanceBand
	tg.UnitCost.Value = tgd.UnitCost
	tg.Grant.Value = tgd.Grant
	return
}

type InclusionSupport struct {
	Grant c.Control[float64] `bson:"grant"`
}

type InclusionSupportDTO struct {
	Grant float64 `form:"grant"`
}

func (icd *InclusionSupportDTO) Map() (ic *InclusionSupport) {
	ic = new(InclusionSupport)
	ic.Grant.Value = icd.Grant
	return
}

type ExceptionalCost struct {
	Id            c.Control[int]     `bson:"id"`
	CostType      c.Control[c.Ccm2]  `bson:"cost_type"`
	Justification c.Control[string]  `bson:"justification"`
	CostSupported c.Control[float64] `bson:"cost_supported"`
	Ratio         c.Control[float64] `bson:"ration"`
	Grant         c.Control[float64] `bson:"grant"`
}

type ExceptionalCostDTO struct {
	Id            int     `form:"id"`
	CostType      int     `form:"cost_type"`
	Justification string  `form:"justification"`
	CostSupported float64 `form:"cost_supported"`
	Ratio         float64 `form:"ratio"`
	Grant         float64 `form:"grant"`
}

func (ecd *ExceptionalCostDTO) Map() (ec *ExceptionalCost) {
	ec = new(ExceptionalCost)
	ec.Id.Value = ecd.Id
	ec.CostType.Value = c.Ccm2(ecd.CostType)
	ec.Justification.Value = ecd.Justification
	ec.CostSupported.Value = ecd.CostSupported
	ec.Ratio.Value = ecd.Ratio
	ec.Grant.Value = ecd.Grant
	return
}

type ExceptionalCosts struct {
	ExceptionalCostList c.Control[[]ExceptionalCost] `bson:"exceptional_cost_list"`
	Grant               c.Control[float64]           `bson:"grant"`
}

type ExceptionalCostsDTO struct {
	ExceptionalCostListDTO []ExceptionalCostDTO `form:"exceptional_cost_list_dto"`
	Grant                  float64              `form:"grant"`
}

func (ecd *ExceptionalCostsDTO) Map() (ec *ExceptionalCosts) {
	ec = new(ExceptionalCosts)
	for _, ecDTO := range ecd.ExceptionalCostListDTO {
		ec.ExceptionalCostList.Value = append(ec.ExceptionalCostList.Value, *ecDTO.Map())
	}
	ec.Grant.Value = ecd.Grant
	return
}

type OrganisationalSupport struct {
	UnitCost c.Control[float64] `bson:"unit_cost"`
	Grant    c.Control[float64] `bson:"grant"`
}

type OrganisationalSupportDTO struct {
	UnitCost float64 `form:"unit_cost"`
	Grant    float64 `form:"grant"`
}

func (osd *OrganisationalSupportDTO) Map() (os *OrganisationalSupport) {
	os = new(OrganisationalSupport)
	os.UnitCost.Value = osd.UnitCost
	os.Grant.Value = osd.Grant
	return
}

type Flow struct {
	Id                                     c.Control[int]                   `bson:"id"`
	CountryDestination                     c.Control[c.Ccm2]                `bson:"country_destination"`
	NumberOfParticipants                   c.Control[int]                   `bson:"number_of_participants"`
	NumberOfParticipantsFewerOpportunities c.Control[int]                   `bson:"number_of_participants_fewer_opportunities"`
	NumberOfAccompanyingPersons            c.Control[int]                   `bson:"number_of_accompanying_persons"`
	GreenTravel                            c.Control[bool]                  `bson:"green_travel"`
	TravelGrant                            c.Control[TravelGrant]           `bson:"travel_grant"`
	InclusionSupport                       c.Control[InclusionSupport]      `bson:"inclusion_support"`
	ExceptionalCosts                       c.Control[ExceptionalCosts]      `bson:"exceptional_costs"`
	OrganisationalSupport                  c.Control[OrganisationalSupport] `bson:"organisational_support"`
	FlowGrant                              c.Control[float64]               `bson:"flow_grant"`
}

type FlowDTO struct {
	Id                                     int                      `form:"id"`
	CountryDestination                     int                      `form:"country_destination"`
	NumberOfParticipants                   int                      `form:"number_of_participants"`
	NumberOfParticipantsFewerOpportunities int                      `form:"number_of_participants_fewer_opportunities"`
	NumberOfAccompanyingPersons            int                      `form:"number_of_accompanying_persons"`
	GreenTravel                            bool                     `form:"green_travel"`
	TravelGrant                            TravelGrantDTO           `form:"travel_grant"`
	InclusionSupport                       InclusionSupportDTO      `form:"inclusion_support"`
	ExceptionalCosts                       ExceptionalCostsDTO      `form:"exceptional_costs"`
	OrganisationalSupport                  OrganisationalSupportDTO `form:"organisational_support"`
	FlowGrant                              float64                  `form:"flow_grant"`
}

func (fd *FlowDTO) Map() (f *Flow) {
	f = new(Flow)
	f.Id.Value = fd.Id
	f.CountryDestination.Value = c.Ccm2(fd.CountryDestination)
	f.NumberOfParticipants.Value = fd.NumberOfParticipants
	f.NumberOfParticipantsFewerOpportunities.Value = fd.NumberOfParticipantsFewerOpportunities
	f.NumberOfAccompanyingPersons.Value = fd.NumberOfAccompanyingPersons
	f.GreenTravel.Value = fd.GreenTravel
	f.TravelGrant.Value = *fd.TravelGrant.Map()
	f.InclusionSupport.Value = *fd.InclusionSupport.Map()
	f.ExceptionalCosts.Value = *fd.ExceptionalCosts.Map()
	f.OrganisationalSupport.Value = *fd.OrganisationalSupport.Map()
	f.FlowGrant.Value = fd.FlowGrant
	return
}

type Activity struct {
	Id                                     c.Control[int]       `bson:"id"`
	Title                                  c.Control[string]    `bson:"title"`
	ActivityType                           c.Control[c.Ccm2]    `bson:"activity_type"`
	CountryStart                           c.Control[c.Ccm2]    `bson:"country_start"`
	Location                               c.Control[string]    `bson:"location"`
	StartDate                              c.Control[time.Time] `bson:"start_date"`
	EndDate                                c.Control[time.Time] `bson:"end_date"`
	NumberOfParticipants                   c.Control[int]       `bson:"number_of_participants"`
	NumberOfParticipantsFewerOpportunities c.Control[int]       `bson:"number_of_participants_fewer_opportunities"`
	NumberOfAccompanyingPersons            c.Control[int]       `bson:"number_of_accompanying_persons"`
	Description                            c.Control[string]    `bson:"description"`
	FlowList                               c.Control[[]Flow]    `bson:"flow_list"`
	ActivityGrant                          c.Control[float64]   `bson:"activity_grant"`
}

type ActivityDTO struct {
	Id                                     int       `form:"id"`
	Title                                  string    `form:"title"`
	ActivityType                           int       `form:"activity_type"`
	CountryStart                           int       `form:"country_start"`
	Location                               string    `form:"location"`
	StartDate                              string    `form:"start_date"`
	EndDate                                string    `form:"end_date"`
	NumberOfParticipants                   int       `form:"number_of_participants"`
	NumberOfParticipantsFewerOpportunities int       `form:"number_of_participants_fewer_opportunities"`
	NumberOfAccompanyingPersons            int       `form:"number_of_accompanying_persons"`
	Description                            string    `form:"description"`
	FlowList                               []FlowDTO `form:"flow_list"`
	ActivityGrant                          float64   `form:"activity_grant"`
}

func (ad *ActivityDTO) Map() (a *Activity) {
	a = new(Activity)
	a.Id.Value = ad.Id
	a.Title.Value = ad.Title
	a.ActivityType.Value = c.Ccm2(ad.ActivityType)
	a.CountryStart.Value = c.Ccm2(ad.CountryStart)
	a.Location.Value = ad.Location
	a.StartDate.Value = getDate(ad.StartDate)
	a.EndDate.Value = getDate(ad.EndDate)
	a.NumberOfParticipants.Value = ad.NumberOfParticipants
	a.NumberOfParticipantsFewerOpportunities.Value = ad.NumberOfParticipantsFewerOpportunities
	a.NumberOfAccompanyingPersons.Value = ad.NumberOfAccompanyingPersons
	a.Description.Value = ad.Description

	for _, fd := range ad.FlowList {
		a.FlowList.Value = append(a.FlowList.Value, *fd.Map())
	}

	a.ActivityGrant.Value = ad.ActivityGrant
	return
}

type Activities struct {
	PurposeActivities c.Control[string]     `bson:"purpose_activities"`
	ActivityList      c.Control[[]Activity] `bson:"activity_list"`
	TotalGrant        c.Control[float64]    `bson:"total_grant"`
}

type ActivitiesDTO struct {
	PurposeActivities string        `form:"purpose_activities"`
	ActivityList      []ActivityDTO `form:"activity_list"`
	TotalGrant        float64       `form:"total_grant"`
}

func (ad *ActivitiesDTO) Map() (a *Activities) {
	a = new(Activities)
	a.PurposeActivities.Value = ad.PurposeActivities

	for _, activitydto := range ad.ActivityList {
		a.ActivityList.Value = append(a.ActivityList.Value, *activitydto.Map())
	}

	a.TotalGrant.Value = ad.TotalGrant
	return
}
