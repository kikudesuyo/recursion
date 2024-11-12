package list

func FireEmployees(employees []string, unemployed []string) []string {
	var fireEmployees []string
	for _, employee := range employees {
		flag := true
		for _, unemployee := range unemployed {
			if employee == unemployee {
				flag = false
				break
			}
		}
		if flag {
			fireEmployees = append(fireEmployees, employee)
		}
	}
	return fireEmployees
}
