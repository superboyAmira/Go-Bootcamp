package comparator

import (
	"fmt"
	"goday01/internal/dbreader"
	"goday01/internal/input"
	"log/slog"
)

const (
	cakeADDED string = "ADDED cake \"%s\"\n"
	cakeREMOVED string = "REMOVED cake \"%s\"\n"

	timeCHANGED string = "CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n"

	ingredientADDED string = "ADDED ingredient \"%s\" for cake \"%s\"\n"
	ingredientREMOVED string = "REMOVED ingredient \"%s\" for cake \"%s\"\n"

	countUnitCHANGED string = "CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n"
	unitCHANGED string = "CHANGED unit for ingredient \"%s\" for cake  \"%s\" - \"%s\" instead of \"%s\"\n"
	unitREMOVED string = "REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n"
)

func Compare(cfg input.CLIcfg, log * slog.Logger) {
	reader_old, err := dbreader.GetReader(cfg.FileType_old)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	reader_new, err := dbreader.GetReader(cfg.FileType_new)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	err, data_old := reader_old.Load(cfg.Path_old, log)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	err, data_new := reader_new.Load(cfg.Path_new, log)
	if  err != nil {
		log.Debug(err.Error())
		return
	}
	for itr, cake_new  := range data_new.Cake {
		var AddedCake bool = true
		for _, cake_old  := range data_old.Cake {
			// removed cake
			if (itr == 0) {
				var RemovedCake bool = true
				for _, cake_new_tmp  := range data_new.Cake {
					if cake_new_tmp.Name == cake_old.Name {
						RemovedCake = false
					}
				}
				if RemovedCake {
					fmt.Printf(cakeREMOVED, cake_old.Name)
				}
			}
			if cake_new.Name == cake_old.Name {
				AddedCake = false
				// time
				if cake_new.Stovetime != cake_old.Stovetime {
					fmt.Printf(timeCHANGED, cake_new.Name, cake_new.Stovetime, cake_old.Stovetime)
				}
				// ingredents
				for itr_ing, ingredient_new := range cake_new.Ingredients {
					var AddedIngr bool = true
					for _, ingredient_old := range cake_old.Ingredients {
						// removed ingredent
						if itr_ing == 0 {
							var RemovedIngr bool = true
							for _, ingredient_new_tmp := range cake_new.Ingredients {
								if ingredient_new_tmp.Name == ingredient_old.Name {
									RemovedIngr = false
								}
							}
							if RemovedIngr {
								fmt.Printf(ingredientREMOVED, ingredient_old.Name, cake_old.Name)
							}
						}
						if ingredient_new.Name == ingredient_old.Name {
							AddedIngr = false
							if ingredient_new.Count != ingredient_old.Count {
								fmt.Printf(countUnitCHANGED, ingredient_old.Name, cake_old.Name, ingredient_new.Count, ingredient_old.Count)
							}
							if ingredient_new.Unit != ingredient_old.Unit {
								fmt.Printf(unitCHANGED, ingredient_old.Name, cake_old.Name, ingredient_new.Unit, ingredient_old.Unit)
							}
						}
					}
					// added ingredent 
					if AddedIngr {
						fmt.Printf(ingredientADDED, ingredient_new.Name, cake_new.Name)
					}
				}
			}
		}
		// added cake
		if AddedCake {
			fmt.Printf(cakeADDED, cake_new.Name)
		}
	}
}

/*
 * Comparison between two files large dimension, based on method Hash-sum comparison
 * wiki link: https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions
 */
func HashByHashComparator(cfg input.CLIcfg, log *slog.Logger) bool {

}