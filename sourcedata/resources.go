package sourcedata

import (
	"image"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha-resources/helper"
	"github.com/wenlng/go-captcha-resources/sourcedata/chars"
)

type GraphImage struct {
	OverlayImage image.Image
	ShadowImage  image.Image
	MaskImage    image.Image
}

// ResourcesFile .
type ResourcesFile struct {
}

func NewResourcesFile() *ResourcesFile {
	return &ResourcesFile{}
}

func (sf *ResourcesFile) GetChineseChars() []string {
	return chars.GetChineseChars()
}

func (sf *ResourcesFile) GetAlphaChars() []string {
	return chars.GetAlphaChars()
}

func (sf *ResourcesFile) GetFonts() (fonts []*truetype.Font, err error) {
	font, err := helper.LoadFont("./resources/fonts/fzshengsksjw_cu/font.ttf")
	if err != nil {
		return fonts, err
	}

	return []*truetype.Font{
		font,
	}, nil
}

func (sf *ResourcesFile) GetImages() ([]image.Image, error) {
	var images = make([]image.Image, 0, 12)
	var img image.Image
	var err error

	img, err = helper.LoadJpeg("resources/images/image-1/image.jpg")
	if err != nil {
		return images, err
	}
	images[0] = img

	img, err = helper.LoadJpeg("resources/images/image-2/image.jpg")
	if err != nil {
		return images, err
	}
	images[1] = img

	img, err = helper.LoadJpeg("resources/images/image-3/image.jpg")
	if err != nil {
		return images, err
	}
	images[2] = img

	img, err = helper.LoadJpeg("resources/images/image-4/image.jpg")
	if err != nil {
		return images, err
	}
	images[3] = img

	img, err = helper.LoadJpeg("resources/images/image-5/image.jpg")
	if err != nil {
		return images, err
	}
	images[4] = img

	return images, nil
}

func (sf *ResourcesFile) GetImagesV2() ([]image.Image, error) {
	var images = make([]image.Image, 0, 12)
	var img image.Image
	var err error

	img, err = helper.LoadJpeg("resources/images/image-v2-1/image.jpg")
	if err != nil {
		return images, err
	}
	images[0] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-2/image.jpg")
	if err != nil {
		return images, err
	}
	images[1] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-3/image.jpg")
	if err != nil {
		return images, err
	}
	images[2] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-4/image.jpg")
	if err != nil {
		return images, err
	}
	images[3] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-5/image.jpg")
	if err != nil {
		return images, err
	}
	images[4] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-6/image.jpg")
	if err != nil {
		return images, err
	}
	images[5] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-7/image.jpg")
	if err != nil {
		return images, err
	}
	images[6] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-8/image.jpg")
	if err != nil {
		return images, err
	}
	images[7] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-9/image.jpg")
	if err != nil {
		return images, err
	}
	images[8] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-10/image.jpg")
	if err != nil {
		return images, err
	}
	images[9] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-11/image.jpg")
	if err != nil {
		return images, err
	}
	images[10] = img

	img, err = helper.LoadJpeg("resources/images/image-v2-12/image.jpg")
	if err != nil {
		return images, err
	}
	images[11] = img
	return images, nil
}

func (sf *ResourcesFile) GetThumbImages() ([]image.Image, error) {
	var images = make([]image.Image, 0, 5)
	var img image.Image
	var err error

	img, err = helper.LoadJpeg("resources/thumbs/thumb-1/thumb.jpg")
	if err != nil {
		return images, err
	}
	images[0] = img

	img, err = helper.LoadJpeg("resources/thumbs/thumb-2/thumb.jpg")
	if err != nil {
		return images, err
	}
	images[1] = img

	img, err = helper.LoadJpeg("resources/thumbs/thumb-3/thumb.jpg")
	if err != nil {
		return images, err
	}
	images[2] = img

	img, err = helper.LoadJpeg("resources/thumbs/thumb-4/thumb.jpg")
	if err != nil {
		return images, err
	}
	images[3] = img

	img, err = helper.LoadJpeg("resources/thumbs/thumb-5/thumb.jpg")
	if err != nil {
		return images, err
	}
	images[4] = img

	return images, nil
}

func (sf *ResourcesFile) GetShapeMaps() (maps map[string]image.Image, err error) {
	shapeImage1, err := helper.LoadPNG("resources/shapes/shape-1/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage2, err := helper.LoadPNG("resources/shapes/shape-2/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage3, err := helper.LoadPNG("resources/shapes/shape-3/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage4, err := helper.LoadPNG("resources/shapes/shape-4/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage5, err := helper.LoadPNG("resources/shapes/shape-5/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage6, err := helper.LoadPNG("resources/shapes/shape-6/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage7, err := helper.LoadPNG("resources/shapes/shape-7/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage8, err := helper.LoadPNG("resources/shapes/shape-8/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage9, err := helper.LoadPNG("resources/shapes/shape-9/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage10, err := helper.LoadPNG("resources/shapes/shape-10/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage11, err := helper.LoadPNG("resources/shapes/shape-11/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage12, err := helper.LoadPNG("resources/shapes/shape-12/shape.png")
	if err != nil {
		return maps, err
	}

	shapeImage13, err := helper.LoadPNG("resources/shapes/shape-13/shape.png")
	if err != nil {
		return maps, err
	}

	return map[string]image.Image{
		"shape1":  shapeImage1,
		"shape2":  shapeImage2,
		"shape3":  shapeImage3,
		"shape4":  shapeImage4,
		"shape5":  shapeImage5,
		"shape6":  shapeImage6,
		"shape7":  shapeImage7,
		"shape8":  shapeImage8,
		"shape9":  shapeImage9,
		"shape10": shapeImage10,
		"shape11": shapeImage11,
		"shape12": shapeImage12,
		"shape13": shapeImage13,
	}, nil
}

func (sf *ResourcesFile) GetTileImages() (images []*GraphImage, err error) {
	tileImage1, err := helper.LoadPNG("resources/tiles/tile-1/tile.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage1, err := helper.LoadPNG("resources/tiles/tile-1/tile-shadow.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage1, err := helper.LoadPNG("resources/tiles/tile-1/tile-mask.png")
	if err != nil {
		return images, nil
	}

	tileImage2, err := helper.LoadPNG("resources/tiles/tile-2/tile.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage2, err := helper.LoadPNG("resources/tiles/tile-2/tile-shadow.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage2, err := helper.LoadPNG("resources/tiles/tile-2/tile-mask.png")
	if err != nil {
		return images, nil
	}

	tileImage3, err := helper.LoadPNG("resources/tiles/tile-2/tile.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage3, err := helper.LoadPNG("resources/tiles/tile-2/tile-shadow.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage3, err := helper.LoadPNG("resources/tiles/tile-2/tile-mask.png")
	if err != nil {
		return images, nil
	}

	tileImage4, err := helper.LoadPNG("resources/tiles/tile-2/tile.png")
	if err != nil {
		return images, nil
	}
	tileShadowImage4, err := helper.LoadPNG("resources/tiles/tile-2/tile-shadow.png")
	if err != nil {
		return images, nil
	}
	tileMaskImage4, err := helper.LoadPNG("resources/tiles/tile-2/tile-mask.png")
	if err != nil {
		return images, nil
	}

	return []*GraphImage{
		{
			OverlayImage: tileImage1,
			ShadowImage:  tileShadowImage1,
			MaskImage:    tileMaskImage1,
		},
		{
			OverlayImage: tileImage2,
			ShadowImage:  tileShadowImage2,
			MaskImage:    tileMaskImage2,
		},
		{
			OverlayImage: tileImage3,
			ShadowImage:  tileShadowImage3,
			MaskImage:    tileMaskImage3,
		},
		{
			OverlayImage: tileImage4,
			ShadowImage:  tileShadowImage4,
			MaskImage:    tileMaskImage4,
		},
	}, nil
}
