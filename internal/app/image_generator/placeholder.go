package image_generator

import (
	"context"
	"log"

	desc "github.com/Sh1karno/image_generator/pkg/api/github.com/Sh1karno/image_generator"
	"github.com/Sh1karno/image_generator/pkg/image"
)

func (i *Implementation) GetPlaceholder(ctx context.Context, req *desc.GetPlaceholderRequest) (*desc.GetPlaceholderResponse, error) {
	log.Println("come request")
	label := image.Label{Text: req.Text, FontSize: req.FoundSize, Color: req.TextColor}
	// Соберём структуру Картинки с нужными полями - высота, ширина, цвет и текст
	placeholder := image.Placeholder{Width: req.Width, Height: req.Height, Color: req.BackgroundColor, Label: label}

	// Сгенерим нашу картинку с текстом
	buff, err := placeholder.Generate()
	if err != nil {
		log.Println(err)
	}

	var img []byte
	if buff != nil {
		img = buff.Bytes()

	}

	resp := &desc.GetPlaceholderResponse{Image: img}

	return resp, nil
}
