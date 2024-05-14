package rpgGo

type Story struct {
	ID     int
	Choice string
	Story  string
}

var Intro = Story{
	ID:     1,
	Choice: "Start",
	Story:  "You are sitting in a tavern, going through your item kit to make sure you have everything to start your journey.",
}

var Stay = Story{
	ID:     2,
	Choice: "Stay",
	Story:  "As you meticulously check your potions and sharpen your dagger, a cloaked figure slides into the seat beside you. \nTheir face is hidden in the shadows of their hood, but you catch a glint of piercing blue eyes.\nYou look like someone who's about to embark on an adventure, \nthe stranger says in a low, raspy voice. \nI have a proposition for you.",
}
