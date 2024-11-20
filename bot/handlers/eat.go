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

	server.Bot.Handle(&tele.Btn{Unique: "eggs"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "cottage_cheese"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "tofu"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "fish"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "oatmeal"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "whole_grain_bread"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "nuts"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "avocado"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "milk"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "protein_shake"}, food.HandleFinalSelection)
	server.Bot.Handle(&tele.Btn{Unique: "none"}, food.HandleFinalSelection)

}
