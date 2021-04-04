package v1

import "testing"

func TestLatitude(t *testing.T) {
	type latitudeTest struct {
		latitudes []float64
		expected  []bool
	}
	latitudes := []float64{100.0, 21.2, 32.4, -34.2, -150}
	expected := []bool{false, true, true, true, false}
	testData := latitudeTest{
		latitudes: latitudes,
		expected:  expected,
	}
	for i := 0; i < len(testData.latitudes); i++ {
		if output := checkLatitude(testData.latitudes[i]); output != testData.expected[i] {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", testData.latitudes[i], output, testData.expected[i])
		}
	}
}
func TestLongitude(t *testing.T) {
	type longitudeTest struct {
		longitudes []float64
		expected   []bool
	}
	longitudes := []float64{1009.0, 21.2, 32.4, -34.2, -2150}
	expected := []bool{false, true, true, true, false}
	testData := longitudeTest{
		longitudes: longitudes,
		expected:   expected,
	}
	for i := 0; i < len(testData.longitudes); i++ {
		if output := checkLongitude(testData.longitudes[i]); output != testData.expected[i] {
			t.Error("Test Failed: {} inputted, {} expected, received: {}")
		}
	}
}
