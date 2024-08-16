package ex01

import "fmt"

type UnknownPlant struct {
    FlowerType  string
    LeafType    string
    Color       int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
    FlowerColor int
    LeafType    string
    Height      int `unit:"inches"`
}

func DescripePlant(plant any) {
    var plant1 UnknownPlant
    var plant2 AnotherUnknownPlant
    var ok1 = false
    var ok2 = false

    plant1, ok1 = plant.(UnknownPlant)
    plant2, ok2 = plant.(AnotherUnknownPlant)
    if !ok1 && !ok2 {
        return
    }

    plant1.Color.Tag

    switch ok1 {
    case true:
        fmt.Printf("FlowerType:%v", plant1.FlowerType)
        fmt.Printf("LeafType:%v", plant1.LeafType)
        fmt.Printf("Color:%v", plant1.Color)
    case false:
        fmt.Printf("FlowerColor:%v", plant2.FlowerColor)
        fmt.Printf("LeafType:%v", plant2.LeafType)
        fmt.Printf("Height:%v", plant2.Height)
    }
}
