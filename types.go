package tickerapi

import "encoding/json"

// Timeframe represents the analysis timeframe.
type Timeframe string

const (
	TimeframeDaily  Timeframe = "daily"
	TimeframeWeekly Timeframe = "weekly"
)

// AssetClass represents the type of financial asset.
type AssetClass string

const (
	AssetClassStock  AssetClass = "stock"
	AssetClassCrypto AssetClass = "crypto"
	AssetClassETF    AssetClass = "etf"
	AssetClassAll    AssetClass = "all"
)

// Direction represents bullish/bearish or buying/selling direction.
type Direction string

const (
	DirectionBullish Direction = "bullish"
	DirectionBearish Direction = "bearish"
	DirectionAll     Direction = "all"

	DirectionBuying  Direction = "buying"
	DirectionSelling Direction = "selling"

	DirectionUndervalued Direction = "undervalued"
	DirectionOvervalued  Direction = "overvalued"
)

// OversoldSeverity represents the minimum severity for oversold scans.
type OversoldSeverity string

const (
	OversoldSeverityOversold     OversoldSeverity = "oversold"
	OversoldSeverityDeepOversold OversoldSeverity = "deep_oversold"
)

// ValuationSeverity represents the minimum severity for valuation scans.
type ValuationSeverity string

const (
	ValuationSeverityDeepValue      ValuationSeverity = "deep_value"
	ValuationSeverityExtremePremium ValuationSeverity = "extreme_premium"
)

// VolumeRatioBand represents volume ratio classification.
type VolumeRatioBand string

const (
	VolumeRatioBandExtremelyLow  VolumeRatioBand = "extremely_low"
	VolumeRatioBandLow           VolumeRatioBand = "low"
	VolumeRatioBandNormal        VolumeRatioBand = "normal"
	VolumeRatioBandElevated      VolumeRatioBand = "elevated"
	VolumeRatioBandHigh          VolumeRatioBand = "high"
	VolumeRatioBandExtremelyHigh VolumeRatioBand = "extremely_high"
)

// SummaryOptions configures the Summary request.
type SummaryOptions struct {
	Timeframe *Timeframe `json:"timeframe,omitempty"`
	Date      *string    `json:"date,omitempty"`
}

// CompareOptions configures the Compare request.
type CompareOptions struct {
	Timeframe *Timeframe `json:"timeframe,omitempty"`
	Date      *string    `json:"date,omitempty"`
}

// WatchlistOptions configures the Watchlist request.
type WatchlistOptions struct {
	Timeframe *Timeframe `json:"timeframe,omitempty"`
}

// ScanOversoldOptions configures the ScanOversold request.
type ScanOversoldOptions struct {
	Timeframe   *Timeframe        `json:"timeframe,omitempty"`
	AssetClass  *AssetClass       `json:"asset_class,omitempty"`
	Sector      *string           `json:"sector,omitempty"`
	MinSeverity *OversoldSeverity `json:"min_severity,omitempty"`
	SortBy      *string           `json:"sort_by,omitempty"` // "severity", "days_oversold", "condition_percentile"
	Limit       *int              `json:"limit,omitempty"`   // 1-50
	Date        *string           `json:"date,omitempty"`
}

// ScanBreakoutsOptions configures the ScanBreakouts request.
type ScanBreakoutsOptions struct {
	Timeframe  *Timeframe  `json:"timeframe,omitempty"`
	AssetClass *AssetClass `json:"asset_class,omitempty"`
	Sector     *string     `json:"sector,omitempty"`
	Direction  *Direction  `json:"direction,omitempty"` // "bullish", "bearish", "all"
	SortBy     *string     `json:"sort_by,omitempty"`   // "volume_ratio", "level_strength", "condition_percentile"
	Limit      *int        `json:"limit,omitempty"`     // 1-50
	Date       *string     `json:"date,omitempty"`
}

// ScanUnusualVolumeOptions configures the ScanUnusualVolume request.
type ScanUnusualVolumeOptions struct {
	Timeframe    *Timeframe       `json:"timeframe,omitempty"`
	AssetClass   *AssetClass      `json:"asset_class,omitempty"`
	Sector       *string          `json:"sector,omitempty"`
	MinRatioBand *VolumeRatioBand `json:"min_ratio_band,omitempty"`
	SortBy       *string          `json:"sort_by,omitempty"` // "volume_percentile"
	Limit        *int             `json:"limit,omitempty"`   // 1-50
	Date         *string          `json:"date,omitempty"`
}

// ScanValuationOptions configures the ScanValuation request.
type ScanValuationOptions struct {
	Timeframe   *Timeframe         `json:"timeframe,omitempty"`
	Sector      *string            `json:"sector,omitempty"`
	Direction   *Direction         `json:"direction,omitempty"`   // "undervalued", "overvalued", "all"
	MinSeverity *ValuationSeverity `json:"min_severity,omitempty"` // "deep_value", "extreme_premium"
	SortBy      *string            `json:"sort_by,omitempty"`     // "valuation_percentile", "pe_vs_history"
	Limit       *int               `json:"limit,omitempty"`       // 1-50
	Date        *string            `json:"date,omitempty"`
}

// ScanInsiderActivityOptions configures the ScanInsiderActivity request.
type ScanInsiderActivityOptions struct {
	Timeframe *Timeframe `json:"timeframe,omitempty"`
	Sector    *string    `json:"sector,omitempty"`
	Direction *Direction `json:"direction,omitempty"` // "buying", "selling", "all"
	SortBy    *string    `json:"sort_by,omitempty"`   // "zone_severity", "shares_volume", "net_ratio"
	Limit     *int       `json:"limit,omitempty"`     // 1-50
	Date      *string    `json:"date,omitempty"`
}

// WatchlistRequest is the request body for the Watchlist endpoint.
type WatchlistRequest struct {
	Tickers   []string   `json:"tickers"`
	Timeframe *Timeframe `json:"timeframe,omitempty"`
}

// SummaryResponse is the response from the Summary endpoint.
type SummaryResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// CompareResponse is the response from the Compare endpoint.
type CompareResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// WatchlistResponse is the response from the Watchlist endpoint.
type WatchlistResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// AssetsResponse is the response from the Assets endpoint.
type AssetsResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// ScanOversoldResponse is the response from the ScanOversold endpoint.
type ScanOversoldResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// ScanBreakoutsResponse is the response from the ScanBreakouts endpoint.
type ScanBreakoutsResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// ScanUnusualVolumeResponse is the response from the ScanUnusualVolume endpoint.
type ScanUnusualVolumeResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// ScanValuationResponse is the response from the ScanValuation endpoint.
type ScanValuationResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// ScanInsiderActivityResponse is the response from the ScanInsiderActivity endpoint.
type ScanInsiderActivityResponse struct {
	Data       json.RawMessage `json:"data"`
	RateLimits RateLimits      `json:"-"`
}

// Ptr returns a pointer to the given value. This is a convenience helper
// for constructing option structs with optional fields.
func Ptr[T any](v T) *T {
	return &v
}
