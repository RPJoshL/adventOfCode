package main

import (
	"os"
	"strconv"

	"rpjosh.de/adventOfCode2022/internal/day_01"
	"rpjosh.de/adventOfCode2022/internal/day_02"
	"rpjosh.de/adventOfCode2022/internal/day_03"
	"rpjosh.de/adventOfCode2022/internal/day_04"
	"rpjosh.de/adventOfCode2022/internal/day_05"
	"rpjosh.de/adventOfCode2022/internal/day_06"
	"rpjosh.de/adventOfCode2022/internal/day_07"
	"rpjosh.de/adventOfCode2022/internal/day_08"
	"rpjosh.de/adventOfCode2022/internal/day_09"
	"rpjosh.de/adventOfCode2022/internal/day_10"
	"rpjosh.de/adventOfCode2022/internal/day_11"
	"rpjosh.de/adventOfCode2022/internal/day_12"
	"rpjosh.de/adventOfCode2022/internal/day_13"
	"rpjosh.de/adventOfCode2022/internal/day_14"
	"rpjosh.de/adventOfCode2022/internal/day_15"
	"rpjosh.de/adventOfCode2022/internal/day_16"
	"rpjosh.de/adventOfCode2022/internal/day_17"
	"rpjosh.de/adventOfCode2022/internal/day_18"
	"rpjosh.de/adventOfCode2022/internal/day_19"
	"rpjosh.de/adventOfCode2022/internal/day_20"
	"rpjosh.de/adventOfCode2022/internal/day_21"
	"rpjosh.de/adventOfCode2022/internal/day_22"
	"rpjosh.de/adventOfCode2022/internal/day_23"
	"rpjosh.de/adventOfCode2022/internal/day_24"
	"rpjosh.de/adventOfCode2022/internal/day_25"
	"rpjosh.de/adventOfCode2022/pkg/utils"
)

func main() {
	year := 2022
	day := 1
	if len(os.Args) >= 2 && os.Args[1] != "" {
		var err error
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			utils.PrintError("Expected the day of the challenge (integer) as argument: %s", err)
		}
	}
	data := utils.GetInputData(year, day)

	challenge := getAllChallenges()[day-1]
	part := 1
	if len(os.Args) >= 3 && os.Args[2] != "" {
		var err error
		part, err = strconv.Atoi(os.Args[2])
		if err != nil {
			utils.PrintError("Expected the part of the challenge (integer) as argument: %s", err)
		}
	}

	if part == 1 {
		utils.CopyToClipboard(challenge.Part1(data))
	} else {
		utils.CopyToClipboard(challenge.Part2(data))
	}
}

type Challenge interface {
	Part1(string) string
	Part2(string) string
}

func getAllChallenges() []Challenge {
	return []Challenge{
		&day_01.Day{},
		&day_02.Day{},
		&day_03.Day{},
		&day_04.Day{},
		&day_05.Day{},
		&day_06.Day{},
		&day_07.Day{},
		&day_08.Day{},
		&day_09.Day{},
		&day_10.Day{},
		&day_11.Day{},
		&day_12.Day{},
		&day_13.Day{},
		&day_14.Day{},
		&day_15.Day{},
		&day_16.Day{},
		&day_17.Day{},
		&day_18.Day{},
		&day_19.Day{},
		&day_20.Day{},
		&day_21.Day{},
		&day_22.Day{},
		&day_23.Day{},
		&day_24.Day{},
		&day_25.Day{},
	}
}
