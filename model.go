package model

import "time"

type LoanSearchRequest struct {
	FilterCriteria []struct {
		Field       string   `json:"field"`
		Operator    string   `json:"operator"`
		SecondValue string   `json:"secondValue,omitempty"`
		Value       string   `json:"valu,omitempty"`
		Values      []string `json:"values,omitempty"`
	} `json:"filterCriteria"`
	SortingCriteria struct {
		Field string `json:"field"`
		Order string `json:"order"`
	} `json:"sortingCriteria,omitempty"`
}

type LoanSearchResponse []struct {
	AccountArrearsSettings struct {
		DateCalculationMethod                     string  `json:"dateCalculationMethod"`
		EncodedKey                                string  `json:"encodedKey"`
		MonthlyToleranceDay                       int     `json:"monthlyToleranceDay"`
		NonWorkingDaysMethod                      string  `json:"nonWorkingDaysMethod"`
		ToleranceCalculationMethod                string  `json:"toleranceCalculationMethod"`
		ToleranceFloorAmount                      float64 `json:"toleranceFloorAmount"`
		TolerancePercentageOfOutstandingPrincipal float64 `json:"tolerancePercentageOfOutstandingPrincipal"`
		TolerancePeriod                           int     `json:"tolerancePeriod"`
	} `json:"accountArrearsSettings"`
	AccountHolderKey         string    `json:"accountHolderKey"`
	AccountHolderType        string    `json:"accountHolderType"`
	AccountState             string    `json:"accountState"`
	AccountSubState          string    `json:"accountSubState"`
	AccruedInterest          float64   `json:"accruedInterest"`
	AccruedPenalty           float64   `json:"accruedPenalty"`
	ActivationTransactionKey string    `json:"activationTransactionKey"`
	AllowOffset              bool      `json:"allowOffset"`
	ApprovedDate             time.Time `json:"approvedDate"`
	ArrearsTolerancePeriod   int       `json:"arrearsTolerancePeriod"`
	Assets                   []struct {
		Amount            float64 `json:"amount"`
		AssetName         string  `json:"assetName"`
		DepositAccountKey string  `json:"depositAccountKey"`
		EncodedKey        string  `json:"encodedKey"`
		GuarantorKey      string  `json:"guarantorKey"`
		GuarantorType     string  `json:"guarantorType"`
		OriginalAmount    float64 `json:"originalAmount"`
		OriginalCurrency  struct {
			Code         string `json:"code"`
			CurrencyCode string `json:"currencyCode"`
		} `json:"originalCurrency"`
	} `json:"assets"`
	AssignedBranchKey string `json:"assignedBranchKey"`
	AssignedCentreKey string `json:"assignedCentreKey"`
	AssignedUserKey   string `json:"assignedUserKey"`
	Balances          struct {
		FeesBalance                float64 `json:"feesBalance"`
		FeesDue                    float64 `json:"feesDue"`
		FeesPaid                   float64 `json:"feesPaid"`
		HoldBalance                float64 `json:"holdBalance"`
		InterestBalance            float64 `json:"interestBalance"`
		InterestDue                float64 `json:"interestDue"`
		InterestFromArrearsBalance float64 `json:"interestFromArrearsBalance"`
		InterestFromArrearsDue     float64 `json:"interestFromArrearsDue"`
		InterestFromArrearsPaid    float64 `json:"interestFromArrearsPaid"`
		InterestPaid               float64 `json:"interestPaid"`
		PenaltyBalance             float64 `json:"penaltyBalance"`
		PenaltyDue                 float64 `json:"penaltyDue"`
		PenaltyPaid                float64 `json:"penaltyPaid"`
		PrincipalBalance           float64 `json:"principalBalance"`
		PrincipalDue               float64 `json:"principalDue"`
		PrincipalPaid              float64 `json:"principalPaid"`
		RedrawBalance              float64 `json:"redrawBalance"`
	} `json:"balances"`
	ClosedDate           time.Time `json:"closedDate"`
	CreationDate         time.Time `json:"creationDate"`
	CreditArrangementKey string    `json:"creditArrangementKey"`
	Currency             struct {
		Code         string `json:"code"`
		CurrencyCode string `json:"currencyCode"`
	} `json:"currency"`
	DaysInArrears       int `json:"daysInArrears"`
	DaysLate            int `json:"daysLate"`
	DisbursementDetails struct {
		DisbursementDate         time.Time `json:"disbursementDate"`
		EncodedKey               string    `json:"encodedKey"`
		ExpectedDisbursementDate time.Time `json:"expectedDisbursementDate"`
		Fees                     []struct {
			Amount                  float64 `json:"amount"`
			EncodedKey              string  `json:"encodedKey"`
			Percentage              float64 `json:"percentage"`
			PredefinedFeeEncodedKey string  `json:"predefinedFeeEncodedKey"`
		} `json:"fees"`
		FirstRepaymentDate time.Time `json:"firstRepaymentDate"`
		TransactionDetails struct {
			EncodedKey              string `json:"encodedKey"`
			InternalTransfer        bool   `json:"internalTransfer"`
			TargetDepositAccountKey string `json:"targetDepositAccountKey"`
			TransactionChannelID    string `json:"transactionChannelId"`
			TransactionChannelKey   string `json:"transactionChannelKey"`
		} `json:"transactionDetails"`
	} `json:"disbursementDetails"`
	EncodedKey     string `json:"encodedKey"`
	FundingSources []struct {
		Amount             float64 `json:"amount"`
		AssetName          string  `json:"assetName"`
		DepositAccountKey  string  `json:"depositAccountKey"`
		EncodedKey         string  `json:"encodedKey"`
		GuarantorKey       string  `json:"guarantorKey"`
		GuarantorType      string  `json:"guarantorType"`
		ID                 string  `json:"id"`
		InterestCommission float64 `json:"interestCommission"`
		SharePercentage    float64 `json:"sharePercentage"`
	} `json:"fundingSources"`
	FuturePaymentsAcceptance string `json:"futurePaymentsAcceptance"`
	Guarantors               []struct {
		Amount            float64 `json:"amount"`
		AssetName         string  `json:"assetName"`
		DepositAccountKey string  `json:"depositAccountKey"`
		EncodedKey        string  `json:"encodedKey"`
		GuarantorKey      string  `json:"guarantorKey"`
		GuarantorType     string  `json:"guarantorType"`
	} `json:"guarantors"`
	ID                            string  `json:"id"`
	InterestAccruedInBillingCycle float64 `json:"interestAccruedInBillingCycle"`
	InterestCommission            float64 `json:"interestCommission"`
	InterestFromArrearsAccrued    float64 `json:"interestFromArrearsAccrued"`
	InterestSettings              struct {
		AccountInterestRateSettings []struct {
			EncodedKey               string    `json:"encodedKey"`
			IndexSourceKey           string    `json:"indexSourceKey"`
			InterestRate             float64   `json:"interestRate"`
			InterestRateCeilingValue float64   `json:"interestRateCeilingValue"`
			InterestRateFloorValue   float64   `json:"interestRateFloorValue"`
			InterestRateReviewCount  float64   `json:"interestRateReviewCount"`
			InterestRateReviewUnit   string    `json:"interestRateReviewUnit"`
			InterestRateSource       string    `json:"interestRateSource"`
			InterestSpread           float64   `json:"interestSpread"`
			ValidFrom                time.Time `json:"validFrom"`
		} `json:"accountInterestRateSettings"`
		AccrueInterestAfterMaturity      bool    `json:"accrueInterestAfterMaturity"`
		AccrueLateInterest               bool    `json:"accrueLateInterest"`
		InterestApplicationMethod        string  `json:"interestApplicationMethod"`
		InterestBalanceCalculationMethod string  `json:"interestBalanceCalculationMethod"`
		InterestCalculationMethod        string  `json:"interestCalculationMethod"`
		InterestChargeFrequency          string  `json:"interestChargeFrequency"`
		InterestRate                     float64 `json:"interestRate"`
		InterestRateReviewCount          float64 `json:"interestRateReviewCount"`
		InterestRateReviewUnit           string  `json:"interestRateReviewUnit"`
		InterestRateSource               string  `json:"interestRateSource"`
		InterestSpread                   float64 `json:"interestSpread"`
		InterestType                     string  `json:"interestType"`
	} `json:"interestSettings"`
	LastAccountAppraisalDate          time.Time `json:"lastAccountAppraisalDate"`
	LastInterestAppliedDate           time.Time `json:"lastInterestAppliedDate"`
	LastInterestReviewDate            time.Time `json:"lastInterestReviewDate"`
	LastLockedDate                    time.Time `json:"lastLockedDate"`
	LastModifiedDate                  time.Time `json:"lastModifiedDate"`
	LastSetToArrearsDate              time.Time `json:"lastSetToArrearsDate"`
	LastTaxRateReviewDate             time.Time `json:"lastTaxRateReviewDate"`
	LatePaymentsRecalculationMethod   string    `json:"latePaymentsRecalculationMethod"`
	LoanAmount                        float64   `json:"loanAmount"`
	LoanName                          string    `json:"loanName"`
	LockedAccountTotalDueType         string    `json:"lockedAccountTotalDueType"`
	LockedOperations                  []string  `json:"lockedOperations"`
	MigrationEventKey                 string    `json:"migrationEventKey"`
	ModifyInterestForFirstInstallment bool      `json:"modifyInterestForFirstInstallment"`
	Notes                             string    `json:"notes"`
	OriginalAccountKey                string    `json:"originalAccountKey"`
	PaymentHolidaysAccruedInterest    float64   `json:"paymentHolidaysAccruedInterest"`
	PaymentMethod                     string    `json:"paymentMethod"`
	PenaltySettings                   struct {
		LoanPenaltyCalculationMethod string  `json:"loanPenaltyCalculationMethod"`
		PenaltyRate                  float64 `json:"penaltyRate"`
	} `json:"penaltySettings"`
	PlannedInstallmentFees []struct {
		Amount            float64   `json:"amount"`
		ApplyOnDate       time.Time `json:"applyOnDate"`
		EncodedKey        string    `json:"encodedKey"`
		InstallmentKey    string    `json:"installmentKey"`
		InstallmentNumber int       `json:"installmentNumber"`
		PredefinedFeeKey  string    `json:"predefinedFeeKey"`
	} `json:"plannedInstallmentFees"`
	PrepaymentSettings struct {
		ApplyInterestOnPrepaymentMethod string `json:"applyInterestOnPrepaymentMethod"`
		ElementsRecalculationMethod     string `json:"elementsRecalculationMethod"`
		PrepaymentRecalculationMethod   string `json:"prepaymentRecalculationMethod"`
		PrincipalPaidInstallmentStatus  string `json:"principalPaidInstallmentStatus"`
	} `json:"prepaymentSettings"`
	PrincipalPaymentSettings struct {
		Amount                       float64 `json:"amount"`
		EncodedKey                   string  `json:"encodedKey"`
		IncludeFeesInFloorAmount     bool    `json:"includeFeesInFloorAmount"`
		IncludeInterestInFloorAmount bool    `json:"includeInterestInFloorAmount"`
		Percentage                   float64 `json:"percentage"`
		PrincipalCeilingValue        float64 `json:"principalCeilingValue"`
		PrincipalFloorValue          float64 `json:"principalFloorValue"`
		PrincipalPaymentMethod       string  `json:"principalPaymentMethod"`
		TotalDueAmountFloor          float64 `json:"totalDueAmountFloor"`
		TotalDuePayment              string  `json:"totalDuePayment"`
	} `json:"principalPaymentSettings"`
	ProductTypeKey string `json:"productTypeKey"`
	RedrawSettings struct {
		RestrictNextDueWithdrawal bool `json:"restrictNextDueWithdrawal"`
	} `json:"redrawSettings"`
	RescheduledAccountKey string `json:"rescheduledAccountKey"`
	ScheduleSettings      struct {
		BillingCycle struct {
			Days []int `json:"days"`
		} `json:"billingCycle"`
		DefaultFirstRepaymentDueDateOffset int    `json:"defaultFirstRepaymentDueDateOffset"`
		FixedDaysOfMonth                   []int  `json:"fixedDaysOfMonth"`
		GracePeriod                        int    `json:"gracePeriod"`
		GracePeriodType                    string `json:"gracePeriodType"`
		HasCustomSchedule                  bool   `json:"hasCustomSchedule"`
		PaymentPlan                        []struct {
			Amount        int    `json:"amount"`
			EncodedKey    string `json:"encodedKey"`
			ToInstallment int    `json:"toInstallment"`
		} `json:"paymentPlan"`
		PeriodicPayment int `json:"periodicPayment"`
		PreviewSchedule struct {
			NumberOfPreviewedInstalments int `json:"numberOfPreviewedInstalments"`
		} `json:"previewSchedule"`
		PrincipalRepaymentInterval int    `json:"principalRepaymentInterval"`
		RepaymentInstallments      int    `json:"repaymentInstallments"`
		RepaymentPeriodCount       int    `json:"repaymentPeriodCount"`
		RepaymentPeriodUnit        string `json:"repaymentPeriodUnit"`
		RepaymentScheduleMethod    string `json:"repaymentScheduleMethod"`
		ScheduleDueDatesMethod     string `json:"scheduleDueDatesMethod"`
		ShortMonthHandlingMethod   string `json:"shortMonthHandlingMethod"`
	} `json:"scheduleSettings"`
	SettlementAccountKey string    `json:"settlementAccountKey"`
	TaxRate              float64   `json:"taxRate"`
	TerminationDate      time.Time `json:"terminationDate"`
	Tranches             []struct {
		Amount              float64 `json:"amount"`
		DisbursementDetails struct {
			DisbursementTransactionKey string    `json:"disbursementTransactionKey"`
			ExpectedDisbursementDate   time.Time `json:"expectedDisbursementDate"`
		} `json:"disbursementDetails"`
		EncodedKey string `json:"encodedKey"`
		Fees       []struct {
			Amount                  float64 `json:"amount"`
			EncodedKey              string  `json:"encodedKey"`
			Percentage              float64 `json:"percentage"`
			PredefinedFeeEncodedKey string  `json:"predefinedFeeEncodedKey"`
		} `json:"fees"`
		TrancheNumber int `json:"trancheNumber"`
	} `json:"tranches"`
}

type ClientResponse struct {
	ActivationDate time.Time `json:"activationDate"`
	Addresses      []struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		EncodedKey  string `json:"encodedKey"`
		IndexInList int    `json:"indexInList"`
		Latitude    int    `json:"latitude"`
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		Longitude   int    `json:"longitude"`
		ParentKey   string `json:"parentKey"`
		Postcode    string `json:"postcode"`
		Region      string `json:"region"`
	} `json:"addresses"`
	ApprovedDate      time.Time `json:"approvedDate"`
	AssignedBranchKey string    `json:"assignedBranchKey"`
	AssignedCentreKey string    `json:"assignedCentreKey"`
	AssignedUserKey   string    `json:"assignedUserKey"`
	BirthDate         string    `json:"birthDate"`
	ClientRoleKey     string    `json:"clientRoleKey"`
	ClosedDate        time.Time `json:"closedDate"`
	CreationDate      time.Time `json:"creationDate"`
	EmailAddress      string    `json:"emailAddress"`
	EncodedKey        string    `json:"encodedKey"`
	FirstName         string    `json:"firstName"`
	Gender            string    `json:"gender"`
	GroupKeys         []string  `json:"groupKeys"`
	GroupLoanCycle    int       `json:"groupLoanCycle"`
	HomePhone         string    `json:"homePhone"`
	ID                string    `json:"id"`
	IDDocuments       []struct {
		Attachments []struct {
			CreationDate     time.Time `json:"creationDate"`
			EncodedKey       string    `json:"encodedKey"`
			FileName         string    `json:"fileName"`
			FileSize         int       `json:"fileSize"`
			ID               int       `json:"id"`
			LastModifiedDate time.Time `json:"lastModifiedDate"`
			Location         string    `json:"location"`
			Name             string    `json:"name"`
			Notes            string    `json:"notes"`
			OwnerKey         string    `json:"ownerKey"`
			OwnerType        string    `json:"ownerType"`
			Type             string    `json:"type"`
		} `json:"attachments"`
		ClientKey                         string `json:"clientKey"`
		DocumentID                        string `json:"documentId"`
		DocumentType                      string `json:"documentType"`
		EncodedKey                        string `json:"encodedKey"`
		IdentificationDocumentTemplateKey string `json:"identificationDocumentTemplateKey"`
		IndexInList                       int    `json:"indexInList"`
		IssuingAuthority                  string `json:"issuingAuthority"`
		ValidUntil                        string `json:"validUntil"`
	} `json:"idDocuments"`
	LastModifiedDate  time.Time `json:"lastModifiedDate"`
	LastName          string    `json:"lastName"`
	LoanCycle         int       `json:"loanCycle"`
	MiddleName        string    `json:"middleName"`
	MigrationEventKey string    `json:"migrationEventKey"`
	MobilePhone       string    `json:"mobilePhone"`
	MobilePhone2      string    `json:"mobilePhone2"`
	Notes             string    `json:"notes"`
	PortalSettings    struct {
		EncodedKey       string    `json:"encodedKey"`
		LastLoggedInDate time.Time `json:"lastLoggedInDate"`
		PortalState      string    `json:"portalState"`
	} `json:"portalSettings"`
	PreferredLanguage   string `json:"preferredLanguage"`
	ProfilePictureKey   string `json:"profilePictureKey"`
	ProfileSignatureKey string `json:"profileSignatureKey"`
	State               string `json:"state"`
}

type GroupResponse struct {
	Addresses []struct {
		City        string `json:"city"`
		Country     string `json:"country"`
		EncodedKey  string `json:"encodedKey"`
		IndexInList int    `json:"indexInList"`
		Latitude    int    `json:"latitude"`
		Line1       string `json:"line1"`
		Line2       string `json:"line2"`
		Longitude   int    `json:"longitude"`
		ParentKey   string `json:"parentKey"`
		Postcode    string `json:"postcode"`
		Region      string `json:"region"`
	} `json:"addresses"`
	AssignedBranchKey string    `json:"assignedBranchKey"`
	AssignedCentreKey string    `json:"assignedCentreKey"`
	AssignedUserKey   string    `json:"assignedUserKey"`
	CreationDate      time.Time `json:"creationDate"`
	EmailAddress      string    `json:"emailAddress"`
	EncodedKey        string    `json:"encodedKey"`
	GroupMembers      []struct {
		ClientKey string `json:"clientKey"`
		Roles     []struct {
			EncodedKey       string `json:"encodedKey"`
			GroupRoleNameKey string `json:"groupRoleNameKey"`
			RoleName         string `json:"roleName"`
			RoleNameID       string `json:"roleNameId"`
		} `json:"roles"`
	} `json:"groupMembers"`
	GroupName         string    `json:"groupName"`
	GroupRoleKey      string    `json:"groupRoleKey"`
	HomePhone         string    `json:"homePhone"`
	ID                string    `json:"id"`
	LastModifiedDate  time.Time `json:"lastModifiedDate"`
	LoanCycle         int       `json:"loanCycle"`
	MigrationEventKey string    `json:"migrationEventKey"`
	MobilePhone       string    `json:"mobilePhone"`
	Notes             string    `json:"notes"`
	PreferredLanguage string    `json:"preferredLanguage"`
}