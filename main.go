package main

type Place struct {
	id         string
	Name       string
	Doors      []*Door
	PlaceItems []*PlaceItem
	PlacesToGo []string
}

type Door struct {
	destination string
	IsLocked    bool
}

type PlaceItem struct {
	Name  string
	Items []*Item
}

type Item struct {
	Name string
}

type Player struct {
	Name string
}

func main() {
	//******************
}

func initGame() {

	var tee Item = Item{
		Name: "чай",
	}
	var tableKitchen PlaceItem = PlaceItem{
		Name:  "на столе",
		Items: []*Item{&tee},
	}
	var kitchen Place = Place{
		id:         "kitchen",
		Name:       "кухня",
		PlaceItems: []*PlaceItem{&tableKitchen},
		PlacesToGo: []string{"hallway"},
	}
	var hallway Place = Place{
		id:         "hallway",
		Name:       "коридор",
		PlacesToGo: []string{"kitchen", "outside", "myRoom"},
	}
	var notes Item = Item{
		Name: "конспекты",
	}
	var keys Item = Item{
		Name: "ключи",
	}
	var backpack Item = Item{
		Name: "рюкзак",
	}
	var tableMyRoom PlaceItem = PlaceItem{
		Name:  "на столе",
		Items: []*Item{&notes, &keys},
	}
	var chairMyRoom PlaceItem = PlaceItem{
		Name:  "на стуле",
		Items: []*Item{&backpack},
	}
	var myRoom Place = Place{
		id:         "myRoom",
		Name:       "комната",
		PlaceItems: []*PlaceItem{&chairMyRoom, &tableMyRoom},
		PlacesToGo: []string{"hallway"},
	}
	var lockedDoor Door = Door{
		destination: "outside",
		IsLocked:    true,
	}
	var outside Place = Place{
		id:         "outside",
		Name:       "улица",
		PlacesToGo: []string{"hallway"},
		Doors:      []*Door{&lockedDoor},
	}
}

func handleCommand(command string) string {
	/*
		данная функция принимает команду от "пользователя"
		и наверняка вызывает какой-то другой метод или функцию у "мира" - списка комнат
	*/
	return "not implemented"
}
