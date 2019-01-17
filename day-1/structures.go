package main

import "fmt"

type Habit struct {
	name string
	description string
	id int
	category string
}

func createHabit( id int, name string, description string, category string ) Habit {
	var h Habit
	h.id = id;
	h.name = name;
	h.description = description;
	h.category = category;
	return h;
}

func printHabit( h Habit ) {
	fmt.Printf( "Habit name : %s\n", h.name);
	fmt.Printf( "Habit description : %s\n", h.description);
	fmt.Printf( "Habit category : %s\n", h.category);
	fmt.Printf( "Habit id : %d\n", h.id); 
}

func main() {
	var habit Habit
	habit = createHabit( 1, "Morning Walk", "Walking daily for about 30-40 minutes", "Good" )
	printHabit( habit );
}