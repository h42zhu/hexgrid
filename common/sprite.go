package common

import (
	"image"
	"strings"

	// this is for init image decode
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/faiface/pixel"
)

// ImageAsset is a map containing game assets
type ImageAsset struct {
	SpriteMap map[string]*pixel.Sprite
}

// NewImageAsset returns an empty ImageAsset
func NewImageAsset() *ImageAsset {
	return &ImageAsset{
		SpriteMap: map[string]*pixel.Sprite{},
	}
}

// LoadPictures ..
func LoadPictures(path []string) (map[string]*pixel.Sprite, error) {
	sm := map[string]*pixel.Sprite{}

	for _, f := range path {
		pic, err := loadPicture(f)
		if err != nil {
			return nil, err
		}
		basename := filepath.Base(f)
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		sprite := pixel.NewSprite(pic, pic.Bounds())
		sm[name] = sprite
	}
	return sm, nil

}

func loadPicture(path string) (pixel.Picture, error) {
	ff, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(ff)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
