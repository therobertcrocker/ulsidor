package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/fogleman/gg"
)

type block struct {
	x        float64
	y        float64
	width    float64
	height   float64
	fill     color.RGBA
	text     string
	fontPath string
	fontSize float64
	align    gg.Align
}

const (
	rewardBlockY    = 1740
	rewardBlockSize = 100
)

var (
	questColors = map[string]color.RGBA{
		// I want red yellow green and blue for the array
		"Hunt":        {0, 0, 0, 0},         // Red
		"Acquisition": {255, 224, 144, 255}, // Yellow
		"Whisper":     {0, 0, 0, 0},         // Green
		"Knowledge":   {0, 0, 0, 0},         // Blue
	}

	levelBlock = block{
		x:        1000,
		y:        80,
		width:    150,
		height:   200,
		fill:     color.RGBA{255, 165, 0, 150}, // Orange color
		text:     "5",
		fontSize: 150,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelHalbfett.ttf",
		align:    gg.AlignCenter,
	}
	titleBlock = block{
		x:        150,
		y:        825,
		width:    1050,
		height:   150,
		fill:     color.RGBA{255, 140, 0, 150}, // Different shade of orange
		text:     "Sleeping at Last",
		fontSize: 80,
		align:    gg.AlignCenter,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelHalbfett.ttf",
	}
	descriptionBlock = block{
		x:        150,
		y:        1025,
		width:    1050,
		height:   480,
		fill:     color.RGBA{255, 99, 71, 150}, // Different shade of orange
		text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ac aliquet nisi. Proin quam quam, scelerisque id semper non, imperdiet at tortor. Aliquam egestas neque ligula, nec placerat turpis malesuada id.",
		fontSize: 48,
		align:    gg.AlignLeft,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelBreit.ttf",
	}
	sourceBlock = block{
		x:        380,
		y:        1515,
		width:    600,
		height:   75,
		fill:     color.RGBA{255, 69, 0, 150}, // Different shade of orange
		text:     "Quest Source",
		fontSize: 36,
		align:    gg.AlignCenter,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelBreit.ttf",
	}
	goldBlock = block{
		x:        160,
		y:        rewardBlockY,
		width:    rewardBlockSize,
		height:   rewardBlockSize,
		fill:     color.RGBA{248, 148, 6, 150}, // Different shade of orange
		text:     "500",
		fontSize: 44,
		align:    gg.AlignCenter,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelHalbfett.ttf",
	}
	treasureBlock = block{
		x:        625,
		y:        rewardBlockY - 5,
		width:    rewardBlockSize,
		height:   rewardBlockSize,
		fill:     color.RGBA{242, 85, 96, 150}, // Different shade of orange
		text:     "7",
		fontSize: 48,
		align:    gg.AlignCenter,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelHalbfett.ttf",
	}
	reputationBlock = block{
		x:        1000,
		y:        rewardBlockY - 5,
		width:    rewardBlockSize,
		height:   rewardBlockSize,
		fill:     color.RGBA{242, 121, 53, 150}, // Different shade of orange
		text:     "+5",
		fontSize: 44,
		align:    gg.AlignCenter,
		fontPath: "../../../internal/modules/questforge/assets/fonts/DePixelHalbfett.ttf",
	}
)

func main() {
	// Load the background image
	backgroundPath := "../../../internal/modules/questforge/assets/acquisition_blank.png"
	backgroundImage, err := gg.LoadImage(backgroundPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the dimensions of the background image
	imageWidth := float64(backgroundImage.Bounds().Dx())
	imageHeight := float64(backgroundImage.Bounds().Dy())

	// Create a drawing context with the image dimensions
	dc := gg.NewContext(int(imageWidth), int(imageHeight))
	dc.DrawImage(backgroundImage, 0, 0)

	// Set color to bright green
	dc.SetRGBA(0, 255, 0, 255)
	// Draw a large point at the center
	dc.DrawPoint(imageWidth/2, imageHeight/2, 10)
	// Fill
	dc.Fill()

	// Draw coordinate points (every 100 pixels)
	//drawCoordinates(dc, 10)

	// Draw rectangles to represent components
	//drawRectangles(dc)

	// Draw text blocks
	drawTextBlocks(dc)

	// Save the visualization as an image
	if err := dc.SavePNG("quest_card_visualization_with_coordinates.png"); err != nil {
		log.Fatal(err)
	}
}

func DrawText(dc *gg.Context, text string, x, y, width, height float64, align gg.Align, fontSize float64, fontPath string, withShadow bool) {
	// Load font face
	dc.LoadFontFace(fontPath, fontSize)

	if withShadow {
		// Shadow details
		shadowOffsetX := -6.0
		shadowOffsetY := 6.0
		dc.SetRGBA(0, 0, 0, 0.5) // Semi-transparent black for shadow

		// Draw the shadow text
		if align == gg.AlignLeft {
			dc.DrawStringWrapped(text, x+shadowOffsetX, y+shadowOffsetY, 0, 0, width, 1.5, align)
		} else {
			// Centered text shadow
			dc.DrawStringAnchored(text, x+(width/2)+shadowOffsetX, y+(height/2)+shadowOffsetY, 0.5, 0.5)
		}

		// Reset color for the main text
		dc.SetRGBA(
			float64(questColors["Acquisition"].R)/255,
			float64(questColors["Acquisition"].G)/255,
			float64(questColors["Acquisition"].B)/255,
			float64(questColors["Acquisition"].A)/255) // Reset color
	}

	// Draw the main text
	if align == gg.AlignLeft {
		dc.DrawStringWrapped(text, x, y, 0, 0, width, 1.5, align)
	} else {
		// Centered text
		dc.DrawStringAnchored(text, x+(width/2), y+(height/2), 0.5, 0.5)
	}
}

func drawTextBlock(dc *gg.Context, b block) {
	// Draw the rectangle (background)
	//drawRect(dc, b.x, b.y, b.width, b.height, b.fill)

	// Use the generic function to draw text (with shadow in this case)
	DrawText(dc, b.text, b.x+10, b.y+10, b.width, b.height, b.align, b.fontSize, b.fontPath, true)
}

func drawTextBlocks(dc *gg.Context) {
	drawTextBlock(dc, levelBlock)
	drawTextBlock(dc, titleBlock)
	drawTextBlock(dc, descriptionBlock)
	drawTextBlock(dc, sourceBlock)
	drawTextBlock(dc, goldBlock)
	drawTextBlock(dc, treasureBlock)
	drawTextBlock(dc, reputationBlock)
}

func drawRect(dc *gg.Context, x, y, width, height float64, fill color.RGBA) {
	dc.SetRGBA(
		float64(fill.R)/255,
		float64(fill.G)/255,
		float64(fill.B)/255,
		float64(fill.A)/255)
	dc.DrawRectangle(x, y, width, height)
	dc.Fill()
}

func drawRectangles(dc *gg.Context) {
	drawRect(dc, levelBlock.x, levelBlock.y, levelBlock.width, levelBlock.height, levelBlock.fill)                               // Level
	drawRect(dc, titleBlock.x, titleBlock.y, titleBlock.width, titleBlock.height, titleBlock.fill)                               // Title
	drawRect(dc, descriptionBlock.x, descriptionBlock.y, descriptionBlock.width, descriptionBlock.height, descriptionBlock.fill) // Description
	drawRect(dc, sourceBlock.x, sourceBlock.y, sourceBlock.width, sourceBlock.height, sourceBlock.fill)                          // Source
	drawRect(dc, goldBlock.x, goldBlock.y, goldBlock.width, goldBlock.height, goldBlock.fill)                                    // Gold
	drawRect(dc, treasureBlock.x, treasureBlock.y, treasureBlock.width, treasureBlock.height, treasureBlock.fill)                // Treasure
	drawRect(dc, reputationBlock.x, reputationBlock.y, reputationBlock.width, reputationBlock.height, reputationBlock.fill)      // Reputation
}

func drawCoordinates(dc *gg.Context, interval float64) {
	// Set color to bright green
	gridColor := color.RGBA{0, 255, 0, 255}
	textColor := color.RGBA{0, 255, 0, 255} // green color for text
	gridSize := 5.0
	gridSpacing := 100.0

	imageWidth := float64(dc.Width())
	imageHeight := float64(dc.Height())

	for x := 0.0; x <= imageWidth; x += gridSpacing {
		for y := 0.0; y <= imageHeight; y += gridSpacing {
			drawGridSquare(dc, x, y, gridSize, gridColor)

			// Set color for text
			dc.SetRGBA(float64(textColor.R)/255.0, float64(textColor.G)/255.0, float64(textColor.B)/255.0, float64(textColor.A)/255.0)

			// Draw the coordinate values
			coordStr := fmt.Sprintf("(%d, %d)", int(x), int(y))
			dc.DrawString(coordStr, x+10, y) // Offsetting text by 10 pixels
		}
	}
}

func drawGridSquare(dc *gg.Context, x, y float64, size float64, fill color.RGBA) {
	dc.SetRGBA(float64(fill.R)/255.0, float64(fill.G)/255.0, float64(fill.B)/255.0, float64(fill.A)/255.0)
	dc.DrawRectangle(x-size/2, y-size/2, size, size)
	dc.Fill()
}
