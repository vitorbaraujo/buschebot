import random
import string
from telegram.ext import Updater, MessageHandler, Filters

import vars

PUNCTUATION_MARKS = '.?!:;-()[]{}/'

updater = Updater(token=vars.API_TOKEN)
dispatcher = updater.dispatcher

phrases = [
    "pego meu carro e vou embora",
    "tchau",
    "oloco",
    "veneza parece uma favela",
    "cansei de ir pra europa",
    "o carro e meu",
]

reply = [
    "fala carai",
    "fala",
]

responses = [
    "sei la",
    "parei de trabalhar como vidente",
]

questions = [
    "como", "qual", "o que",
    "porque", "porquê", "por que", "pq",
]

def echo(bot, update):
    chat_id = update.message.chat_id
    sender_name = update.message.from_user.first_name + ' ' + update.message.from_user.last_name
    sender = update.message.from_user.username
    message = update.message.text
    answer = ""

    if "@busche_bot" in message:
        answer = random.choice(reply)
    else:
        if sender == 'jpbusche' and len(message.split()) > 1 and '?' not in message and message[-1] not in PUNCTUATION_MARKS:
            answer = message + '?'
        elif 'bot' in message.lower().split() or 'bot?' in message.lower().split():
            answer = random.choice(reply)
        elif '?' in message and len(message) > 1:
            if any(x in message.lower() for x in questions):
                answer = random.choice(responses)
            elif 'quem' in message:
                answer = random.choice(responses)
            else:
                answer = random.choice(["sim","nao","nao po"])

    print('Message "{}" sent by [{},{}] -> {}'.format(message, sender, sender_name, answer))
    bot.send_message(chat_id=chat_id, text=answer)

echo_handler = MessageHandler(Filters.text, echo)

dispatcher.add_handler(echo_handler)

updater.start_polling()
