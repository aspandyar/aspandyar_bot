package handlers

import (
	tele "gopkg.in/telebot.v3"
)

type ServerBot struct {
	Bot *tele.Bot
}

var (
	Menu         = &tele.ReplyMarkup{ResizeKeyboard: true}
	BtnBreakfast = Menu.Text("🍳 Breakfast/Завтрак")
	BtnLunch     = Menu.Text("🥪 Lunch/Обед")
	BtnSnack     = Menu.Text("🍪 Snack/Перекус")
	BtnDinner    = Menu.Text("🍲 Dinner/Ужин")

	MealSelector  = &tele.ReplyMarkup{}
	BtnProteins   = MealSelector.Data("Белки", "proteins")
	BtnCarbs      = MealSelector.Data("Углеводы", "carbs")
	BtnFats       = MealSelector.Data("Здоровые жиры", "fats")
	BtnAdditional = MealSelector.Data("Дополнительно", "additional")

	BtnEggs            = MealSelector.Data("Яйца (2-4 шт.)", "eggs", "eggs")
	BtnCottageCheese   = MealSelector.Data("Творог или греческий йогурт", "cottage_cheese")
	BtnTofu            = MealSelector.Data("Тофу", "tofu")
	BtnFish            = MealSelector.Data("Рыба (тунец, семга)", "fish")
	BtnOatmeal         = MealSelector.Data("Овсяная каша с фруктами", "oatmeal")
	BtnWholeGrainBread = MealSelector.Data("Цельнозерновой хлеб с авокадо", "whole_grain_bread")
	BtnNuts            = MealSelector.Data("Здоровые жиры: Орехи", "nuts")
	BtnAvocado         = MealSelector.Data("Авокадо", "avocado")
	BtnMilk            = MealSelector.Data("Стакан молока", "milk")
	BtnProteinShake    = MealSelector.Data("Протеиновый коктейль", "protein_shake")
)

func InitMarkups() {
	Menu.Reply(
		Menu.Row(BtnBreakfast, BtnLunch),
		Menu.Row(BtnSnack, BtnDinner),
	)
}

func HandleEat(c tele.Context) error {
	return c.Send("Choose your meal for breakfast:", MealSelector)
}

func HandleBreakfast(c tele.Context) error {
	MealSelector.Inline(
		MealSelector.Row(BtnProteins, BtnCarbs),
		MealSelector.Row(BtnFats, BtnAdditional),
	)

	return c.Send("Вы выбрали Завтрак. Now, choose your protein source:", MealSelector)
}

func HandleProteins(c tele.Context) error {
	MealSelector.Inline(
		MealSelector.Row(BtnEggs, BtnCottageCheese),
		MealSelector.Row(BtnTofu, BtnFish),
	)

	return c.Send("Выберите белки (Proteins):", MealSelector)
}

func HandleCarbs(c tele.Context) error {
	carbsSelector := &tele.ReplyMarkup{}
	carbsSelector.Inline(
		carbsSelector.Row(BtnOatmeal, BtnWholeGrainBread),
		carbsSelector.Row(BtnNuts, BtnAvocado),
	)

	return c.Send("Выберите углеводы (Carbs):", carbsSelector)
}

func HandleAdditional(c tele.Context) error {
	additionalSelector := &tele.ReplyMarkup{}
	additionalSelector.Inline(
		additionalSelector.Row(BtnMilk, BtnProteinShake),
	)

	return c.Send("Дополнительные опции:", additionalSelector)
}

func HandleSelection(c tele.Context) error {
	selectedOption := c.Callback().Unique

	var response string
	switch selectedOption {
	case "eggs":
		response = "You chose Eggs (2-4 pieces) as your protein."
	case "cottage_cheese":
		response = "You chose Cottage Cheese or Greek Yogurt as your protein."
	case "tofu":
		response = "You chose Tofu as your protein."
	case "fish":
		response = "You chose Fish (Tuna, Salmon, etc.) as your protein."
	case "oatmeal":
		response = "You chose Oatmeal with fruits as your carbs."
	case "whole_grain_bread":
		response = "You chose Whole Grain Bread with Avocado as your carbs."
	case "nuts":
		response = "You chose Nuts as your healthy fats."
	case "avocado":
		response = "You chose Avocado as your healthy fat."
	case "milk":
		response = "You chose Milk for additional protein and calcium."
	case "protein_shake":
		response = "You chose a Protein Shake."
	default:
		response = "Invalid option."
	}

	return c.Send(response)
}

func (server *ServerBot) RegisterEatHandlers() {
	server.Bot.Handle("/eat", HandleEat)
	server.Bot.Handle(&BtnBreakfast, HandleBreakfast)
	server.Bot.Handle(&BtnProteins, HandleProteins)
	server.Bot.Handle(&BtnCarbs, HandleCarbs)
	server.Bot.Handle(&BtnFats, HandleSelection)
	server.Bot.Handle(&BtnAdditional, HandleAdditional)

	server.Bot.Handle(&BtnEggs, HandleSelection)
	server.Bot.Handle(&BtnCottageCheese, HandleSelection)
	server.Bot.Handle(&BtnTofu, HandleSelection)
	server.Bot.Handle(&BtnFish, HandleSelection)
	server.Bot.Handle(&BtnOatmeal, HandleSelection)
	server.Bot.Handle(&BtnWholeGrainBread, HandleSelection)
	server.Bot.Handle(&BtnNuts, HandleSelection)
	server.Bot.Handle(&BtnAvocado, HandleSelection)
	server.Bot.Handle(&BtnMilk, HandleSelection)
	server.Bot.Handle(&BtnProteinShake, HandleSelection)
}
