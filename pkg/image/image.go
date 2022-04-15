package image

import (
	"bytes"
	"fmt"
	"github.com/Sh1karno/image_generator/pkg/colors"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"

	"github.com/golang/freetype/truetype"
)

const (
	// Default parameters ...
	imgColorDefault         = "E5E5E5"
	msgColorDefault         = "AAAAAA"
	imgWDefault             = 300
	imgHDefault             = 300
	fontSizeDefault         = 0
	dpiDefault      float64 = 72

	fontFileDefault = "wqy-zenhei.ttf"
)

type Placeholder struct {
	Width    int64
	Height   int64
	Color    string
	Text     string
	FontSize int64
	Label    Label
}

type Label struct {
	Text     string
	FontSize int64
	Color    string
}

// Generate - соберёт картинку по нужным размерам, цветом и текстом.
func (p *Placeholder) Generate() (*bytes.Buffer, error) {
	// Если есть размеры и нет требований по Тексту - соберём Текст по умолчанию.
	if ((p.Width > 0 || p.Height > 0) && p.Text == "") || p.Text == "" {
		p.Text = fmt.Sprintf("%d x %d", p.Width, p.Height)
	}
	// Если нет требований по размеру шрифта - подберём его исходя из размеров картинки.
	if p.FontSize == 0 {
		p.FontSize = p.Width / 10
		if p.Height < p.Width {
			p.FontSize = p.Height / 5
		}
	}
	// Переведём цвет из строки в color.RGBA.
	clr, err := colors.ToRGBA(p.Color)
	if err != nil {
		return nil, err
	}
	// Создадим in-memory картинку с нужными размерами.
	m := image.NewRGBA(image.Rect(0, 0, int(p.Width), int(p.Height)))
	// Отрисуем картинку:
	// - по размерам (Bounds)
	// - и с цветом (Uniform - обёртка над color.Color c Image функциями)
	// - исходя из точки (Point), как базовой картинки
	// - заполним цветом нашу Uniform (draw.Src)
	draw.Draw(m, m.Bounds(), image.NewUniform(clr), image.Point{}, draw.Src)
	// Добавим текст в картинку.
	if err = p.addLabel(m); err != nil {
		return nil, err
	}
	var img image.Image = m
	// Выделим память под нашу данные (байты картинки).
	buffer := &bytes.Buffer{}
	// Закодируем картинку в нашу аллоцированную память.
	err = jpeg.Encode(buffer, img, nil)

	return buffer, err
}

// drawLabel - добавит текст на картинку.
func (p *Placeholder) addLabel(m *image.RGBA) error {
	// Разберём цвет текста из строки в RGBA.
	clr, err := colors.ToRGBA(p.Label.Color)
	if err != nil {
		return err
	}
	// Получим шрифт (должен работать и с латиницей и с кириллицей).
	fontBytes, err := ioutil.ReadFile(fontFileDefault)
	if err != nil {
		return err
	}
	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		return err
	}
	// Подготовим Drawer для отрисовки текста на картинке.
	d := &font.Drawer{
		Dst: m,
		Src: image.NewUniform(clr),
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size:    float64(p.FontSize),
			DPI:     dpiDefault,
			Hinting: font.HintingNone,
		}),
	}
	// Зададим базовую линию.
	d.Dot = fixed.Point26_6{
		X: (fixed.I(int(p.Width)) - d.MeasureString(p.Text)) / 2,
		Y: fixed.I(int((p.Height+p.FontSize)/2 - 12)),
	}
	// Непосредственно отрисовка текста в нашу RGBA картинку.
	d.DrawString(p.Text)

	return nil
}
