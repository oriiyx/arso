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
				strength::float AS strength  
		FROM (SELECT id,
					  polygon ->> 'name' AS name,
					  polygon ->> 'strength' AS strength,
					  ST_MakePolygon(ST_AddPoint(line, ST_STARTPOINT(line))) AS geom
			  FROM (SELECT id,
							polygon,
							ST_MakeLine(array_agg(ST_SetSRID(ST_MakePoint((v ->> 'lon'):: FLOAT, (v ->> 'lat'):: FLOAT),
															4326))) AS line
					FROM "Message",
						  jsonb_array_elements(polygons) AS polygon, jsonb_array_elements(polygon -> 'polygon') WITH ORDINALITY arr(v, idx)
					GROUP BY id, POLYGON) AS line_query) AS subquery
		WHERE ST_CONTAINS(geom, ST_SetSRID(ST_MakePoint(?, ?), 4326))
		ORDER BY strength DESC
		LIMIT 1;`

	if err := db.Bun.NewRaw(query, input.Lon, input.Lat).Scan(ctx, &msData); err != nil {
		return nil, err
	}

	return msData, nil
}
