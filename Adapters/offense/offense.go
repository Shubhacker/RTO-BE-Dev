package offense

func GetFineWithOffense(offenseCommited []string) map[string]int {
	AllOffenseWithFine := make(map[string]int, 0)

	AllOffenseWithFine["improper number plate_1"] = 500
	AllOffenseWithFine["illegal parking_1"] = 500
	AllOffenseWithFine["improper number plate_2"] = 1500
	AllOffenseWithFine["illegal parking_2"] = 1500
	AllOffenseWithFine["Not obeying the orders from the Authorities"] = 2000
	AllOffenseWithFine["Not sharing information"] = 2000
	AllOffenseWithFine["Not adhering to road rules"] = 500
	AllOffenseWithFine["Driving/riding without a valid Driving Licence"] = 5000
	AllOffenseWithFine["Driving/riding without Driving Licence"] = 5000
	AllOffenseWithFine["Driving an unauthorised vehicle without valid licence"] = 5000
	AllOffenseWithFine["Driving/riding a vehicle after disqualification"] = 10000
	AllOffenseWithFine["Overspeeding_Bike"] = 1000
	AllOffenseWithFine["Overspeeding_Car"] = 2000
	AllOffenseWithFine["Rash/dangerous driving_1"] = 1000
	AllOffenseWithFine["Rash/dangerous driving_2"] = 10000
	AllOffenseWithFine["Driving/riding under the influence of intoxicating substances/alcohol_1"] = 10000
	AllOffenseWithFine["Drink and Drive_1"] = 10000
	AllOffenseWithFine["Driving/riding under the influence of intoxicating substances/alcohol_2"] = 15000
	AllOffenseWithFine["Drink and Drive_2"] = 15000
	AllOffenseWithFine["Driving/riding in a mentally/physically unfit state_1"] = 1000
	AllOffenseWithFine["Driving/riding in a mentally/physically unfit state_2"] = 2000
	AllOffenseWithFine["Driving/riding a vehicle without valid motor insurance_1"] = 2000
	AllOffenseWithFine["Driving/riding a vehicle without valid motor insurance_2"] = 4000
	AllOffenseWithFine["Illegal racing and overspeeding_1"] = 5000
	AllOffenseWithFine["Illegal racing and overspeeding_2"] = 10000
	AllOffenseWithFine["Driving an oversized vehicle"] = 5000
	AllOffenseWithFine["Accident-related offences_1"] = 5000
	AllOffenseWithFine["Accident-related offences_2"] = 10000
	AllOffenseWithFine["Driving/riding a vehicle without a valid Registration Certificate_1"] = 5000
	AllOffenseWithFine["Driving/riding a vehicle without a valid Registration Certificate_2"] = 10000
	AllOffenseWithFine["Driving a vehicle without permit"] = 10000
	AllOffenseWithFine["Driving/riding a vehicle while using a mobile phone"] = 5000
	AllOffenseWithFine["Overloading the vehicle"] = 20000
	AllOffenseWithFine["Overloading passengers"] = 1000
	AllOffenseWithFine["Overloading a two-wheeler"] = 2000
	AllOffenseWithFine["Driving a two-wheeler triple seat"] = 2000
	AllOffenseWithFine["Not wearing a helmet while riding a two-wheeler"] = 1000
	AllOffenseWithFine["Not wearing seatbelt while driving"] = 1000
	AllOffenseWithFine["Not giving way for emergency vehicles"] = 10000
	AllOffenseWithFine["Offences committed by juveniles(Below 18 age)"] = 25000
	AllOffenseWithFine["Offence related to enforcing officers (Example: Offering bribe)"] = 2
	AllOffenseWithFine["Aggregator offences (Licence-related)"] = 25000
	AllOffenseWithFine["Usage of horn in Silent Zone"] = 2000
	AllOffenseWithFine["Travelling without a ticket in public transport vehicles"] = 500

	UserCommitedOffense := make(map[string]int, 0)

	for _, offense := range offenseCommited {
		UserCommitedOffense[offense] = AllOffenseWithFine[offense]
	}

	return UserCommitedOffense
}
