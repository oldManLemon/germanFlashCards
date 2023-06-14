import random

# Define the words and their initial weights
# words = {'tisch': 99, 'bier': 50, 'vogel': 32, 'baum': 7}
words = {'tisch': 80, 'bier': 36,'tanz': 36, 'vogel': 32, 'baum': 7}

# Function to choose a word with weighted randomness
def choose_word():
    total_weight = sum(words.values())
    print("total weight: " + str(total_weight))
    rand_weight = random.uniform(0, 100)
    print("rand_weight: " + str(rand_weight))
    weight_sum = 0
    normalized_words = {word: (weight / total_weight) * 100 for word, weight in words.items()}
    print(normalized_words)
    for word, weight in  normalized_words.items():
        print("weight: " + str(weight))
        weight_sum += weight
        print("weight_sum: " + str(weight_sum))
        

        if weight_sum > rand_weight:
            return word



# Choose a word with weighted randomness
for i in range(1000):
    chosen_word = choose_word()
    # print(chosen_word)
    # print("chosen Word: "+chosen_word)
    with open('readme.txt', 'a') as f:
        f.write(chosen_word+'\n')

for w in words.keys():
    with open('readme.txt', 'r') as f:
        data = f.read()
        rate = data.count(w)
        print(w+": "+ str(rate))
