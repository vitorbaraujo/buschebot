require 'telegram_bot'

bot = TelegramBot.new(token: '377093324:AAHJTh7QedP3RbO9Q64AOnrlRZ1_tXjxvDw')

bot.get_updates(fail_silently: true) do |message|
  puts "@#{message.from.username}: #{message.text}"

  command = message.get_command_for(bot)

  message.reply do |reply|
    if message.from.username == "jpbusche"
      unless command[-1] == '?'
        reply.text = "#{message.text}?"
        reply.send_with(bot)
      end
    end
  end
end
