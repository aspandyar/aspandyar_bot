package food

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

var userChoices = make(map[string]string)

var (
	BtnProteins   = tele.Btn{Text: "Белки", Unique: "proteins"}
	BtnCarbs      = tele.Btn{Text: "Углеводы", Unique: "carbs"}
	BtnFats       = tele.Btn{Text: "Здоровые жиры", Unique: "fats"}
	BtnAdditional = tele.Btn{Text: "Дополнительно", Unique: "additional"}
)

func HandleBreakfast(c tele.Context) error {
	MealSelector := &tele.ReplyMarkup{}
	MealSelector.Inline(
		MealSelector.Row(BtnProteins, BtnCarbs),
		MealSelector.Row(BtnFats, BtnAdditional),
	)

	return c.Send("Вы выбрали Завтрак. Выберите категорию:", MealSelector)
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
			nextStepButtons.Row(
				tele.Btn{Text: "Яйца (2-4 шт.)", Unique: "eggs"},
				tele.Btn{Text: "Творог или греческий йогурт", Unique: "cottage_cheese"},
				tele.Btn{Text: "Тофу", Unique: "tofu"},
				tele.Btn{Text: "Рыба (тунец, семга)", Unique: "fish"},
			),
		)
		userChoices["category"] = "Proteins"

	case "carbs":
		nextStepMessage = "Выберите углеводы (Carbs):"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(
				tele.Btn{Text: "Овсяная каша с фруктами", Unique: "oatmeal"},
				tele.Btn{Text: "Цельнозерновой хлеб с авокадо", Unique: "whole_grain_bread"},
			),
		)
		userChoices["category"] = "Carbs"

	case "fats":
		nextStepMessage = "Выберите здоровые жиры (Fats):"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(
				tele.Btn{Text: "Орехи", Unique: "nuts"},
				tele.Btn{Text: "Авокадо", Unique: "avocado"},
			),
		)
		userChoices["category"] = "Fats"

	case "additional":
		nextStepMessage = "Дополнительные опции:"
		nextStepButtons = &tele.ReplyMarkup{}
		nextStepButtons.Inline(
			nextStepButtons.Row(
				tele.Btn{Text: "Стакан молока", Unique: "milk"},
				tele.Btn{Text: "Протеиновый коктейль", Unique: "protein_shake"},
				tele.Btn{Text: "Пропустить", Unique: "none"},
			),
		)
		userChoices["category"] = "Additional"
	}

	c.Edit(nextStepMessage, nextStepButtons)

	return nil
}

func HandleFinalSelection(c tele.Context) error {
	selectedOption := c.Callback().Unique
	userChoices["item"] = selectedOption

	var summary string
	for key, value := range userChoices {
		summary += fmt.Sprintf("%s: %s\n", key, value)
	}

	c.Send("Вы завершили выбор. Ваши выборы:\n" + summary)

	return nil
}
