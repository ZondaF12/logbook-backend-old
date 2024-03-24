package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ZondaF12/logbook-backend/internal/models"
)

func FetchVehicleDetails(postData models.NewVehiclePostData) (models.NewVehicle, error) {
	vehicleData, err := DoVehicleInfoRequest(postData.Registration)
	if err != nil {
		return models.NewVehicle{}, err
	}

	motData, err := DoVehicleMotRequest(postData.Registration)
	if err != nil {
		return models.NewVehicle{}, err
	}

	var taxDate string
	if vehicleData.TaxDueDate != "" {
		taxDate = vehicleData.TaxDueDate
	} else {
		taxDate = vehicleData.TaxStatus
	}

	var motDate string
	if vehicleData.MotExpiryDate != "" {
		motDate = vehicleData.MotExpiryDate
	} else {
		motDate = motData[0].MotTestExpiryDate
	}

	var registeredDate string
	if motData[0].FirstUsedDate != "" {
		registeredDate = strings.Replace(motData[0].FirstUsedDate, ".", "-", 2)
	} else {
		date, err := time.Parse("2006-01-02", motDate)
		if err != nil {
			fmt.Println(err)
		}

		registeredDate = date.AddDate(-3, 0, 1).Format("2006-01-02")
	}

	var model string
	if postData.Model != "" {
		model = postData.Model
	} else {
		model = motData[0].Model

	}

	newVehicle := models.NewVehicle{
		Registration: postData.Registration,
		Color:        motData[0].PrimaryColour,
		EngineSize:   uint16(vehicleData.EngineCapacity),
		Make:         vehicleData.Make,
		Model:        model,
		MotDate:      motDate,
		Registered:   registeredDate,
		TaxDate:      taxDate,
		Year:         uint16(vehicleData.YearOfManufacture),
		Description:  postData.Description,
		Nickname:     postData.Nickname,
		Images:       postData.Images,
	}

	return newVehicle, nil
}

var (
	dvla_api_key = os.Getenv("DVLA_API")
	dvsa_api_key = os.Getenv("DVSA_MOT_API_KEY")
)

func DoVehicleInfoRequest(registration string) (models.VehicleData, error) {
	jsonBody := []byte(fmt.Sprintf(`{"registrationNumber": "%s"}`, registration))
	bodyReader := bytes.NewBuffer(jsonBody)

	requestURL := "https://driver-vehicle-licensing.api.gov.uk/vehicle-enquiry/v1/vehicles"
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return models.VehicleData{}, fmt.Errorf("could not create request %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", dvla_api_key)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.VehicleData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.VehicleData{}, errors.New("invalid registration number")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.VehicleData{}, err
	}

	var vehicleResponse models.VehicleData
	if err = json.Unmarshal(body, &vehicleResponse); err != nil {
		fmt.Printf("Error: %v", err)
	}

	return vehicleResponse, nil
}

func DoVehicleMotRequest(registration string) (models.MotData, error) {
	requestURL := "https://beta.check-mot.service.gov.uk/trade/vehicles/mot-tests/?registration=" + registration
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return models.MotData{}, fmt.Errorf("could not create request %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", dvsa_api_key)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.MotData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.MotData{}, errors.New("invalid registration number")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.MotData{}, err
	}

	var vehicleResponse models.MotData
	if err = json.Unmarshal(body, &vehicleResponse); err != nil {
		fmt.Printf("Error: %v", err)
	}

	return vehicleResponse, nil
}
