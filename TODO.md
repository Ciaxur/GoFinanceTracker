## To-Do List
---

#### [ FEATURES ]
- Year Summary
    - Similair Outcome to Monthly Summary
- Generate Graph for Year
- When Always Check Current Year for Month Number

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
