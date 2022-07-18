package functional_go

import (
	"strconv"
	"testing"
)

func TestMapDoublingIntegers(t *testing.T) {
	function_which_doubles_an_integer := func(number int) int {
		return number * 2
	}

	integers_to_double := []int{1, 2, 3}

	doubled_integers := Map(
		integers_to_double,
		function_which_doubles_an_integer,
	)

	if len(doubled_integers) != 3 {
		t.Fatalf(`len(%q) != 3`, doubled_integers)
	}
	if doubled_integers[0] != 2 ||
		doubled_integers[1] != 4 ||
		doubled_integers[2] != 6 {
		t.Fatalf(`%q != []int{2, 4, 6}`, doubled_integers)
	}
}

func TestMapConcatenatingStrings(t *testing.T) {
	function_which_concatenates_world := func(words string) string {
		return words + " world"
	}

	greetings := []string{"hello", "goodbye"}

	world_greetings := Map(greetings, function_which_concatenates_world)

	if len(world_greetings) != 2 {
		t.Fatalf(`len(%q) != 2`, world_greetings)
	}
	if world_greetings[0] != "hello world" ||
		world_greetings[1] != "goodbye world" {
		t.Fatalf(`%q != []string{'hello world', 'goodbye world'}`, world_greetings)
	}
}

func TestMapIntegersToStrings(t *testing.T) {
	numbers := []int{1, 2}

	strings := Map(
		numbers,
		func(number int) string {
			return strconv.Itoa(number)
		},
	)

	if strings[0] != "1" || strings[1] != "2" {
		t.Fatalf(`%q != []string{"1", "2"}`, strings)
	}
}

func TestFilterRemovesEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	even_numbers := Filter(
		numbers,
		func(number int) bool { return number%2 == 0 },
	)

	if len(even_numbers) != 2 ||
		even_numbers[0] != 2 ||
		even_numbers[1] != 4 {
		t.Fatalf(`%q != []int{2, 4}`, even_numbers)
	}
}

func TestFilterRemovesEmptyStrings(t *testing.T) {
	strings := []string{"hello", "", "world"}

	words := Filter(
		strings,
		func(word string) bool { return len(word) > 0 },
	)

	if len(words) != 2 ||
		words[0] != "hello" ||
		words[1] != "world" {
		t.Fatalf(`%q != []string{"hello", "world"}`, words)
	}
}

func TestReduceSumsIntegers(t *testing.T) {
	numbers := []int{1, 2, 3}

	sum := Reduce(numbers, 0, func(accumulator int, nextNumber int) int { return accumulator + nextNumber })

	if sum != 6 {
		t.Fatalf(`%q != 6`, sum)
	}
}

func TestReduceOnlyKeepsEvenNumbers(t *testing.T) {
	numbers := []int{1, 2, 3}

	even_numbers := Reduce(numbers, []int{}, func(accumulator []int, nextNumber int) []int {
		if nextNumber%2 == 0 {
			return append(accumulator, nextNumber)
		} else {
			return accumulator
		}
	})

	if len(even_numbers) != 1 || even_numbers[0] != 2 {
		t.Fatalf(`%q != []int{2}`, even_numbers)
	}
}

func TestReduceRemovesEmptyStrings(t *testing.T) {
	strings := []string{"hello", "", "world"}

	words := Reduce(strings, []string{}, func(accumulator []string, nextString string) []string {
		if len(nextString) > 0 {
			return append(accumulator, nextString)
		} else {
			return accumulator
		}
	})

	if len(words) != 2 || words[0] != "hello" || words[1] != "world" {
		t.Fatalf(`%q != []string{"hello", "world"}`, words)
	}
}
