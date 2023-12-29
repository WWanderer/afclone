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

type InclusionSupport struct {
	Grant c.Control[float64] `bson:"grant"`
}

type ExceptionalCost struct {
	Id            c.Control[int]     `bson:"id"`
	CostType      c.Control[c.Ccm2]  `bson:"cost_type"`
	Justification c.Control[string]  `bson:"justification"`
	CostSupported c.Control[float64] `bson:"cost_supported"`
	Ration        c.Control[float64] `bson:"ration"`
	Grant         c.Control[float64] `bson:"grant"`
}

type ExceptionalCosts struct {
	ExceptionalCostList c.Control[[]ExceptionalCost] `bson:"exceptional_cost_list"`
	Grant               c.Control[float64]           `bson:"grant"`
}

type OrganisationalSupport struct {
	UnitCost c.Control[float64] `bson:"unit_cost"`
	Grant    c.Control[float64] `bson:"grant"`
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
	FlowList                               c.Control[Flow]      `bson:"flow_list"`
	ActivityGrant                          c.Control[float64]   `bson:"activity_grant"`
}

type Activities struct {
	PurposeActivities c.Control[string]      `bson:"purpose_activities"`
	ActivityList      c.Control[[]Activity]  `bson:"activity_list"`
	TotalGrant        c.Control[float64]     `bson:"total_grant"`
}

