package main

import "fmt"


func cekJamLebih(jamParkir int) int {

	var jamParkirLebih int

    if jamParkir > 2 {
        jamParkirLebih = jamParkir - 2
    } else {
        jamParkirLebih = 0
    }
    return jamParkirLebih
}

func hitungBiayaParkir(jamParkir int, statusMember bool, hariLibur bool) int {

	tarifDasar := 5000
	tarifJamKetiga := 2000
	tarifHariLibur := 3000
	totalBiaya := 0

	if jamParkir <= 2 {
		totalBiaya = tarifDasar
	}else if jamParkir <= 0{
		fmt.Print("Masukkan jam parkir yang valid\n")
	}else{
		totalBiaya = tarifDasar + (tarifJamKetiga * cekJamLebih(jamParkir))
	}

	if statusMember{
		if jamParkir <=5{
			totalBiaya = totalBiaya - (totalBiaya * 50 / 100)
		}else{
			totalBiaya = totalBiaya - (totalBiaya * 30 / 100)
		}
		
	}

	if hariLibur{
		totalBiaya+=tarifHariLibur
	}

	fmt.Print("Total Biaya Parkir: ", totalBiaya, "\n")
	return totalBiaya
}

func main() {
    // Test case 1: 4 jam, bukan member, bukan hari libur
    hitungBiayaParkir(4, false, false)
    // Test case 2: 2 jam, member, hari libur
    hitungBiayaParkir(2, true, true)
}