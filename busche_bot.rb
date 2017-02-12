require 'telegram/bot'

token = 'bot_api_token'
BUSCHE_USERNAME = 'jpbusche'
PARIS_STICKER = 'CAADAQADrQADZ6LPCqQOfJqDfhiwAg'

Telegram::Bot::Client.run(token) do |bot|
  bot.listen do |message|
    #if message.from.username == BUSCHE_USERNAME
      if message.text[-1] != '?'
        bot.api.send_message(chat_id: message.chat.id, text: "#{message.text}?")
      else
        bot.api.send_sticker(chat_id: message.chat.id, sticker: PARIS_STICKER)
      end
    #end
  end
end
