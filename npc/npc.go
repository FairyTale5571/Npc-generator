package npc

import (
	"Npc-generator/store"
	"fmt"
)

func getTemplate() map[string]string {
	return map[string]string{
		"class": "string",
		"object": "string",
		"position": "array", // [posWorld,vectorDir,vectorUp]
		"animations": "array",
		"action": "string",
		"name": "string",
		
		"loadout": "class",
		"mainWeapon": "string",
		"handWeapon": "string",
		"launcherWeapon": "string",
		"face": "string",
		"headgear": "string",
		"uniform": "string",
		"vest": "string",
		"backpack": "string",
	}
}

func CreateStruct() []store.Unit {
	var units []store.Unit

	return units
}

func Pos(mode string, unit store.Unit) []float32 {
	var res []float32
	switch mode {
	case "pos":
		res = []float32{unit.ObjectPos.Pos.X,unit.ObjectPos.Pos.Y,unit.ObjectPos.Pos.Z}
	case "dir":
		res = []float32{unit.ObjectPos.VectorDir.X,unit.ObjectPos.VectorDir.Y,unit.ObjectPos.VectorDir.Z}
	case "up":
		res = []float32{unit.ObjectPos.VectorUp.X,unit.ObjectPos.VectorUp.Y,unit.ObjectPos.VectorUp.Z}
	}
	return res
}

func CreateClass() {
	class := "class ConfigNpc\n{"
	class += "\n\t"
	for _, elem := range CreateStruct() {
		class += "class " + elem.VariableName
		class += "\t{"
		class += fmt.Sprintf(`objectClassname = "%s";\n`,elem.ObjectClassName)
		class += fmt.Sprintf(`objPosition[] = {{%f,%f,%f},{%f,%f,%f},{%f,%f,%f}};\n`,Pos("pos",elem)[0],Pos("pos",elem)[1],Pos("pos",elem)[2],
			Pos("dir",elem)[0],Pos("dir",elem)[1],Pos("dir",elem)[2],
			Pos("up",elem)[0],Pos("up",elem)[1],Pos("up",elem)[2])
		class += fmt.Sprintf(`animations[] = {%s};\n`,elem.Animations)
		class += "class Equipment\n{"
		class += "\t"
		class += fmt.Sprintf(`face = "%s";\n`,elem.Equipment.Unit.Face)
		class += fmt.Sprintf(`headgear = "%s";\n`,elem.Equipment.Unit.Headgear)
		class += fmt.Sprintf(`glass = "%s";\n`,elem.Equipment.Unit.Glass)
		class += fmt.Sprintf(`uniform = "%s";\n`,elem.Equipment.Unit.Uniform)
		class += fmt.Sprintf(`vest = "%s";\n`,elem.Equipment.Unit.Vest)
		class += fmt.Sprintf(`backpack = "%s";\n`,elem.Equipment.Unit.Backpack)

		class += fmt.Sprintf(`rifle = "%s";\n`,elem.Equipment.Weapons.Rifle)
		class += fmt.Sprintf(`pistol = "%s";\n`,elem.Equipment.Weapons.Pistol)
		class += fmt.Sprintf(`launcher = "%s";\n`,elem.Equipment.Weapons.Launcher)
		class += fmt.Sprintf("};\n")

		class += "};\n"

	}

	class += "};\n"

}

