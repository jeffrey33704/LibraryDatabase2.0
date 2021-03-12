package API

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
	"gocv.io/x/gocv"
	"image"
	_ "image/png"
	"os"
	"projects/LibraryDatabase2.0/rand"
)

var SaveFile string

// Creates a file with a random string name in order to capture the scan as an image.
func WebcamInput() {
	SaveFile = rand.String(5) + ".png"
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device %v\n", deviceID)
		return
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}

	if img.Empty() {
		fmt.Printf("no image on device %v\n", deviceID)
		return
	}

	gocv.IMWrite(SaveFile, img)
}

// Activates WebcamInput() to obtain images of barcodes for processing.
// If an error occurs in processing (if an EAN13 is not obtainable, either due to
// an illegible scan or a barcode that is not EAN13), ProcessImage() returns an empty
// string and a false boolean value.  If successful, ProcessImage() returns the barcode
// number and a true value.
// Furthermore, whether ProcessImage() is successful, or not, the SaveFile created by
// WebcamInput() is deleted, in order to save resources.
func ProcessImage() (string, bool) {
	WebcamInput()
	open, _ := os.Open(SaveFile)
	img, _, _ := image.Decode(open)
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	bcReader := oned.NewEAN13Reader()

	result, err := bcReader.DecodeWithoutHints(bmp)
	resultString := fmt.Sprintf("%v", result)
	if err != nil {
		os.Remove(SaveFile)
		return "", false
	} else {
		os.Remove(SaveFile)
		return resultString, true
	}
}

// Calls forth ProcessImage() to implement the acquisition of an ISBN13 from a barcode.
// If a valid scan occurs (meaning, if a readable scan is processed) an ISBN is exported.
// If ProcessImage() returns a false validity, a loop continues to call ProcessImage()
// until a true value is obtained.
func ExportISBN() string {
	var exportIsbn string
	fmt.Println("Scanning barcode...")
	for {
		result, validity := ProcessImage()
		if validity {
			fmt.Println("Barcode found!")
			exportIsbn = result
			break
		}
	}
	return exportIsbn
}