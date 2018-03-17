require 'telegram/bot'

token = '377093324:AAHJTh7QedP3RbO9Q64AOnrlRZ1_tXjxvDw'
BUSCHE_USERNAME = 'jpbusche'
PARIS_STICKER = 'CAADAQADrQADZ6LPCqQOfJqDfhiwAg'

Telegram::Bot::Client.run(token) do |bot|
  bot.listen do |message|
    if message.from.username == BUSCHE_USERNAME
      if message.text.count('?') == 0
        bot.api.send_message(chat_id: message.chat.id, text: "#{message.text}?") 
      else
        bot.api.send_sticker(chat_id: message.chat.id, sticker: PARIS_STICKER)
      end
    end
  end
end
