# Game Design 
 
Simple flash card game for learning German Aritcles. Der Die oder Das. 

Answer file would contain `der Apfel, the Apple` for example. Maybe as CSV. This allows as to determine if correct. `Der Die Das Apfel` is Presented like ![design](./img/initial_design.png)

Then it should keep track of when you answered correctly and when you answered incorrectly, asking the user more frequently about ones that were incorrect rarly but still required to ask the correct ones again. 

## Steps overview

1. Get the information
2. Store inforamtion
3. Display information
4. Basic gameplay
5. Points? 
6. Web or Desktop App?

# Step one: Getting word info.  

Lets use go to call an API. https://en.wiktionary.org/w/api.php has the information we need. We can scrape the word and store word + ARTICLE. 

DE API is different
https://de.wiktionary.org/api/rest_v1/#/Page%20content



CROSS pLatform GUI APP... could also write as web app... shall seee
https://dev.to/aurelievache/learning-go-by-examples-part-7-create-a-cross-platform-gui-desktop-app-in-go-44j1


When playing the "game" how to store results and store for later so that incorrect guess are repeated more often. 
----
## CONT. 
So after fluffing about about and failing to find a decent API which is unfortunate, we are forced to conclude that there is no decent API. Wikitionary's API only functions well for English, so useless for us. DWDS API which is another I found, didn't provide the ariticle in the API. DUDEN seems to have an API, but I can make such small amounts of requests 20. per month on the free account of 500 per month fo 39€ or something I gave up. It wasn't useful information. 

I then looked at Deepl.api which may end up being useful but it was still not giving me exactly what I wanted. So I have instead gone from webpage scraping. I am now scraping the wikitionary entries where I can the article and I can also get the english translatation. Below is the list of test words 
```golang
getGender("Tisch")
getGender("Tür")
getGender("Vogel")
getGender("Baum")
getGender("Apfel")
getGender("banane")

    //---OUTPUT----
m
table
f
door
m
bird
m
tree
m
apple
f
```
I have stuck with this approach because the amount of information available for extracting will probably grow and scraping wiktionary is pretty useful at this point.


The extractor is kind of working as well as expected at this point, I am retreiving the article and the translation. So we can now focus on getting the next steps. 

# Step 2 Data persistance

As I want to keep the data and score it I need to persist the data. I could store it on an ever growing CSV file, but given we are manipulatign the data or at least plan to in the future, I am going to choose a Database. 

I spent some time looking at how many other applications get away with this and Sqllite popped up again and agian. It appears to use standard SQL syntax meaning that we should be able to easily transition to other DB's if needed. However as it so goes we will start with SQLLITE as a docker container. `docker run --rm -it --name sql --network test -v "$(pwd):/workspace" -w /workspace keinos/sqlite3` Gets it up for testing purposes. 

I have dumped the DB so no we should be able to use .open ./ger_dict.db to perist the data if the container dies. Saves us getting these words.

--- 
OK learnt somethings about SQLite, so it essentailly a small file the lives as as example.db locally. There is is then a small engine instlled that allows to deal with that file. So the docker container is somewhat worthless at this point. INstll sqlite work forward from there. The Data is sucessfully saving is persisting. 

Now we need to consider the scoring aspect and how we will persist it. 

# Step 3 Adding a score

We need to keep track of the words score. So if you get it correct, Ihave written a small python script where i was experimenting with weighted random. So currently this is how I would propose to score the words in the db. The higher the score worse you did, ie greater chance of being called into the game, lesser score the better history, lesser chance of being called into game. I also considered adding a new column, this would only mark the word a new once to garuntee it called into the game, with 1 being new 0 being not new. Or what ever the boolean is in SQLite.  I could also just give it a high score, insuring it makes it into the game?? 

See DB below.

```sql
|aritcle|dword|eword | score | new
|m|Vogel|bird | 6 | 0 
|f|Tür|door | 5 | 1
|m|Tisch|table | 5 | 0
|f|Wurst|sausage | 5 | 0 
|n|Bier|beer | 5 | 1
|m|Hund|dog | 5 | 0

```

# Step 4 Seperation

We have now implemented some basic elements. 

I would now like to seperate out aspects of the program. So far we have focused solely on adding new words. This is great but now we need to seperate that into more functions. I am going to start with 4

1. Add new word function
2. Pull Set of words for game
3. Edit words
4. Delete words?

With these functions taken care of I a think we can then build out the rest of the game? 
