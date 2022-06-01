package util

import (
	"math"

	"github.com/zeromicro/go-zero/core/logx"
)

const EARTH_RADIUS = 6378.137
const RAD = math.Pi / 180.0

// Check points distance valid (in km)
func CheckPointsDistance(clat float64, clng float64, olat float64, olng float64, max_range float64) bool {
	clat = clat * RAD
	clng = clng * RAD
	olat = olat * RAD
	olng = olng * RAD
	theta := olng - clng
	dist := math.Acos(math.Sin(clat)*math.Sin(olat) + math.Cos(clat)*math.Cos(olat)*math.Cos(theta))
	logx.Info("Distance: ", dist, " Max Range: ", max_range)
	return dist <= max_range
}
