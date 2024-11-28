package food

import (
	tele "gopkg.in/telebot.v3"
)

type Food struct {
	CategoryChoice string
	FoodSelection  string
}

var foodInstance []Food

var (
	BtnProteins   = tele.Btn{Text: "Белки", Unique: "proteins"}
	BtnCarbs      = tele.Btn{Text: "Углеводы", Unique: "carbs"}
	BtnFats       = tele.Btn{Text: "Здоровые жиры", Unique: "fats"}
	BtnAdditional = tele.Btn{Text: "Дополнительно", Unique: "additional"}

	BtnEggs            = tele.Btn{Text: "Яйца (2-4 шт.)", Unique: "eggs"}
	BtnCottageCheese   = tele.Btn{Text: "Творог или греческий йогурт", Unique: "cottage_cheese"}
	BtnTofu            = tele.Btn{Text: "Тофу", Unique: "tofu"}
	BtnFish            = tele.Btn{Text: "Рыба (тунец, семга)", Unique: "fish"}
	BtnOatmeal         = tele.Btn{Text: "Овсяная каша с фруктами", Unique: "oatmeal"}
	BtnWholeGrainBread = tele.Btn{Text: "Цельнозерновой хлеб с авокадо", Unique: "whole_grain_bread"}
	BtnNuts            = tele.Btn{Text: "Орехи", Unique: "nuts"}
	BtnAvocado         = tele.Btn{Text: "Авокадо", Unique: "avocado"}
	BtnMilk            = tele.Btn{Text: "Стакан молока", Unique: "milk"}
	BtnProteinShake    = tele.Btn{Text: "Протеиновый коктейль", Unique: "protein_shake"}
	BtnNone            = tele.Btn{Text: "Пропустить", Unique: "none"}
)

func HandleBreakfast(c tele.Context) error {
	MealSelector := &tele.ReplyMarkup{}
	MealSelector.Inline(
		MealSelector.Row(BtnProteins, BtnCarbs),
		MealSelector.Row(BtnFats, BtnAdditional),
	)

	return c.Send("Выберите категорию:", MealSelector)
}

func HandleSelection(c tele.Context) error {
	selectedOption := c.Callback().Unique
	var nextStepMessage string
	var nextStepButtons *tele.ReplyMarkup

	switch selectedOption {
	case "proteins":
		nextStepMessage = "Выберите белки (Proteins):"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(BtnEggs, BtnCottageCheese),
			nextStepButtons.Row(BtnTofu, BtnFish),
		)
	case "carbs":
		nextStepMessage = "Выберите углеводы (Carbs):"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(BtnOatmeal, BtnWholeGrainBread),
		)
	case "fats":
		nextStepMessage = "Выберите здоровые жиры (Fats):"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(BtnNuts, BtnAvocado),
		)
	case "additional":
		nextStepMessage = "Дополнительные опции:"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(BtnMilk, BtnProteinShake, BtnNone),
		)
	}

	return c.Edit(nextStepMessage, nextStepButtons)
}

func HandleFinalSelection(c tele.Context) error {
	err := c.Respond()
	if err != nil {
		return err
	}

	err = c.Delete()
	if err != nil {
		return err
	}

	choosenFood := c.Callback().Unique
	if choosenFood != "" {

	}

	return HandleBreakfast(c)
}
