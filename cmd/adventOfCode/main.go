package main

import (
	"fmt"
	"os"
	"strconv"

	"git.rpjosh.de/RPJosh/go-logger"
	"rpjosh.de/adventOfCode/internal/2023/day_01"
	"rpjosh.de/adventOfCode/internal/2023/day_02"
	"rpjosh.de/adventOfCode/internal/2023/day_03"
	"rpjosh.de/adventOfCode/internal/2023/day_04"
	"rpjosh.de/adventOfCode/internal/2023/day_05"
	"rpjosh.de/adventOfCode/internal/2023/day_06"
	"rpjosh.de/adventOfCode/internal/2023/day_07"
	"rpjosh.de/adventOfCode/internal/2023/day_08"
	"rpjosh.de/adventOfCode/internal/2023/day_09"
	"rpjosh.de/adventOfCode/internal/2023/day_10"
	"rpjosh.de/adventOfCode/internal/2023/day_11"
	"rpjosh.de/adventOfCode/internal/2023/day_12"
	"rpjosh.de/adventOfCode/internal/2023/day_13"
	"rpjosh.de/adventOfCode/internal/2023/day_14"
	"rpjosh.de/adventOfCode/internal/2023/day_15"
	"rpjosh.de/adventOfCode/internal/2023/day_16"
	"rpjosh.de/adventOfCode/internal/2023/day_17"
	"rpjosh.de/adventOfCode/internal/2023/day_18"
	"rpjosh.de/adventOfCode/internal/2023/day_19"
	"rpjosh.de/adventOfCode/internal/2023/day_20"
	"rpjosh.de/adventOfCode/internal/2023/day_21"
	"rpjosh.de/adventOfCode/internal/2023/day_22"
	"rpjosh.de/adventOfCode/internal/2023/day_23"
	"rpjosh.de/adventOfCode/internal/2023/day_24"
	"rpjosh.de/adventOfCode/internal/2023/day_25"
	"rpjosh.de/adventOfCode/pkg/utils"
)

func main() {

	// Configure logger
	logger.SetGlobalLogger(&logger.Logger{
		Level: logger.LevelDebug,
		File:  &logger.FileLogger{},
	})

	year := 2023
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

	testRun := len(os.Args) >= 4 && os.Args[3] != ""
	if testRun {
		data = utils.GetFromClipboard()
	}

	var result string
	if part == 1 {
		result = challenge.Part1(data)
	} else {
		result = challenge.Part2(data)
	}

	if !testRun {
		utils.CopyToClipboard(result)
	} else {
		fmt.Print(result + "\n")
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
