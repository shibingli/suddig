package configs

// Takes a string in and returns a proccesed string.
type StringFunc func(string) string

// Calculates the distance between the two strings.
type DistanceFunc func(string, string) float64

// Normalizes the score to a similarity percentage.
type ScoreFunc func(string, string, float64) float64

type Config struct {
	// A decimal number from 0-1
	MinScore float64
	// A function that gets ran on any input, great for normalizing strings
	StringFunc StringFunc
	// A function that takes in 2 strings and compares them returning the distance as a whole number
	DistanceFunc DistanceFunc
	// Takes the 2 inputs strings and the distance, and returns a score value
	ScoreFunc ScoreFunc
}
