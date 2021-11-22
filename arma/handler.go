package arma

import "Npc-generator/npc"

func ArgsHandle(funcName string, args []string) string {
	switch funcName {
	case "create":
		if len(args) < 0 {
			return "Not enough lens"
		}
		npc.CreateClass()
	}
	return ""
}