package main

import (
	"fmt"
	"time"
)

func question4() {
	var total int
	fmt.Print("Jumlah Cuti Bersama: ")
	fmt.Scan(&total)

	var joinDateStr string
	fmt.Print("Tanggal join karyawan: ")
	fmt.Scan(&joinDateStr)

	var vacationPlanDateStr string
	fmt.Print("Tanggal rencana cuti: ")
	fmt.Scan(&vacationPlanDateStr)

	var numbersofdayvacation int
	fmt.Print("Durasi cuti (hari): ")
	fmt.Scan(&numbersofdayvacation)

	joinDate, _ := time.Parse("2006-01-02", joinDateStr)
	vacationPlanDate, _ := time.Parse("2006-01-02", vacationPlanDateStr)

	result := true
	var reason string

	eligibleTimeToLeave := joinDate.AddDate(0, 0, 180)

	if vacationPlanDate.Before(eligibleTimeToLeave) {
		result = false
		reason = "Karena belum 180 hari sejak tanggal join karyawan"
	}

	lastDayOfTheYear, _ := time.Parse("2006-01-02", "2021-12-31")
	difBetweenJoinAndLastDayOfTheYear := lastDayOfTheYear.Sub(eligibleTimeToLeave)
	dayDiff := int(difBetweenJoinAndLastDayOfTheYear.Hours() / 24) // number of days

	totalPersonalLeave := 14 - total

	amountOfPersonalLeave := dayDiff / 365 * totalPersonalLeave // number of personal leave
	fmt.Println(amountOfPersonalLeave)

	if numbersofdayvacation > amountOfPersonalLeave {
		result = false
		reason = "Karena hanya boleh mengambil " + fmt.Sprint(amountOfPersonalLeave) + " hari"
	}

	fmt.Println(result)
	if !result {
		fmt.Println(reason)
	}
}
