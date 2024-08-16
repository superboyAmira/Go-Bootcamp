package ex01

import (
	"errors"
	"fmt"
	"reflect"
)

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

func DescripePlant(plant any) (error) {
    typ := reflect.TypeOf(plant)
    if typ == nil  {
        return errors.New("reflect failed to parse type")
    }
    if typ.Kind() != reflect.Struct {
        return errors.ErrUnsupported
    }
    value := reflect.ValueOf(plant)

    switch typ.Name() {
    case "UnknownPlant":
        for i := 0; i < value.NumField(); i++ {
            field := typ.Field(i)
            value := value.Field(i).Interface()
            tag := field.Tag.Get("color_scheme")

            if tag != "" {
                fmt.Printf("%v(%v=%v):%v\n", field.Name, "color_scheme", tag, value)
            } else {
                fmt.Printf("%v:%v\n", field.Name, value)
            }
        }
    case "AnotherUnknownPlant":
        for i := 0; i < value.NumField(); i++ {
            field := typ.Field(i)
            value := value.Field(i).Interface()
            tag := field.Tag.Get("unit")

            if tag != "" {
                fmt.Printf("%v(%v=%v):%v\n", field.Name, "unit", tag, value)
            } else {
                fmt.Printf("%v:%v\n", field.Name, value)
            }
        }
    default:
        return errors.ErrUnsupported
    }
    return nil
}
