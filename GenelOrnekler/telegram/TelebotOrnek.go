/*
------------------------------------
ixakblt - ibrahim AKBULUT
------------------------------------
Web Site :ixakblt
------------------------------------
All Sites : @ixakblt
------------------------------------
*/

// @RaifPy 'ye teşekkürler

package main

import (
	"log"
	"time"

	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	botToken := "1498255154:AAH0wKkdsdgsdhshd_fsdfsdfskdfsdfh5iqyiphQ_cjiCVTzfsdfsdfdsgsd0n9c9HZ24"
	bot, err := tb.NewBot(telebot.Settings{Token: botToken, Verbose: true, Poller: &telebot.LongPoller{Timeout: 10 * time.Second}}) //standart tanımlama
	if err != nil {
		log.Fatalln(err)
		return
	}

	bot.Handle(telebot.OnText, func(m *telebot.Message) { // Text emsaj atılırsa yapılacaklar
		a, _ := bot.Send(m.Sender, "selam") //gönderene selam diye yanıt ver
		time.Sleep(3 * time.Second)         //3 saniye bekle
		bot.Edit(a, "naber knk")            // selamı naber knk olarak değiştir

	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		bot.Reply(m, "hi") //gönderilen mesajı yanıtla

	})

	bot.Handle("/komut", func(m *telebot.Message) { // /komut komutu verilirse yapılacaklar

		bot.Send(m.Sender, "S")

	})

	bot.Handle(telebot.OnText, func(m *telebot.Message) {
		p := &tb.Photo{File: tb.FromDisk("205.jpg")} // tb içindeki photo struct 'ına FromDisk ile sistemden dosya yolu verek dosya atadık bu jpg
		v := &tb.Video{File: tb.FromDisk("q.mp4")}   // yukardaki işlemin video hali
		bot.Send(m.Sender, p)                        // ayrı ayrı göndermek için
		bot.SendAlbum(m.Sender, tb.Album{p, v})

	})

	bot.Handle(telebot.OnPhoto, func(m *telebot.Message) {
		//foto atıldığında yapılacak işlem aynı şekilde On diyerek vscode üzerinde ctrl spaceye basarsanız diğer methodları görebilirsiniz
		p := &tb.Photo{File: tb.FromDisk("205.jpg")}
		bot.Send(m.Sender, p)

	})

	bot.Start()
}
