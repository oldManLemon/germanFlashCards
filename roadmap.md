# Game Design 
 
Simple flash card game for learning German Aritcles. Der Die oder Das. 

Answer file would contain `der Apfel, the Apple` for example. Maybe as CSV. This allows as to determine if correct. `Der Die Das Apfel` is Presented like ![design](./img/initial_design.png)

Then it should keep track of when you answered correctly and when you answered incorrectly, asking the user more frequently about ones that were incorrect rarly but still required to ask the correct ones again. 

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