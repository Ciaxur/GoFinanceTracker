## To-Do List
---

#### [ FEATURES ]
- Year Summary
    - Similair Outcome to Monthly Summary
- Generate Graph for Year
- When Always Check Current Year for Month Number
- Save Each Year in a seperate JSON File
    - Implements "Load Year" Option
        - Display all the Years Available to Load (But default is to load most recent Year)
    - When Searching for Month Number, it should be the most recent or the one specifically loaded in

#### [ PERFORMANCE ]
- Use "make" to pre-allocate space for incomming data (12 Blocks / Year)
    - Have a config file storage to save performance information such as how much data is stored in file in order to allocate
- Make sure to sort each slice read from JSON

#### [ DEVELOPMENT ]
- Implement Continuous Integration (TravisCI)
