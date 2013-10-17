// Package holtwinters import http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
// st[i] = alpha * y[i] / it[i - period] + (1.0 - alpha) * (st[i - 1] + bt[i - 1])
// bt[i] = gamma * (st[i] - st[i - 1]) + (1 - gamma) * bt[i - 1]
// it[i] = beta * y[i] / st[i] + (1.0 - beta) * it[i - period]
// ft[i + m] = (st[i] + (m * bt[i])) * it[i - period + m]
package holtwinters

import (
	"errors"
)

// Forecast method is the entry point. it calculates the initial values and
// returns the forecast for the future m periods.
//
// y - Time series data.
// alpha - Exponential smoothing coefficients for level, trend,
//        seasonal components.
// beta - Exponential smoothing coefficients for level, trend,
//        seasonal components.
// gamma - Exponential smoothing coefficients for level, trend,
//         seasonal components.
// perdiod - A complete season's data consists of L periods. And we need
//           to estimate the trend factor from one period to the next. To
//           accomplish this, it is advisable to use two complete seasons;
//           that is, 2L periods.
// m - Extrapolated future data points.
//   - 4 quarterly,
//   - 7 weekly,
//   - 12 monthly
func Forecast(y []float64, alpha, beta, gamma float64, period, m int) (forecast []float64, err error) {

	if err = validateArguments(y, alpha, beta, gamma, period, m); err != nil {
		forecast = nil
		return
	}

	seasons := len(y) / period
	a0 := initialLevel(y)
	b0 := initialTrend(y, period)
	seasonal := seasonalIndicies(y, period, seasons)

	forecast = calculateHoltWinters(y, a0, b0, alpha, beta, gamma, seasonal, period, m)

	return
}

//
// Validate input.
func validateArguments(y []float64, alpha, beta, gamma float64, period, m int) (err error) {

	if len(y) == 0 {
		err = errors.New("value of y should be not null")
	}

	if m <= 0 {
		err = errors.New("value of m must be greater than 0")
	}

	if m > period {
		err = errors.New("value of m must be <= period")
	}

	if (alpha < 0.0) || (alpha > 1.0) {
		err = errors.New("value of Alpha should satisfy 0.0 <= alpha <= 1.0")
	}

	if (beta < 0.0) || (beta > 1.0) {
		err = errors.New("value of Beta should satisfy 0.0 <= beta <= 1.0")
	}

	if (gamma < 0.0) || (gamma > 1.0) {
		err = errors.New("value of Gamma should satisfy 0.0 <= gamma <= 1.0")
	}
	return
}

// This method realizes the Holt-Winters equations.
// Forecast for m periods.
func calculateHoltWinters(y []float64, a0, b0, alpha, beta, gamma float64, initialSeasonalIndices []float64, period, m int) []float64 {

	st := make([]float64, len(y))
	bt := make([]float64, len(y))
	it := make([]float64, len(y))
	ft := make([]float64, len(y)+m)

	st[1] = a0
	bt[1] = b0

	for i := 0; i < period; i++ {
		it[i] = initialSeasonalIndices[i]
	}

	for i := 2; i < len(y); i++ {

		// overall smoothing
		if (i - period) >= 0 {
			st[i] = alpha*y[i]/it[i-period] + (1.0-alpha)*(st[i-1]+bt[i-1])
		} else {
			st[i] = alpha*y[i] + (1.0-alpha)*(st[i-1]+bt[i-1])
		}

		// trend smoothing
		bt[i] = gamma*(st[i]-st[i-1]) + (1-gamma)*bt[i-1]

		// seasonal smoothing
		if (i - period) >= 0 {
			it[i] = beta*y[i]/st[i] + (1.0-beta)*it[i-period]
		}

		// forecast
		if (i + m) >= period {
			ft[i+m] = (st[i] + (float64(m) * bt[i])) * it[i-period+m]
		}
	}

	return ft
}

// See: http://robjhyndman.com/researchtips/hw-initialization/
func initialLevel(y []float64) float64 {
	return y[0]
}

// See: http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
func initialTrend(y []float64, period int) float64 {

	var sum float64
	sum = 0

	for i := 0; i < period; i++ {
		sum += (y[period+i] - y[i])
	}

	return sum / float64(period*period)
}

// See: http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
func seasonalIndicies(y []float64, period, seasons int) []float64 {

	seasonalAverage := make([]float64, seasons)
	seasonalIndices := make([]float64, period)

	averagedObservations := make([]float64, len(y))

	for i := 0; i < seasons; i++ {
		for j := 0; j < period; j++ {
			seasonalAverage[i] += y[(i*period)+j]
		}
		seasonalAverage[i] /= float64(period)
	}

	for i := 0; i < seasons; i++ {
		for j := 0; j < period; j++ {
			averagedObservations[(i*period)+j] = y[(i*period)+j] / seasonalAverage[i]
		}
	}

	for i := 0; i < period; i++ {
		for j := 0; j < seasons; j++ {
			seasonalIndices[i] += averagedObservations[(j*period)+i]
		}
		seasonalIndices[i] /= float64(seasons)
	}

	return seasonalIndices
}
