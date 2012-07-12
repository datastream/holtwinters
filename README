# Holt-Winters Triple Exponential Smoothing Algorithm

Based on ![Ruby version](https://github.com/cmdrkeene/holt_winters)

![Algorithm](http://www.itl.nist.gov/div898/handbook/pmc/section4/eqns/ts26.gif)

    The equations are intended to give more weight to recent observations and less weights to observations further in the past.
    These weights are geometrically decreasing by a constant ratio.

# Usage

## Forecast()

It calculates the initial values and returns the forecast for __m__ periods.

    # y           Time series array
    # alpha       Level smoothing coefficient
    # beta        Trend smoothing coefficient (increasing beta tightens fit)
    # gamma       Seasonal smoothing coefficient
    # period      A complete season's data consists of L periods. And we need
    #             to estimate the trend factor from one period to the next. To
    #             accomplish this, it is advisable to use two complete seasons;
    #             that is, 2L periods.
    # m           Extrapolated future data points
    #             - 4 quarterly
    #             - 7 weekly
    #             - 12 monthly
    Frecast(y, alpha, beta, gamma, period, m)



For details, see:
http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc43.htm
http://www.itl.nist.gov/div898/handbook/pmc/section4/pmc435.htm
