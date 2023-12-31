package apirestgin

var StoragePeople = []DomainPerson{
	{UUID: "61be616a-176e-4f41-8730-53a8366981d5", FirstName: "Chelcey", Email: "mimi_solarjub@cultural.jnt", Ege: 16},
	{UUID: "918e414a-b031-4706-bf74-bfaea011aaba", FirstName: "Alta", Email: "zedrick_pickleshe@titanium.gek", Ege: 30},
	{UUID: "e9844c1f-5669-4929-95c8-8410ac52e53d", FirstName: "Shawanda", Email: "shareka_pitts8@floral.md", Ege: 40},
	{UUID: "2f1d8415-e141-4569-b34b-0aba2d968178", FirstName: "Ryna", Email: "noura_vaillancourtwl@wind.jux", Ege: 20},
	{UUID: "bab4ef60-2044-4192-aaa7-49a50ae77317", FirstName: "Paolo", Email: "aubrie_burchfieldd@marking.fv", Ege: 23},
}

func getAll() *[]DomainPerson {
	return &StoragePeople
}

func add(p DomainPerson) {
	StoragePeople = append(StoragePeople, p)
}

func delete(s *[]DomainPerson, uuid string) Message {

	var slice []DomainPerson = *s
	for index, a := range slice {
		if a.UUID == uuid {
			*s = append(slice[:index], slice[index+1:]...)
			return Message{Msg: "Delete uuid " + uuid}
		}
	}

	return Message{Msg: "Not found uuid " + uuid}
}

func getPersonByUUID(s *[]DomainPerson, uuid string) DomainPerson {
	for _, a := range StoragePeople {
		if a.UUID == uuid {
			return a
		}
	}

	return DomainPerson{}
}
