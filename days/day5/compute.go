package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/naruto678/aoc-go/globals"
)

const (
	SEED_TO_SOIL_MAP            = "seed-to-soil map:"
	SOIL_TO_FERTILIZER_MAP      = "soil-to-fertilizer map:"
	FERTILIZER_TO_WATER_MAP     = "fertilizer-to-water map:"
	WATER_TO_LIGHT_MAP          = "water-to-light map:"
	LIGHT_TO_TEMPERATURE_MAP    = "light-to-temperature map:"
	TEMPERATURE_TO_HUMIDITY_MAP = "temperature-to-humidity map:"
	HUMIDITY_TO_LOCATION_MAP    = "humidity-to-location map:"
)

type Pair struct {
	dest, src, length int
}

type Almanac struct {
	seeds                 []int
	seedToSoilMap         []Pair
	soilToFertilizerMap   []Pair
	fertilizerToWaterMap  []Pair
	waterToLightMap       []Pair
	lightToTempMap        []Pair
	tempToHumidityMap     []Pair
	humidityToLocationMap []Pair
}

func NewAlmanac(content string) *Almanac {
	lines := strings.Split(content, "\n")
	all_seeds, found := strings.CutPrefix(lines[0], "seeds: ")
	if !found {
		panic("invalid format")
	}

	almanac := &Almanac{
		seeds:                 []int{},
		seedToSoilMap:         []Pair{},
		soilToFertilizerMap:   []Pair{},
		fertilizerToWaterMap:  []Pair{},
		waterToLightMap:       []Pair{},
		lightToTempMap:        []Pair{},
		tempToHumidityMap:     []Pair{},
		humidityToLocationMap: []Pair{},
	}

	for _, seed := range strings.Fields(all_seeds) {
		seedId, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		almanac.seeds = append(almanac.seeds, seedId)
	}

	parseFields := func(line string) (int, int, int) {
		fields := strings.Fields(line)
		dest, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		src, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		length, err := strconv.Atoi(fields[2])
		if err != nil {
			panic(err)
		}
		return dest, src, length
	}

	var currentMap *[]Pair

	for i := 1; i < len(lines); i++ {
		switch lines[i] {
		case SEED_TO_SOIL_MAP:
			currentMap = &almanac.seedToSoilMap
		case SOIL_TO_FERTILIZER_MAP:
			currentMap = &almanac.soilToFertilizerMap
		case FERTILIZER_TO_WATER_MAP:
			currentMap = &almanac.fertilizerToWaterMap
		case WATER_TO_LIGHT_MAP:
			currentMap = &almanac.waterToLightMap
		case LIGHT_TO_TEMPERATURE_MAP:
			currentMap = &almanac.lightToTempMap
		case TEMPERATURE_TO_HUMIDITY_MAP:
			currentMap = &almanac.tempToHumidityMap
		case HUMIDITY_TO_LOCATION_MAP:
			currentMap = &almanac.humidityToLocationMap
		default:

			if currentMap != nil && len(strings.TrimSpace(lines[i])) > 0 {
				fields := strings.Fields(lines[i])
				if len(fields) != 3 {
					panic(fmt.Errorf("%s invalid fields ", fields))
				}
				dest, src, length := parseFields(lines[i])
				pair := Pair{dest, src, length}
				*currentMap = append(*currentMap, pair)
			}
		}

	}

	return almanac

}

func (a *Almanac) getMap(id int) []Pair {
	switch id {
	case 1:
		return a.seedToSoilMap
	case 2:
		return a.soilToFertilizerMap
	case 3:
		return a.fertilizerToWaterMap
	case 4:
		return a.waterToLightMap
	case 5:
		return a.lightToTempMap
	case 6:
		return a.tempToHumidityMap
	case 7:
		return a.humidityToLocationMap
	default:
		panic("invalid count")
	}
}

func (a *Almanac) GetLocationNumber(seedId int) int {
	id := seedId
	for i := 1; i <= 7; i++ {
		pairs := a.getMap(i)
		for _, pair := range pairs {
			src, dest, length := pair.src, pair.dest, pair.length
			//oldId := id
			if id >= src && id < src+length {
				diff := id - src
				id = diff + dest
				break
			}
			//fmt.Println(fmt.Sprintf("%d ---> %d %v %d", oldId, id, pairs, src))
		}

	}
	return id

}

func computeFirst(content []byte) {
	strContent := string(content)
	almanac := NewAlmanac(strContent)
	min_location := math.MaxInt
	for _, seed := range almanac.seeds {
		min_location = min(min_location, almanac.GetLocationNumber(seed))
	}
	fmt.Println(fmt.Sprintf("computed the first part %d", min_location))

}

func computeSecond(content []byte) {

}

func init() {
	globals.FuncMap["5-1"] = computeFirst
	globals.FuncMap["5-2"] = computeSecond
}
