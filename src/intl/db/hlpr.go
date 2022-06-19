package db

import "fmt"

func GenerateSelectStatement(sel map[string]map[string]string) string {
	stm := "SELECT"
	for key, props := range sel {
		for prop, alias := range props {
			if key == "" {
				stm = fmt.Sprintf("%s %s as %s,", stm, prop, alias)
			} else {
				stm = fmt.Sprintf("%s %s.%s as %s,", stm, key, prop, alias)
			}
		}
	}
	if stm == "SELECT" {
		return "SELECT *"
	} else {
		return stm[:len(stm)-1]
	}
}

func GenerateWhereAndStatement(
	varStartCount int,
	strQry map[string]string,
	strContainsQry map[string]string,
	nullableQry map[string]interface{},
) (string, []interface{}, int) {
	stm := "WHERE"
	var vals []interface{}

	varCount := varStartCount
	for key, qry := range strQry {
		if qry != "" {
			stm = fmt.Sprintf("%s %s = $%d AND", stm, key, varCount)
			vals = append(vals, qry)
			varCount++
		}
	}

	for key, qry := range strContainsQry {
		if qry != "" {
			stm = fmt.Sprintf("%s %s LIKE $%d AND", stm, key, varCount)
			vals = append(vals, fmt.Sprintf("%%%s%%", qry))
			varCount++
		}
	}

	for key, qry := range nullableQry {
		if qry != nil {
			stm = fmt.Sprintf("%s %s = $%d AND", stm, key, varCount)
			vals = append(vals, qry)
			varCount++
		}
	}

	if varCount == varStartCount {
		return "", vals, varCount
	} else {
		return stm[:len(stm)-4], vals, varCount
	}
}
