package thoughtfactory

type EmptyThought struct {
	Thought
}

func (t *EmptyThought) Process() string {
	return "I don't really understand..."
}