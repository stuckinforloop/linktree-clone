package serializer

import (
	"encoding/json"
	"fmt"

	"github.com/neel229/linktree-clone/internal/shortener"
)

type Redirect struct {
}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, &redirect); err != nil {
		return nil, fmt.Errorf("error unmarshaling data: %v", err)
	}
	return redirect, nil
}

func (r *Redirect) Encode(redirect *shortener.Redirect) ([]byte, error) {
	data, err := json.Marshal(redirect)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %v", err)
	}
	return data, nil
}
