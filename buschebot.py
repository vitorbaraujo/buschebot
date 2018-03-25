import re
import random
import string
from telegram.ext import Updater, MessageHandler, Filters

import vars
import sentences as s

PUNCTUATION_MARKS = '.?!:;-()[]{}/'
PQ_RGX = 'po*r* *q'
BOT_RGX = '(^| |[^a-z])(busche|bot)($| |[^a-z])'
BOT_NAME = '@busche_bot'
BUSCHE_ID = 152231281
RGX_OPTIONS = re.M | re.I

updater = Updater(token=vars.API_TOKEN)
dispatcher = updater.dispatcher

message_counter = {}

def echo(bot, update):
    chat_id = update.message.chat_id
    user = update.message.from_user
    first_name = user.first_name.lower()
    name = 'temp_name'
    if user.first_name:
        name = user.first_name
    if user.last_name:
        name += ' ' + user.last_name
    username = user.username
    message = update.message.text
    answer = ""

    is_question = re.search(PQ_RGX, message, RGX_OPTIONS) \
        or any(q in message.lower() for q in s.questions)
    length = len(message.split())

    if name in message_counter and message_counter[name] >= 5:
        answer = random.choice(s.stop)
        if answer[-1] == ' ' and first_name:
            answer += first_name
        message_counter[name] = 0
    else:
        if user.id == BUSCHE_ID and length > 1 and '?' not in message:
            answer = message + '?'
        elif '?' in message and len(set(list(message.replace(' ','')))) != 1:
            if is_question:
                answer = random.choice(s.responses)
            elif 'o que' in message.lower():
                answer = 'sei la po'
            else:
                answer = random.choice(s.yes_no)
        elif BOT_NAME in message:
            answer = random.choice(s.reply)
        elif re.search(BOT_RGX, message, RGX_OPTIONS):
            answer = random.choice(s.reply)

    if len(answer) > 0:
        if name in message_counter:
            message_counter[name] += 1
        else:
            message_counter[name] = 0

    print('{}({}) -> {} times'.format(name, username, message_counter[name]))
    print('  message: "{}"'.format(message))
    print('  reply: "{}"'.format(answer))
    print('-' * 50)

    bot.send_message(chat_id=chat_id, text=answer)

echo_handler = MessageHandler(Filters.text, echo)
dispatcher.add_handler(echo_handler)

updater.start_polling()
