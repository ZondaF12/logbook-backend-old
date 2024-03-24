package models

type (
	Vehicle struct {
		ID            string   `json:"id"`
		UserID        string   `json:"user_id"`
		Registration  string   `json:"registration"`
		Color         string   `json:"color"`
		Description   string   `json:"description"`
		EngineSize    uint16   `json:"engine_size"`
		Images        []string `json:"images"`
		InsuranceDate string   `json:"insurance_date"`
		Make          string   `json:"make"`
		Model         string   `json:"model"`
		MotDate       string   `json:"mot_date"`
		Nickname      string   `json:"nickname"`
		Registered    string   `json:"registered"`
		ServiceDate   string   `json:"service_date"`
		TaxDate       string   `json:"tax_date"`
		Year          uint16   `json:"year"`
	}

	NewVehicle struct {
		Registration  string   `json:"registration"`
		Color         string   `json:"color"`
		Description   string   `json:"description"`
		EngineSize    uint16   `json:"engine_size"`
		Images        []string `json:"images"`
		InsuranceDate string   `json:"insurance_date"`
		Make          string   `json:"make"`
		Model         string   `json:"model"`
		MotDate       string   `json:"mot_date"`
		Nickname      string   `json:"nickname"`
		Registered    string   `json:"registered"`
		ServiceDate   string   `json:"service_date"`
		TaxDate       string   `json:"tax_date"`
		Year          uint16   `json:"year"`
	}

	VehicleData struct {
		RegistrationNumber       string `json:"registrationNumber"`
		TaxStatus                string `json:"taxStatus"`
		TaxDueDate               string `json:"taxDueDate"`
		ArtEndDate               string `json:"artEndDate"`
		MotStatus                string `json:"motStatus"`
		Make                     string `json:"make"`
		YearOfManufacture        int    `json:"yearOfManufacture"`
		EngineCapacity           int    `json:"engineCapacity"`
		Co2Emissions             int    `json:"co2Emissions"`
		FuelType                 string `json:"fuelType"`
		MarkedForExport          bool   `json:"markedForExport"`
		Colour                   string `json:"colour"`
		TypeApproval             string `json:"typeApproval"`
		RevenueWeight            int    `json:"revenueWeight"`
		EuroStatus               string `json:"euroStatus"`
		DateOfLastV5CIssued      string `json:"dateOfLastV5CIssued"`
		MotExpiryDate            string `json:"motExpiryDate"`
		Wheelplan                string `json:"wheelplan"`
		MonthOfFirstRegistration string `json:"monthOfFirstRegistration"`
	}

	MotData []struct {
		Registration      string     `json:"registration"`
		Make              string     `json:"make"`
		Model             string     `json:"model"`
		FirstUsedDate     string     `json:"firstUsedDate"`
		FuelType          string     `json:"fuelType"`
		PrimaryColour     string     `json:"primaryColour"`
		MotTestExpiryDate string     `json:"MotTestExpiryDate"`
		MotTests          []MotTests `json:"motTests"`
	}

	MotTests struct {
		CompletedDate  string `json:"completedDate"`
		TestResult     string `json:"testResult"`
		ExpiryDate     string `json:"expiryDate"`
		OdometerValue  string `json:"odometerValue"`
		OdometerUnit   string `json:"odometerUnit"`
		MotTestNumber  string `json:"motTestNumber"`
		RfrAndComments []any  `json:"rfrAndComments"`
	}

	NewVehiclePostData struct {
		Registration string   `json:"registration"`
		Images       []string `json:"images"`
		Nickname     string   `json:"nickname"`
		Model        string   `json:"model"`
		Description  string   `json:"description"`
	}
)
