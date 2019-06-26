package main

// LogicKeeper performs numerous checks on a Wumpus to make sure their stats
// are within range
//
// Requires 2 Arguments
// UserWumpus Wumpus should be the Wumpus you want to check and correct
//
// Returns CorrectedWumpus Wumpus which is the original Wumpus with values
// that are within range
func LogicKeeper(UserWumpus Wumpus) (CorrectedWumpus Wumpus) {
	if UserWumpus.Age > 14 {
		CorrectedWumpus.Age = 14
	} else if UserWumpus.Age < 0 {
		CorrectedWumpus.Age = 0
	} else {
		CorrectedWumpus.Age = UserWumpus.Age
	}

	if UserWumpus.Health > 10 {
		CorrectedWumpus.Health = 10
	} else if UserWumpus.Health < 0 {
		CorrectedWumpus.Health = 0
	} else {
		CorrectedWumpus.Health = UserWumpus.Health
	}

	if UserWumpus.Energy > 10 {
		CorrectedWumpus.Energy = 10
	} else if UserWumpus.Energy < 0 {
		CorrectedWumpus.Energy = 0
	} else {
		CorrectedWumpus.Energy = UserWumpus.Energy
	}

	if UserWumpus.Happiness > 10 {
		CorrectedWumpus.Happiness = 10
	} else if UserWumpus.Happiness < 0 {
		CorrectedWumpus.Happiness = 0
	} else {
		CorrectedWumpus.Happiness = UserWumpus.Happiness
	}

	if UserWumpus.Credits < 0 {
		CorrectedWumpus.Credits = 0
	} else {
		CorrectedWumpus.Credits = UserWumpus.Credits
	}

	CorrectedWumpus.Color = UserWumpus.Color
	CorrectedWumpus.Sick = UserWumpus.Sick
	CorrectedWumpus.Sleeping = UserWumpus.Sleeping
	CorrectedWumpus.Left = UserWumpus.Left
	return CorrectedWumpus
}
