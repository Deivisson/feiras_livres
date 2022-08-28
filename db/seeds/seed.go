package seeds

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Deivisson/feiras_livres/domain"
)

func ImportCsvFile(repo domain.FairRepository) {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if ok, err := repo.HasAny(); err != nil {
		log.Fatal("Error on check if has any fair")
	} else if ok {
		return
	}

	file, err := os.Open(fmt.Sprintf("%s/db/seeds/DEINFO_AB_FEIRASLIVRES_2014.csv", path))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fairs := []domain.Fair{}
	firstRow := true
	for scanner.Scan() {
		if firstRow {
			firstRow = false
			continue
		}
		parts := strings.Split(scanner.Text(), ",")
		// The last item when doesn't have value in the last column, It'll split into only 16 parts.
		// The code bellow gonna equalize columns to avoid "index out of range" exceptions
		if len(parts) < 17 {
			parts = append(parts, "")
		}
		fair := domain.Fair{
			Id: parts[0],
			FairCommon: domain.FairCommon{
				Longitude:         parts[1],
				Latitude:          parts[2],
				Sector:            parts[3],
				Area:              parts[4],
				DistrictCode:      parts[5],
				District:          parts[6],
				SubprefectureCode: parts[7],
				SubprefectureName: parts[8],
				Region5:           parts[9],
				Region8:           parts[10],
				FairName:          parts[11],
				Registry:          parts[12],
				Address:           parts[13],
				Number:            parserNumber(parts[14]),
				Neighborhood:      parts[15],
				Reference:         parts[16],
			},
		}
		fairs = append(fairs, fair)
	}

	if err := repo.BulkCreate(fairs); err != nil {
		log.Fatal(err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func parserNumber(value string) string {
	fmt.Println(value)
	if value == "S/N" {
		return value
	}
	return strings.Replace(value, ".000000", "", -1)
}
