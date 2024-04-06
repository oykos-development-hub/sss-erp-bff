package structs

import "time"

type FixedDeposit struct {
	ID                   int                    `json:"id"`
	OrganizationUnitID   int                    `json:"organization_unit_id"`
	Subject              string                 `json:"subject"`
	JudgeID              int                    `json:"judge_id"`
	CaseNumber           string                 `json:"case_number"`
	DateOfRecipiet       *time.Time             `json:"date_of_recipiet"`
	DateOfCase           *time.Time             `json:"date_of_case"`
	DateOfFinality       *time.Time             `json:"date_of_finality"`
	DateOfEnforceability *time.Time             `json:"date_of_enforceability"`
	DateOfEnd            *time.Time             `json:"date_of_end"`
	AccountID            int                    `json:"account_id"`
	FileID               int                    `json:"file_id"`
	Status               string                 `json:"status"`
	Type                 string                 `json:"type"`
	Items                []FixedDepositItem     `json:"items"`
	Dispatches           []FixedDepositDispatch `json:"dispatches"`
	Judges               []FixedDepositJudge    `json:"judges"`
	CreatedAt            time.Time              `json:"created_at"`
	UpdatedAt            time.Time              `json:"updated_at"`
}

type FixedDepositWill struct {
	ID                 int                        `json:"id"`
	OrganizationUnitID int                        `json:"organization_unit_id"`
	Subject            string                     `json:"subject"`
	FatherName         string                     `json:"father_name"`
	DateOfBirth        time.Time                  `json:"date_of_birth"`
	JMBG               string                     `json:"jmbg"`
	CaseNumberSI       string                     `json:"case_number_si"`
	CaseNumberRS       string                     `json:"case_number_rs"`
	DateOfReceiptSI    *time.Time                 `json:"date_of_receipt_si"`
	DateOfReceiptRS    *time.Time                 `json:"date_of_receipt_rs"`
	DateOfEnd          *time.Time                 `json:"date_of_end"`
	Status             string                     `json:"status"`
	FileID             int                        `json:"file_id"`
	Judges             []FixedDepositJudge        `json:"judges"`
	Dispatches         []FixedDepositWillDispatch `json:"dispatches"`
	CreatedAt          time.Time                  `json:"created_at"`
	UpdatedAt          time.Time                  `json:"updated_at"`
}

type FixedDepositWillDispatch struct {
	ID             int       `json:"id"`
	WillID         int       `json:"will_id"`
	DispatchTypeID int       `json:"dispatch_type_id"`
	JudgeID        int       `json:"judge_id"`
	CaseNumber     string    `json:"case_number"`
	DateOfDispatch time.Time `json:"date_of_dispatch"`
	FileID         int       `json:"file_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type FixedDepositItem struct {
	ID                 int        `json:"id"`
	DepositID          int        `json:"deposit_id"`
	CategoryID         int        `json:"category_id"`
	TypeID             int        `json:"type_id"`
	UnitID             int        `json:"unit_id"`
	CurrencyID         int        `json:"currency_id"`
	Amount             float32    `json:"amount"`
	SerialNumber       string     `json:"serial_number"`
	DateOfConfiscation *time.Time `json:"date_of_confiscation"`
	CaseNumber         string     `json:"case_number"`
	JudgeID            int        `json:"judge_id"`
	FileID             int        `json:"file_id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

type FixedDepositDispatch struct {
	ID           int        `json:"id"`
	DepositID    int        `json:"deposit_id"`
	CategoryID   int        `json:"category_id"`
	TypeID       int        `json:"type_id"`
	UnitID       int        `json:"unit_id"`
	CurrencyID   int        `json:"currency_id"`
	Amount       float32    `json:"amount"`
	SerialNumber string     `json:"serial_number"`
	DateOfAction *time.Time `json:"date_of_action"`
	Subject      string     `json:"subject"`
	ActionID     int        `json:"action_id"`
	CaseNumber   string     `json:"case_number"`
	JudgeID      int        `json:"judge_id"`
	FileID       int        `json:"file_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type FixedDepositJudge struct {
	ID          int        `json:"id"`
	JudgeID     int        `json:"judge_id"`
	DepositID   int        `json:"deposit_id"`
	WillID      int        `json:"will_id"`
	DateOfStart time.Time  `json:"date_of_start"`
	DateOfEnd   *time.Time `json:"date_of_end"`
	FileID      int        `json:"file_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
