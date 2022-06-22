package valuate

type Create struct {
	AppID      string  `json:"app_id"`
	ProfitData string  `json:"profit_data"`
	Choices    Choices `json:"choices"`
	State      uint8   `json:"state"`
}
type Choices struct {
	CompetitiveLandscape         []int `json:"competitive_landscape"`
	BarrierIndustry              []int `json:"barrier_industry"`
	IndustryHook                 []int `json:"industry_hook"`
	Threat                       []int `json:"threat"`
	PolicySupportPower           []int `json:"policy_support_power"`
	PolicySupportType            []int `json:"policy_support_type"`
	Founder                      []int `json:"founder"`
	ManagementExp                []int `json:"management_exp"`
	AverageWorkingYears          []int `json:"average_working_years"`
	UndergraduatesRadio          []int `json:"undergraduates_radio"`
	CertificateRadio             []int `json:"certificate_radio"`
	IsTraining                   []int `json:"is_training"`
	BusinessArea                 []int `json:"business_area"`
	MarketingMethod              []int `json:"marketing_method"`
	BusinessProcessing           []int `json:"business_processing"`
	Loyalty                      []int `json:"loyalty"`
	IntellectualPropertyNumbers  []int `json:"intellectual_property_numbers"`
	RdInvestmentRadio            []int `json:"rd_investment_radio"`
	PrdType                      []int `json:"prd_type"`
	InnovationLevel              []int `json:"innovation_level"`
	CoreTurnoverRadio            []int `json:"core_turnover_radio"`
	IsInternalControlSystem      []int `json:"is_internal_control_system"`
	IsDishonestyRecord           []int `json:"is_dishonesty_record"`
	IsRegularPhysicalExamination []int `json:"is_regular_physical_examination"`
	ExitStrategy                 []int `json:"exit_strategy"`
}
type Search struct {
	AppID    string `form:"app_id"`    /*企业ID*/
	Page     int    `form:"page"`      /*页*/
	PageSize int    `form:"page_size"` /*页大小*/
}
