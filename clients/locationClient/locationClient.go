package locationClient

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	azure "github.com/MSOpenTech/azure-sdk-for-go"
	"strings"
)

const (
	azureLocationListURL   = "locations"
	invalidLocationError   = "Invalid location: %s. Available locations: %s"
	paramNotSpecifiedError = "Parameter %s is not specified."
)

func ResolveLocation(location string) error {
	if len(location) == 0 {
		return fmt.Errorf(paramNotSpecifiedError, "location")
	}

	locations, err := GetLocationList()
	if err != nil {
		return err
	}

	for _, existingLocation := range locations.Locations {
		if existingLocation.Name != location {
			continue
		}

		return nil
	}

	var availableLocations bytes.Buffer
	for _, existingLocation := range locations.Locations {
		availableLocations.WriteString(existingLocation.Name + ", ")
	}

	return errors.New(fmt.Sprintf(invalidLocationError, location, strings.Trim(availableLocations.String(), ", ")))
}

func GetLocationList() (LocationList, error) {
	locationList := LocationList{}

	response, err := azure.SendAzureGetRequest(azureLocationListURL)
	if err != nil {
		return locationList, err
	}

	err = xml.Unmarshal(response, &locationList)
	if err != nil {
		return locationList, err
	}

	return locationList, nil
}

func GetLocation(location string) (*Location, error) {
	if len(location) == 0 {
		return nil, fmt.Errorf(paramNotSpecifiedError, "location")
	}

	locations, err := GetLocationList()
	if err != nil {
		return nil, err
	}

	for _, existingLocation := range locations.Locations {
		if existingLocation.Name != location {
			continue
		}

		return &existingLocation, nil
	}

	var availableLocations bytes.Buffer
	for _, existingLocation := range locations.Locations {
		availableLocations.WriteString(existingLocation.Name + ", ")
	}

	return nil, errors.New(fmt.Sprintf(invalidLocationError, location, strings.Trim(availableLocations.String(), ", ")))
}
