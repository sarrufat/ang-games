package common

// Piece
type Piece struct {
	// Not used
	//	Label   string `json:"label"`
	Letter  string `json:"letter"`
	Npieces int    `json:"npieces"`
}

// Problem
type Problem struct {
	Dim    string  `json:"dim""`
	Pieces []Piece `json:"pieces"`
}

type Combination struct {
	Positions   []ResultPosition `json:"positions"`

}
// Result
type Result struct {
	Done        bool             `json:"done"`
	Millis      int64            `json:"ms"`
	NIterations int32            `json:"iterations"`
	Combination   []Combination `json:"combinations"`
}
type TResult struct {
	Result
	TaskId string
}
// Position
type ResultPosition struct {
	Piece string `json:"piece"`
	X     int    `json:"x"`
	Y     int    `json:"y"`
}
