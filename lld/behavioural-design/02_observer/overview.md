### Observer Pattern

Suppose we have a weather station that records temperature and multiple devices (eg display units ) want to show the latest temperature. Without using the observer pattern, the weather station would have to explicitly inform each device about the temperature change, which results in tight coupling between the station and devices



## Observer pattern benefits

-> Loose Coupling: The subject (e.g WeatherStation) doesnt need to know about the specific observers. it just notifies them 

-> Scalability: New observers (e.g new display devices) can easily be added without changing the subject

-> Flexibility: Observers can be dynamically added or removed at runtime


## Observer Pattern Use cases

-> Event Listeners: GUI frameworks often use the observer pattern to implement event listeners for handling button clicks, input changes, etc

-> Stock Price Monitoring: When a stock price changes, multiple subscribers (like investors or systems) can be notified of the change

-> New Publishing Systems: News articles are published ( subject ), and subscribers (users) are notified whenever a new article is available

-> Social Media Notifications: Users can subscribe to updates from specific accounts and when an account posts (subject), all followers ( observers ) are notified

-> Logging Systems: Different logging handlers can observe events and log them as needed, such as to the consule, file or remote server