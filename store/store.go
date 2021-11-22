package store


type Unit struct {
	VariableName 		string
	ObjectClassName 	string
	ObjectPos struct{
		Pos struct{
			X float32
			Y float32
			Z float32
		}
		VectorDir struct{
			X float32
			Y float32
			Z float32
		}
		VectorUp struct{
			X float32
			Y float32
			Z float32
		}
	}
	Animations []string
	Equipment struct{
		Weapons struct{
			Rifle 		string
			Pistol 		string
			Launcher 	string
		}
		Unit struct{
			Face 		string
			Headgear 	string
			Glass 		string
			Uniform 	string
			Vest 		string
			Backpack 	string
		}
	}
}