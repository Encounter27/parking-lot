package parking

type lot struct {
	floors []floor
}

// Considering the fact that the parking management software will be completely
// won by parking woner. Parking woner can have more than one parking in the city
// and he wants to centralize the management of all parking he owns.
var allParkings map[string]lot

func createParkingFromConfig(c Config) {

}

func AddAParking(c Config) error {

	return nil
}
