# Image Manuplation
This Go code demonstrates the use of constants, basic image operations, and a function from the `image` package. Let's go through each part of the code:

```go
package main

import (
	"fmt"
	"image"
	"image/jpeg" // use "image/png" for png files
	"os"
)

func main() {
	// Change the file path to the location of the file
	filepath := "/Users/gnbhavithran/Non_python_programs/Go_lang/Test.jpg"
	file, err := os.Open(filepath)

  //Displays error if file not found or unaccessable
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the image file  
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	// Create a new file to save the reshaped image 
	outputFile, err := os.Create("/Users/gnbhavithran/Non_python_programs/Go_lang/output1.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputFile.Close()

	scale := 0.1  //Change the scale to the required size
	newWidth := int(float64(img.Bounds().Dx()) * scale)
	newHeight := int(float64(img.Bounds().Dy()) * scale)

	// Create a new RGBA image with the specified dimensions
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Scale the original image to fit within the new dimensions
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			// Map the coordinates from the new image to the original image
			origX := x * width / newWidth
			origY := y * height / newHeight

			// Set the color of the pixel in the new image to the color of the corresponding pixel in the original image
			newImg.Set(x, y, img.At(origX, origY))
		}
	}

  // Encoding the values into JPG 
	err = jpeg.Encode(outputFile, newImg, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Image resized to %d%% and saved successfully!", int(scale * 100))
}
```

When you run this Go program, it will output the values of the constants and the result of the mathematical operations to the console.

The output should be something like:

```
Image resized to 10% of Orginal size and saved successfully!
```

### Explanation:

1. `img.Bounds()`: Declares the boundary of the `Image`(img). `Dx` and `Dy` maximum value in X and Y axis.

2. `image.NewRGBA()`: An RGBA image is a standard image type where each pixel has four components: red, green, blue, and alpha (transparency) with boundaries given in "image.Rect()". 

3. `image.Rect()`: representing the rectangle defined by the given coordinates.

4. `image.Decode()`: decode image data from an io.Reader into an image.Image object.
