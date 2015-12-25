package types
import "fmt"

type EmptyThought struct {
	SentenceThought
}

func (t *EmptyThought) Process() {
	fmt.Println("I don't really understand...")
}