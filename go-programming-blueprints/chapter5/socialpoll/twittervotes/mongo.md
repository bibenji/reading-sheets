use ballots
db.polls.insert({"title":"Test poll","options":["fraise","banane"]})

use ballots
db.polls.find().pretty()

