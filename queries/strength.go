package queries

import (
	"context"
	"errors"

	"github.com/oriiyx/arso/db"
)

type MaxStrength struct {
	Lat float64
	Lon float64
}

type MaxStrengthData struct {
	ID       int
	Name     string
	Area     float64
	Strength float64
}

func (ms MaxStrength) Validate() error {
	if ms.Lat == 0 || ms.Lon == 0 {
		return errors.New("latitude and longitude are required")
	}

	return nil
}

func (ms MaxStrength) Execute(input MaxStrength) (maxStrengthData []MaxStrengthData, error error) {
	if err := ms.Validate(); err != nil {
		return nil, err
	}

	ctx := context.Background()
	msData := make([]MaxStrengthData, 0)

	query := `
    SELECT id,
           name,
           ST_AREA(geom) AS area,
           strength AS strength_numeric  -- strength is already a float
    FROM public."Polygon"
    WHERE ST_CONTAINS(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326))
    ORDER BY strength_numeric DESC
    LIMIT 1;`

	if err := db.Bun.NewRaw(query, input.Lon, input.Lat).Scan(ctx, &msData); err != nil {
		return nil, err
	}

	return msData, nil
}
