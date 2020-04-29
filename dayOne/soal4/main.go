package main

import (
	"fmt"
	"time"
)

func main() {
	january := time.Date(2020, time.January, 3, 0, 0, 0, 0, time.UTC)
	february := time.Date(2020, time.February, 10, 0, 0, 0, 0, time.UTC)

	different := january.Sub(february)

	days := int(different.Hours() / 24)

	fmt.Printf("Dari tanggal 3 January 2020 sampai 10 February 2020 sudah %d hari yang terlewati \n", days)
}
