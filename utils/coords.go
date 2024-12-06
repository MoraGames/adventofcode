package utils

type CoordInterface interface {
	GetCol() int
	GetRow() int
}

type MatrixCoord struct {
	Row int
	Col int
}

func (c MatrixCoord) GetRow() int { return c.Row }
func (c MatrixCoord) GetCol() int { return c.Col }
