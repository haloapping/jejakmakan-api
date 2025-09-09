package owner

func AddValidation(req AddReq) map[string][]string {
	validation := make(map[string][]string)

	if req.Images == "" {
		validation["images"] = append(validation["images"], "cannot empty")
	}

	if req.Name == "" {
		validation["name"] = append(validation["name"], "cannot empty")
	}

	return validation
}
