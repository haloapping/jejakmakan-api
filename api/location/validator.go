package location

func AddValidation(req AddReq) map[string][]string {
	validation := make(map[string][]string)

	if req.District == "" {
		validation["district"] = append(validation["district"], "cannot empty")
	}
	if req.City == "" {
		validation["city"] = append(validation["city"], "cannot empty")
	}
	if req.Province == "" {
		validation["province"] = append(validation["province"], "cannot empty")
	}
	if req.PostalCode == "" {
		validation["postalCode"] = append(validation["postalCode"], "cannot empty")
	}
	if req.Details == "" {
		validation["details"] = append(validation["details"], "cannot empty")
	}

	return validation
}
