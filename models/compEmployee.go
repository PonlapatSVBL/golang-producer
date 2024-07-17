package models

type EmployeeStruct struct {
	EmployeeId                    string `json:"employee_id" db:"employee_id"`
	ApplicationId                 string `json:"application_id" db:"application_id"`
	EmployeeCode                  string `json:"employee_code" db:"employee_code"`
	FingCode                      string `json:"fing_code" db:"fing_code"`
	EmployeeTypeCode              string `json:"employee_type_code" db:"employee_type_code"`
	EmployeeTypeGroupId           string `json:"employee_type_group_id" db:"employee_type_group_id"`
	EmployeeNickname              string `json:"employee_nickname" db:"employee_nickname"`
	EmployeeNicknameEn            string `json:"employee_nickname_en" db:"employee_nickname_en"`
	EmployeeName                  string `json:"employee_name" db:"employee_name"`
	EmployeeLastName              string `json:"employee_last_name" db:"employee_last_name"`
	EmployeeNameEn                string `json:"employee_name_en" db:"employee_name_en"`
	EmployeeLastNameEn            string `json:"employee_last_name_en" db:"employee_last_name_en"`
	EmployeeTitleLv               string `json:"employee_title_lv" db:"employee_title_lv"`
	EmployeeGender                string `json:"employee_gender" db:"employee_gender"`
	EmployeeForeigner             string `json:"employee_foreigner" db:"employee_foreigner"`
	CountryId                     string `json:"country_id" db:"country_id"`
	EmployeeStatus                string `json:"employee_status" db:"employee_status"`
	PositionId                    string `json:"position_id" db:"position_id"`
	CompanyId                     string `json:"company_id" db:"company_id"`
	BranchId                      string `json:"branch_id" db:"branch_id"`
	DepartmentId                  string `json:"department_id" db:"department_id"`
	DivisionId                    string `json:"division_id" db:"division_id"`
	SectionId                     string `json:"section_id" db:"section_id"`
	Mobilephone                   string `json:"mobilephone" db:"mobilephone"`
	Emailaddress                  string `json:"emailaddress" db:"emailaddress"`
	Salary                        string `json:"salary" db:"salary"`
	SalaryLaw                     string `json:"salary_law" db:"salary_law"`
	SalaryPerWeek                 string `json:"salary_per_week" db:"salary_per_week"`
	SalaryPerWeekTypeLv           string `json:"salary_per_week_type_lv" db:"salary_per_week_type_lv"`
	PaymentMethod                 string `json:"payment_method" db:"payment_method"`
	SocialInsuranceMethodLv       string `json:"social_insurance_method_lv" db:"social_insurance_method_lv"`
	SocialInsuranceMethodConstant string `json:"social_insurance_method_constant" db:"social_insurance_method_constant"`
	SocialInsuranceDeductLv       string `json:"social_insurance_deduct_lv" db:"social_insurance_deduct_lv"`
	TaxMethodLv                   string `json:"tax_method_lv" db:"tax_method_lv"`
	TaxMethodConstant             string `json:"tax_method_constant" db:"tax_method_constant"`
	TaxMethodRate                 string `json:"tax_method_rate" db:"tax_method_rate"`
	TaxDeductLv                   string `json:"tax_deduct_lv" db:"tax_deduct_lv"`
	TaxStartMonth                 string `json:"tax_start_month" db:"tax_start_month"`
	DaysPerMonth                  string `json:"days_per_month" db:"days_per_month"`
	HoursPerDay                   string `json:"hours_per_day" db:"hours_per_day"`
	BirthDt                       string `json:"birth_dt" db:"birth_dt"`
	IdNo                          string `json:"id_no" db:"id_no"`
	SsoNo                         string `json:"sso_no" db:"sso_no"`
	OptCode                       string `json:"opt_code" db:"opt_code"`
	PersonId                      string `json:"person_id" db:"person_id"`
	LineUserId                    string `json:"line_user_id" db:"line_user_id"`
	PlayerId                      string `json:"player_id" db:"player_id"`
	AppleId                       string `json:"apple_id" db:"apple_id"`
	FacebookId                    string `json:"facebook_id" db:"facebook_id"`
	GoogleId                      string `json:"google_id" db:"google_id"`
	LineTokenId                   string `json:"line_token_id" db:"line_token_id"`
	LineTokenTodolistId           string `json:"line_token_todolist_id" db:"line_token_todolist_id"`
	DeviceId                      string `json:"device_id" db:"device_id"`
	Photograph                    string `json:"photograph" db:"photograph"`
	CompanyPaymentAccountId       string `json:"company_payment_account_id" db:"company_payment_account_id"`
	BankId                        string `json:"bank_id" db:"bank_id"`
	BankBranchCode                string `json:"bank_branch_code" db:"bank_branch_code"`
	BankAccountCode               string `json:"bank_account_code" db:"bank_account_code"`
	BankOtherName                 string `json:"bank_other_name" db:"bank_other_name"`
	WorkCycleIdJson               string `json:"work_cycle_id_json" db:"work_cycle_id_json"`
	WorkCycleFormat               string `json:"work_cycle_format" db:"work_cycle_format"`
	HolidayDayJson                string `json:"holiday_day_json" db:"holiday_day_json"`
	HolidayFormat                 string `json:"holiday_format" db:"holiday_format"`
	AuthFirst                     string `json:"auth_first" db:"auth_first"`
	AuthSecond                    string `json:"auth_second" db:"auth_second"`
	ClockInout                    string `json:"clock_inout" db:"clock_inout"`
	TrialRange                    string `json:"trial_range" db:"trial_range"`
	EffectiveDt                   string `json:"effective_dt" db:"effective_dt"`
	BeginDt                       string `json:"begin_dt" db:"begin_dt"`
	ContractExpireDt              string `json:"contract_expire_dt" db:"contract_expire_dt"`
	RetirementYear                string `json:"retirement_year" db:"retirement_year"`
	SignoutFlag                   string `json:"signout_flag" db:"signout_flag"`
	SignoutRequestDt              string `json:"signout_request_dt" db:"signout_request_dt"`
	SignoutDt                     string `json:"signout_dt" db:"signout_dt"`
	OutDt                         string `json:"out_dt" db:"out_dt"`
	SsoOutDt                      string `json:"sso_out_dt" db:"sso_out_dt"`
	SignoutTypeFlag               string `json:"signout_type_flag" db:"signout_type_flag"`
	SignoutRemark                 string `json:"signout_remark" db:"signout_remark"`
	RoundMonthConfig              string `json:"round_month_config" db:"round_month_config"`
	RoundXtraConfig               string `json:"round_xtra_config" db:"round_xtra_config"`
	RoundOtConfig                 string `json:"round_ot_config" db:"round_ot_config"`
	RoundWorktimeConfig           string `json:"round_worktime_config" db:"round_worktime_config"`
	HolidayApplyConfig            string `json:"holiday_apply_config" db:"holiday_apply_config"`
	ImportLogId                   string `json:"import_log_id" db:"import_log_id"`
	PersonalConfig                string `json:"personal_config" db:"personal_config"`
	Address                       string `json:"address" db:"address"`
	Address1                      string `json:"address1" db:"address1"`
	Address2                      string `json:"address2" db:"address2"`
	Address3                      string `json:"address3" db:"address3"`
	Address4                      string `json:"address4" db:"address4"`
	Address5                      string `json:"address5" db:"address5"`
	Address6                      string `json:"address6" db:"address6"`
	Address7                      string `json:"address7" db:"address7"`
	Address8                      string `json:"address8" db:"address8"`
	Address9                      string `json:"address9" db:"address9"`
	CountryCode                   string `json:"country_code" db:"country_code"`
	StateCode                     string `json:"state_code" db:"state_code"`
	DistrictCode                  string `json:"district_code" db:"district_code"`
	SubdistrictCode               string `json:"subdistrict_code" db:"subdistrict_code"`
	PostCode                      string `json:"post_code" db:"post_code"`
	CurrentAddress                string `json:"current_address" db:"current_address"`
	CurrentAddress1               string `json:"current_address1" db:"current_address1"`
	CurrentAddress2               string `json:"current_address2" db:"current_address2"`
	CurrentAddress3               string `json:"current_address3" db:"current_address3"`
	CurrentAddress4               string `json:"current_address4" db:"current_address4"`
	CurrentAddress5               string `json:"current_address5" db:"current_address5"`
	CurrentAddress6               string `json:"current_address6" db:"current_address6"`
	CurrentAddress7               string `json:"current_address7" db:"current_address7"`
	CurrentAddress8               string `json:"current_address8" db:"current_address8"`
	CurrentAddress9               string `json:"current_address9" db:"current_address9"`
	CurrentCountryCode            string `json:"current_country_code" db:"current_country_code"`
	CurrentStateCode              string `json:"current_state_code" db:"current_state_code"`
	CurrentDistrictCode           string `json:"current_district_code" db:"current_district_code"`
	CurrentSubdistrictCode        string `json:"current_subdistrict_code" db:"current_subdistrict_code"`
	CurrentPostCode               string `json:"current_post_code" db:"current_post_code"`
	HashtagDesc                   string `json:"hashtag_desc" db:"hashtag_desc"`
	ReferenceCode1                string `json:"reference_code_1" db:"reference_code_1"`
	ReferenceCode2                string `json:"reference_code_2" db:"reference_code_2"`
	ReferenceCode3                string `json:"reference_code_3" db:"reference_code_3"`
	ReferenceCode4                string `json:"reference_code_4" db:"reference_code_4"`
	ReferenceCode5                string `json:"reference_code_5" db:"reference_code_5"`
	SlipEncryption                string `json:"slip_encryption" db:"slip_encryption"`
	BlacklistFlag                 string `json:"blacklist_flag" db:"blacklist_flag"`
	MicrosoftEntraIdId            string `json:"microsoft_entra_id_id" db:"microsoft_entra_id_id"`
	CoaAccountGroupId             string `json:"coa_account_group_id" db:"coa_account_group_id"`
	OrderNo                       string `json:"order_no" db:"order_no"`
	ServerId                      string `json:"server_id" db:"server_id"`
	InstanceServerId              string `json:"instance_server_id" db:"instance_server_id"`
	InstanceServerChannelId       string `json:"instance_server_channel_id" db:"instance_server_channel_id"`
	PublishFlag                   string `json:"publish_flag" db:"publish_flag"`
	ApproveFlag                   string `json:"approve_flag" db:"approve_flag"`
	ApproveRemark                 string `json:"approve_remark" db:"approve_remark"`
	GlobalFlag                    string `json:"global_flag" db:"global_flag"`
	ReadOnlyFlag                  string `json:"read_only_flag" db:"read_only_flag"`
	SysDelFlag                    string `json:"sys_del_flag" db:"sys_del_flag"`
	RemoteIp                      string `json:"remote_ip" db:"remote_ip"`
	LanguageCode                  string `json:"language_code" db:"language_code"`
	CreatedBy                     string `json:"created_by" db:"created_by"`
	Created                       string `json:"created" db:"created"`
	LastUpdBy                     string `json:"last_upd_by" db:"last_upd_by"`
	LastUpd                       string `json:"last_upd" db:"last_upd"`
	SsoStartMonth                 string `json:"sso_start_month" db:"sso_start_month"`
}
