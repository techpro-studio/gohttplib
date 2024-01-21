package location

import (
	"github.com/techpro-studio/gohttplib"
	"github.com/techpro-studio/gohttplib/validator"
	"net/http"
	"strconv"
)

func ParseGeoLocation(longitudeRaw string, latitudeRaw string) (float64, float64, error) {
	longitude, err := strconv.ParseFloat(longitudeRaw, 64)
	if err != nil {
		return 0, 0, gohttplib.NewServerError(400, "longitude", "INVALID_FLOAT", "INVALID_FLOAT", nil)
	}
	latitude, err := strconv.ParseFloat(latitudeRaw, 64)
	if err != nil {
		return 0, 0, gohttplib.NewServerError(400, "longitude", "INVALID_FLOAT", "INVALID_FLOAT", nil)
	}

	errs := validator.ValidateValue(longitude, validator.LongitudeValidators("longitude"))
	if len(errs) > 0 {
		return 0, 0, gohttplib.ServerError{StatusCode: 400, Errors: gohttplib.Errors{Errors: errs}}
	}
	errs = validator.ValidateValue(latitude, validator.LatitudeValidators("latitude"))
	if len(errs) > 0 {
		return 0, 0, gohttplib.ServerError{StatusCode: 400, Errors: gohttplib.Errors{Errors: errs}}
	}
	return longitude, latitude, nil
}

func LocationParametersFromRequest(req *http.Request, defaultMaxDistance int64) (*LocationParameters, error) {
	longitudeRaw := gohttplib.GetParameterFromURLInRequest(req, "longitude")
	latitudeRaw := gohttplib.GetParameterFromURLInRequest(req, "latitude")
	if longitudeRaw == nil || latitudeRaw == nil {
		return nil, nil
	}

	longitude, latitude, err := ParseGeoLocation(*longitudeRaw, *latitudeRaw)

	var minDistance int64 = 0
	maxDistance := defaultMaxDistance

	maxDistanceRaw := gohttplib.GetParameterFromURLInRequest(req, "max_distance")
	minDistanceRaw := gohttplib.GetParameterFromURLInRequest(req, "min_distance")

	if maxDistanceRaw != nil {

		maxDistance, err = strconv.ParseInt(*maxDistanceRaw, 10, 64)
		if err != nil {
			return nil, gohttplib.NewServerError(400, "max_distance", "INVALID_INT", "INVALID_INT", nil)
		}

		errs := validator.ValidateValue(maxDistance, validator.RequiredIntValidators("max_distance", validator.DistanceValidator("max_distance")))
		if len(errs) > 0 {
			return nil, gohttplib.ServerError{StatusCode: 400, Errors: gohttplib.Errors{Errors: errs}}
		}
	}

	if minDistanceRaw != nil {

		minDistance, err = strconv.ParseInt(*minDistanceRaw, 10, 64)
		if err != nil {
			return nil, gohttplib.NewServerError(400, "min_distance", "INVALID_INT", "INVALID_INT", nil)
		}

		errs := validator.ValidateValue(minDistance, validator.RequiredIntValidators("min_distance", validator.DistanceValidator("min_distance")))
		if len(errs) > 0 {
			return nil, gohttplib.ServerError{StatusCode: 400, Errors: gohttplib.Errors{Errors: errs}}
		}
	}

	if minDistance > maxDistance {
		return nil, gohttplib.NewServerError(400, "min_distance", "Min distance more tham max", "MIN_MAX_ERROR", nil)
	}

	return &LocationParameters{Latitude: latitude, Longitude: longitude, MinDistance: minDistance, MaxDistance: maxDistance}, nil

}
