import random

# Define the words and their initial weights
words = {'tisch': 5, 'bier': 5, 'vogel': 5, 'baum': 5}

# Function to choose a word with weighted randomness
def choose_word():
    total_weight = sum(words.values())
    # print("total weight: "+ str(total_weight))
    rand_weight = random.uniform(0, total_weight)
    # print("rand_weight: "+str(rand_weight))
    weight_sum = 0
    for word, weight in words.items():
        # print(word)
        # print(weight)
        weight_sum += weight
        # print("weight_sum: "+ str(weight_sum))

        if weight_sum >= rand_weight:
            return word



# Choose a word with weighted randomness
for i in range(1000):
    chosen_word = choose_word()
    # print("chosen Word: "+chosen_word)
    with open('readme.txt', 'a') as f:
        f.write(chosen_word+'\n')


for w in words.keys():
    with open('readme.txt', 'r') as f:
        data = f.read()
        rate = data.count(w)
        print(w+": "+ str(rate))
