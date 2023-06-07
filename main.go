package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Kagami/go-face"
)

const dataDir = "./data"

func main() {
	fmt.Println("Go Facer ")

	rec, err := face.NewRecognizer(filepath.Join(dataDir, "models"))
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}

	fmt.Println("Recognizer Initialized")
	//count number of faces in the image
	image := filepath.Join(dataDir, "images", "avengers-02.jpeg")
	faces, err := rec.RecognizeFile(image)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	fmt.Println("Total Number of Faces in Image: ", len(faces))

	var samples []face.Descriptor //array samples of type face.Descriptor
	var avengers []int32          //array of indexes of avengers
	for i, f := range faces {
		samples = append(samples, f.Descriptor)
		// Each face is unique on that image so goes to its own category.
		avengers = append(avengers, int32(i))
	}
	// Name the categories, i.e. people on the image.
	labels := []string{
		"Tony Stark",
		"Dr Strange",
		"Bruce Banner",
		"Hulk",
	}
	// Pass samples to the recognizer.
	rec.SetSamples(samples, avengers)

	// classification of some not yet known image.
	one := filepath.Join(dataDir, "images", "tony-stark.jpg")
	tonyStark, err := rec.RecognizeSingleFile(one)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if tonyStark == nil { //image is empty
		log.Fatalf("Not a single face on the image or image is empty")
	}
	ID := rec.Classify(tonyStark.Descriptor)
	if ID < 0 { //don't exists in our refernce images data
		log.Fatalf("Can't classify the image")
	}

	fmt.Print("Face ID : ")
	fmt.Print(ID)
	fmt.Print(" Classified as : ")
	fmt.Println(labels[ID])

	one = filepath.Join(dataDir, "images", "dr-strange.jpg")
	drStrange, err := rec.RecognizeSingleFile(one)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if drStrange == nil {
		log.Fatalf("Not a single face on the image or image is empty")
	}
	ID = rec.Classify(drStrange.Descriptor)
	if ID < 0 {
		log.Fatalf("Can't classify the image")
	}

	fmt.Print("Face ID : ")
	fmt.Print(ID)
	fmt.Print(" Classified as : ")
	fmt.Println(labels[ID])

	one = filepath.Join(dataDir, "images", "wong.jpg")
	wong, err := rec.RecognizeSingleFile(one)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if wong == nil {
		log.Fatalf("Not a single face on the image or image is empty")
	}
	ID = rec.Classify(wong.Descriptor)
	if ID < 0 {
		log.Fatalf("Can't classify the image")
	}
	fmt.Print("Face ID : ")
	fmt.Print(ID)
	fmt.Print(" Classified as : ")
	fmt.Println(labels[ID])
}
