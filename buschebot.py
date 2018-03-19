import random
import string
from telegram.ext import Updater, MessageHandler, Filters

import vars

updater = Updater(token=vars.API_TOKEN)
dispatcher = updater.dispatcher

def echo(bot, update):
    chat_id = update.message.chat_id
    sender = update.message.from_user.username
    message = update.message.text

    print('Message "{}" sent by {}'.format(message, sender))

    punctuation_marks = '.?!:;-()[]{}/'

    if sender == 'jpbusche':
        if message.find('?') and message[-1] not in punctuation_marks:
            bot.send_message(chat_id=chat_id, text=message+'?')
        else:
            bot.send_sticker(chat_id=chat_id, sticker=random.choice(vars.STICKERS))

echo_handler = MessageHandler(Filters.text, echo)
dispatcher.add_handler(echo_handler)

updater.start_polling()
