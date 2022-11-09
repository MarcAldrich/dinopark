package dinopark

type Error struct {
	error
	code    int
	message string
}
