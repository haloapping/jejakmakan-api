package food

func AddValidation(req AddReq) map[string][]string {
	validation := make(map[string][]string)

	if req.UserId == "" {
		validation["userId"] = append(validation["userId"], "cannot empty")
	}

	if req.OwnerId == "" {
		validation["ownerId"] = append(validation["ownerId"], "cannot empty")
	}

	if req.LocationId == "" {
		validation["locationId"] = append(validation["locationId"], "cannot empty")
	}

	if req.Images == "" {
		validation["images"] = append(validation["iamges"], "cannot empty")
	}

	if req.Name == "" {
		validation["name"] = append(validation["name"], "cannot empty")
	}

	if req.Description == "" {
		validation["description"] = append(validation["description"], "cannot empty")
	}

	if req.Price < 0 {
		validation["price"] = append(validation["price"], "cannot negative value")
	}

	if req.Review == "" {
		validation["review"] = append(validation["review"], "cannot empty")
	}

	return validation
}
