package ex01

import (
	"errors"
	"testing"
)

func TestDescripte(t *testing.T) {
	plant1 := AnotherUnknownPlant{
		FlowerColor: 10,
		LeafType: "lanceolate",
		Height: 15,
	}

	err := DescripePlant(plant1)
	if err != nil {
		t.Errorf("err != nil, %v", err.Error())
		return
	}
}

func TestDescripePlant_UnknownPlant(t *testing.T) {
	plant := UnknownPlant{
		FlowerType: "Rose",
		LeafType:   "Serrated",
		Color:      16711680,
	}

	err := DescripePlant(plant)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestDescripePlant_AnotherUnknownPlant(t *testing.T) {
	plant := AnotherUnknownPlant{
		FlowerColor: 255,
		LeafType:    "Oval",
		Height:      12,
	}

	err := DescripePlant(plant)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestDescripePlant_UnsupportedType(t *testing.T) {
	notAPlant := "Just a string"

	err := DescripePlant(notAPlant)
	if err == nil {
		t.Errorf("Expected an error for unsupported type, but got none")
	} else if !errors.Is(err, errors.ErrUnsupported) {
		t.Errorf("Expected unsupported type error, but got: %v", err)
	}
}

func TestDescripePlant_NilInput(t *testing.T) {
	var plant interface{} = nil

	err := DescripePlant(plant)
	if err == nil {
		t.Errorf("Expected an error for nil input, but got none")
	} else if err.Error() != "reflect failed to parse type" {
		t.Errorf("Expected 'reflect failed to parse type' error, but got: %v", err)
	}
}