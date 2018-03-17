import telegram
from telegram import InlineQueryResultArticle, InputTextMessageContent
import string
from telegram.ext import Updater, MessageHandler, InlineQueryHandler, Filters

import vars

bot = telegram.Bot(token=vars.API_TOKEN)
updater = Updater(token=vars.API_TOKEN)
dispatcher = updater.dispatcher

print('This is me: {}'.format(bot.get_me()))

def echo(bot, update):
    chat_id = update.message.chat_id
    sender = update.message.chat.username
    message = update.message.text

    print('Sent by', sender)

    punctuation_marks = '.?!:;-()[]{}/'

    if sender == 'jpbusche':
        if message.find('?') and message[-1] not in punctuation_marks:
            bot.send_message(chat_id=chat_id, text=message+'?')
        else:
            bot.send_sticker(chat_id=chat_id, sticker=vars.PARIS_STICKER)

echo_handler = MessageHandler(Filters.text, echo)
dispatcher.add_handler(echo_handler)

updater.start_polling()
