## To-Do List
---

#### [ FEATURES ]
- Text Hints (Bold, Colors, Italics...)

#### [ PERFORMANCE ]
- Use "make" to pre-allocate space for incomming data
    - Have a config file storage to save performance information
       such as how much data is stored in file in order to allocate
- Make sure to sort each slice read from JSON
- Load only most recent (from bottom of JSON)
    - Have a "Load More" option
- Binary Search through Data Block for Month
- Save by Appending NEW data rather than all

#### [ DEVELOPMENT ]
- Implement Continuous Integration (TravisCI)