package telegram

import (
	reversoApi "github.com/BRUHItsABunny/go-reverso-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

type TranslationResult struct {
	translations []string
	contextResults map[string]string
}

func NewTranslationResult() *TranslationResult {
	return &TranslationResult{
		translations: make([]string, 100),
		contextResults: make(map[string]string, 100),
	}
}

func (b *Bot) handleMessage(msg *tgbotapi.Message) error {
	langs := reversoApi.GetLanguages()
	trResp, err := b.client.Translate(msg.Text, langs[reversoApi.LanguageEnglish], langs[reversoApi.LanguageRussian])
	if err != nil {
		return nil
	}

	result := NewTranslationResult()
	result.translations = trResp.Translation

	replacer := strings.NewReplacer("<em>", "", "</em>", "")

	for _, res := range trResp.ContextResults.Results {
		for i := 0; i < len(res.SourceExamples); i++ {
			result.contextResults[replacer.Replace(res.SourceExamples[i])] = replacer.Replace(res.TargetExamples[i])
		}
	}

	var answer string
	answer = strings.Join(result.translations, ", ")
	answer += "\n"
	for key, val := range result.contextResults {
		answer += key + "---" + val + "\n"
	}

	newMsg := tgbotapi.NewMessage(msg.Chat.ID, answer)
	if _, err := b.bot.Send(newMsg); err != nil {
		return err
	}
	return nil
}

