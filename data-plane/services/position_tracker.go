package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type LapType int
type LapPhase string

const (
	Outlap LapType = iota
	HotLap
	Inlap
)

const (
	Sector1Index int = 88
	Sector2Index int = 21
	Sector3Index int = 54
)

const (
	Pitlane      LapPhase = "pitlane"
	TrackSector  LapPhase = "tracksector"
	EntrySector1 LapPhase = "entry-sector-1"
	ExitSector1  LapPhase = "exit-sector-1"
	EntrySector2 LapPhase = "entry-sector-2"
	ExitSector2  LapPhase = "exit-sector-2"
	EntrySector3 LapPhase = "entry-sector-3"
	ExitSector3  LapPhase = "exit-sector-3"
)

// Session specific data
var CurrentLapType LapType
var CurrentLapPhase LapPhase
var CurrentLapNumber int
var LiveMicroSectorInfo []MicroSectorInfo
var CurrentLapStartTime time.Time

// Track specific data
var mircoSectors [][]Point
var trackData []TrackData

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type MicroSectorInfo struct {
	Lap           int
	TimeOfArrival time.Time
	Stopwatch     time.Duration
	Index         int
	SectorType    string
}

type TrackData struct {
	Index      int       `json:"index"`
	XPoints    []float64 `json:"x"`
	YPoints    []float64 `json:"y"`
	SectorType string    `json:"type"`
}

func pointInPolygon(p Point, polygon []Point) bool {
	inside := false
	n := len(polygon)

	for i, v := range polygon {
		j := (i + 1) % n
		vi := v
		vj := polygon[j]

		// Check if point is between the y-coords of the edge
		intersects := ((vi.Y > (p.Y)) != (vj.Y > p.Y)) &&
			(p.X < (vj.X-vi.X)*(p.Y-vi.Y)/(vj.Y-vi.Y)+vi.X)

		if intersects {
			inside = !inside
		}
	}
	return inside
}

func StartUp() {
	file, err := os.Open("services/killarney_sectors_v4_annotated.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if err := json.Unmarshal(bytes, &trackData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, td := range trackData {
		var polygon []Point
		for i := 0; i < 4; i++ {
			polygon = append(polygon, Point{td.XPoints[i], td.YPoints[i]})
		}
		mircoSectors = append(mircoSectors, polygon)
	}

	LiveMicroSectorInfo = make([]MicroSectorInfo, len(mircoSectors))
	// Test point
	CurrentLapType = Outlap
	CurrentLapPhase = Pitlane
	CurrentLapNumber = 0

	TrySendTelegramMsg("Starting new session")
	// for _, p := range points {
	// 	sleepTime := time.Millisecond*1000 + time.Millisecond*time.Duration(rand.Intn(201)-100)
	// 	// sleepTime := time.Millisecond*100 + time.Millisecond*time.Duration(rand.Intn(21)-10)
	// 	time.Sleep(sleepTime)
	// 	processGPSData(p)
	// }
}

func TrySendTelegramMsg(s string) {
	telegramKey, ok := os.LookupEnv("TELEGRAMAPI_KEY")
	if ok {
		telegramMsg := url.QueryEscape(s)
		telegramPayload := fmt.Sprintf("https://api.telegram.org/%s/sendMessage?chat_id=1362017106&text=%s", telegramKey, telegramMsg)
		resp, err := http.Get(telegramPayload)
		if err != nil {
			fmt.Println("Could not send the following telegram message:", s)
			return
		}
		defer resp.Body.Close()
	}
}

func ProcessGPSData(p Point) {
	for j, miniSector := range mircoSectors {
		if !pointInPolygon(p, miniSector) {
			continue
		}

		currentTimestamp := time.Now()
		if CurrentLapType == HotLap && (CurrentLapNumber == 2 || CurrentLapNumber == 3) &&
			LiveMicroSectorInfo[trackData[j].Index].Lap == CurrentLapNumber-1 {

			now := currentTimestamp.Sub(CurrentLapStartTime)
			previousTime := LiveMicroSectorInfo[trackData[j].Index].Stopwatch
			fmt.Println("delta", now, previousTime, now-previousTime)
		}

		LiveMicroSectorInfo[trackData[j].Index].Lap = CurrentLapNumber
		LiveMicroSectorInfo[trackData[j].Index].TimeOfArrival = currentTimestamp
		LiveMicroSectorInfo[trackData[j].Index].Index = trackData[j].Index
		LiveMicroSectorInfo[trackData[j].Index].SectorType = trackData[j].SectorType
		//New lap condition
		if trackData[j].SectorType == string(ExitSector1) && CurrentLapPhase == EntrySector1 {
			CurrentLapNumber += 1
			LiveMicroSectorInfo[trackData[j].Index].Lap += 1
			CurrentLapType = HotLap
			CurrentLapPhase = LapPhase(trackData[j].SectorType)

			lastTiming, microSectorGap, ok := findLastEntryTiming(&LiveMicroSectorInfo, trackData[j].Index)
			if !ok {
				fmt.Println("Failed to find last entry time", trackData[j].Index)
				continue
			}

			samplingTimeDiff := LiveMicroSectorInfo[trackData[j].Index].TimeOfArrival.Sub(lastTiming.TimeOfArrival)
			startOfSectorGap, _ := findGapToStartOfSector(trackData[j].Index, strings.Split(trackData[j].SectorType, "-")[2])
			delta := (int(samplingTimeDiff.Milliseconds()) / microSectorGap) * startOfSectorGap
			correctedStartTime := LiveMicroSectorInfo[trackData[j].Index].TimeOfArrival.Add(-time.Duration(delta) * time.Millisecond)
			if !CurrentLapStartTime.IsZero() {
				lapTime := correctedStartTime.Sub(CurrentLapStartTime)
				fmt.Println("Lap Time:", lapTime)
				TrySendTelegramMsg(fmt.Sprint(lapTime))
			}
			CurrentLapStartTime = correctedStartTime

		} else if trackData[j].SectorType != string(TrackSector) {
			CurrentLapPhase = LapPhase(trackData[j].SectorType)
			fmt.Println("Current phase:", CurrentLapPhase, "Lap:", CurrentLapNumber)
		}
		if !CurrentLapStartTime.IsZero() {
			LiveMicroSectorInfo[trackData[j].Index].Stopwatch = LiveMicroSectorInfo[trackData[j].Index].TimeOfArrival.Sub(CurrentLapStartTime)
		}
		break
	}
}

func findGapToStartOfSector(currentIndex int, sectorName string) (int, error) {
	sectorNumber, _ := strconv.Atoi(sectorName)
	switch sectorNumber {
	case 1:
		return currentIndex - Sector1Index, nil
	case 2:
		return currentIndex - Sector2Index, nil
	case 3:
		return currentIndex - Sector3Index, nil
	default:
		return 0, errors.New("unexpected sector type")
	}
}

func findLastEntryTiming(liveMicroSectorInfo *[]MicroSectorInfo, currentIndex int) (*MicroSectorInfo, int, bool) {
	for i := 1; i < 10; i++ {
		lookBehindIndex := (currentIndex - i) % len(*liveMicroSectorInfo)
		microSectorInfo := (*liveMicroSectorInfo)[lookBehindIndex]
		if strings.HasPrefix(microSectorInfo.SectorType, "entry-sector") && !microSectorInfo.TimeOfArrival.IsZero() {
			return &microSectorInfo, i, true
		}
	}
	return nil, 0, false
}
