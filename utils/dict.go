package utils

// map
var filterMap = []string{
	",district",
	",register_capital",
	",estimate_value",
	",industry",
	",app_id",
	",name",
	",registration_number",
}
var sortMap = []string{
	",register_capital",
	",estimate_value",
	",name",
}

func ParseFilter(k int) string {
	return filterMap[k]
}

func ParseSortColumn(k int) string {
	return sortMap[k]
}

//单选题字典
var ValKey2IndexChoice = []string{
	"competitive_landscape",
	"barrier_industry",
	"industry_hook",
	"threat",
	"policy_support_power",
	"policy_support_type",
	"founder",
	"management_exp",
	"average_working_years",
	"undergraduates_radio",
	"certificate_radio",
	"is_training",
	"business_area",
	"marketing_method",
	"business_processing",
	"loyalty",
	"intellectual_property_numbers",
	"rd_investment_radio",
	"prd_type",
	"innovation_level",
	"core_turnover_radio",
	"is_internal_control_system",
	"is_dishonesty_record",
	"is_regular_physical_examination",
	"exit_strategy",
}

//企业估值数据字典
var ValKey2IndexData = []string{
	"year",
	"money_funds",
	"bill_receivable",
	"accounts_receivable",
	"stock",
	"total_current_assets",
	"fixed_assets",
	"construction_in_progress",
	"total_un_current_assets",
	"total_current_liabilities",
	"total_un_current_liabilities",
	"total_revenue",
	"sales",
	"first_sales_radio",
	"second_sales_radio",
	"third_sales_radio",
	"total_profit",
	"operating_profit",
	"gross_profit_margin",
	"total_cost",
	"operating_cost",
	"financial_expenses",
	"management_expense",
	"sales_expense",
	"research_and_development_expenses",
	"taxes_and_surcharges",
	"free_cash",
	"assets_and_liabilities",
	"current_ratio",
	"quick_ratio",
	"net_cash_flow_from_operating_activities",
	"total_liabilities",
	"total_assets",
	"short_term_loan",
	"liabilities_maturity_one_year",
	"long_term_loan",
	"accounts_payable",
	"bills_payable",
	"total_liabilities2",
	"net_profit",
	"securities",
}
