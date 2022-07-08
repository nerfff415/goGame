package main

type Place struct {
	Name       string
	Doors      []*Door
	PlaceItems []*PlaceItem
	PlacesToGo []*Place
}

type Door struct {
	destination Place
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
		Name:       "Kitchen",
		PlaceItems: []*PlaceItem{&tableKitchen},
		PlacesToGo: []*Place{&hallway},
	}
	var hallway Place = Place{
		Name:       "коридор",
		PlacesToGo: []*Place{&kitchen, &myRoom, &outside},
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
		Name:       "комната",
		PlaceItems: []*PlaceItem{&chairMyRoom, &tableMyRoom},
		PlacesToGo: []*Place{&hallway},
	}
	var lockedDoor Door = Door{
		destination: outside,
		IsLocked:    true,
	}
	var outside Place = Place{
		Name:       "улица",
		PlacesToGo: []*Place{&hallway},
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
