package handlers

import (
	"github.com/aspandyar/aspandyar_bot/bot/handlers/food"
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
)

func InitMarkups() {
	Menu.Reply(
		Menu.Row(BtnBreakfast, BtnLunch),
		Menu.Row(BtnSnack, BtnDinner),
	)
}

func (server *ServerBot) RegisterEatHandlers() {
	server.Bot.Handle(&BtnBreakfast, food.HandleBreakfast)

	server.Bot.Handle(&food.BtnProteins, food.HandleSelection)
	server.Bot.Handle(&food.BtnCarbs, food.HandleSelection)
	server.Bot.Handle(&food.BtnFats, food.HandleSelection)
	server.Bot.Handle(&food.BtnAdditional, food.HandleSelection)

	server.Bot.Handle(&food.BtnEggs, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnCottageCheese, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnTofu, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnFish, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnOatmeal, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnWholeGrainBread, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnNuts, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnAvocado, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnMilk, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnProteinShake, food.HandleFinalSelection)
	server.Bot.Handle(&food.BtnNone, food.HandleFinalSelection)
}
