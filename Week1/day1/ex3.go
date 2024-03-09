package day1

import "fmt"

type employee interface {
	get_salary(*Time) float32
}

type Freelancer struct {
	salary_per_hour float32
}

type Fulltime struct {
	salary_per_month float32
}

type Contractor struct {
	salary_per_month float32
}

func (freelancer Freelancer) get_salary(time *Time) float32 {
	time.convert_to_hour()
	return freelancer.salary_per_hour * time.hour
}

func (contractor Contractor) get_salary(time *Time) float32 {
	time.convert_to_month()
	return contractor.salary_per_month * time.month
}

func (fulltime Fulltime) get_salary(time *Time) float32 {
	time.convert_to_month()
	return fulltime.salary_per_month * time.month
}

type Time struct {
	hour  float32
	day   float32
	month float32
}

func (time *Time) convert_month_to_day() {
	time.day += time.month * 30
}

func (time *Time) convert_day_to_hour() {
	time.hour += time.day * 24
}

func (time *Time) convert_hour_to_day() {
	time.day += time.hour / 24
}

func (time *Time) convert_day_to_month() {
	time.month += time.day / 30
}

func (time *Time) convert_to_hour() {
	time.convert_month_to_day()
	time.convert_day_to_hour()
}

func (time *Time) convert_to_month() {
	time.convert_hour_to_day()
	time.convert_day_to_month()
}

func calculate_salary(emp employee, time Time) {
	fmt.Printf("Employee's salary %f\n", emp.get_salary(&time))
}

func Ex3() {

	time := Time{
		24,
		1,
		0,
	}

	fl := Freelancer{100}
	calculate_salary(fl, time)

	ft := Fulltime{3000}
	calculate_salary(ft, time)

	ct := Contractor{60000}
	calculate_salary(ct, time)
}
