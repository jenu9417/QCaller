# QCaller

## Assumptions, Concerns and More.


Why GO?
-------
QCaller is simple API service providing CRUD operations. It doesn't involve heavy business logics which requires much of inheritance, generics or reflection. The basic requirement of QCaller is simple multithreaded request handling at scale and who else does it better than Go.


Why ElasticSearch?
------------------
ElasticSearch is better for search usecases since it stores using inverted indices. ES is made specially for these usecases.


Why Aerospike?
--------------
Aerospike is a simple key value store, which provides blazing fast record fetch and is distributed in nature.


Data Model
----------
```
ID            -   String    // Compound id made up of source-id and number
Name          -   String    // Contact name
SourceID      -   String    // Source from which the contact was uploaded
Country       -   String    // Country for which the contact is valid
CountryCode   -   String    // International Calling Code for the country
Number        -   String    // Number
LastUpdated   -   Long      // UTC timestamp of last modification
```

**`ID`** - Made it compound id of sourceid and number. 
- This makes it easy to perform operations like get, delete which requires id. 
- Making it random, requires it to store at the client end, which will be a bad idea.
- Since sourceid + number is unique, we can store only one name per contact. Logical this makes sense too. This is the same conflict that will occur in Phone Caller ID when there are multiple contacts with same number. Hence this is a necessary evil to deal with.

**`Country & CountryCode`** - This is for the long term. 
- Ideally it shouldn't be limited to a particular country. 
- Also, country based indexing in es, helps to perform efficient search.
- Why both? - There are countries with same country code. Eg: USA and Canada.

**`Number`** - Single Number per record. This is done for multiple reasons.
- Incase of update operation on name, since we dont store id at client end, it would be a night mare to keep track of changes.
- Efficient searching. Single number per records helps es to index and search it better.
- Keeping it simple.
- Hashing of numbers is not done, to keep it simple.

**`LastUpdated`** - This is helpful for 2 reasons.
- We can get incremental list of updates happened in a timerange, which will ease syncing to aerospike.
- To keep a track of last edit.


One number per record
---------------------
Ideally the contacts in a phone can store multiple numbers against the same name. But we cannot use the same model, because that will be more inefficient for contact retrieval since we need to search inside arrays. Hence going with the one number per record model. 

This assumes the client is capable of two things:

- The contacts are sent in a one number per record model, while syncing. Splitting a single contact to multiple records.

- The country and country code logic. Ideally number are stored along with country code. Hence the client to should map the number to proper country and country code before syncing contact. This is possible since we already have country code with number. Hence in mobile its not an issue. In web search(if we are providing), we can get the input from user. (True caller does it that way ;-) )

These two logics simple and can be easily shifted to QCaller side, if we want to keep the client thin and dumb.


Data Operation Overview
-----------------------
The contacts are imported from source and synced to ElasticSearch in the data model specified. The contacts are stored in ES on country based index and in a one number per record operation. The contacts in ES are synced to Aerospike on a preferential basis. Only the top 1 result will be synced to AS. Hence the mapping there is number => record. We should be doing a periodic refresh of records in AS on a daily basis, to maintain top results in AS.


Contact Edit
------------
When a customer edits, 2 usecases
	
  **a)** edit name. - No issues, the entry corresponding to the numbers in ES will be updated with new name.
	
  **b)** edit number - Here, since our id is a compound id of number and sourceid, we cannot form the id. Also we dont store id at client end. Hence the updated contact will be added as new record in ES. The old contact will remain in ES. We could do periodic re-sync of records from client to clean these entries.

This update strategy was choosen from 5 possibilities. This one seemed more efficient.
- Keep track of changes at phone level and give atomic updates to create and delete - if possible best. requires old num and new num.
- **[Choosen]** Give only new updates. do a monthly refresh to purge old contacts - ok approach but maintainance is there.
- Keep a snapshot for each user - very bad
- Remove and reindex the whole contact - requires the contact to be single name multi number format. even then its difficult.
- Check if exists? - name exists num change, num exists name change.  -  bad logic


Get & Delete APIs
------------------
These apis should ideally done based on id logic. Since we dont store id at client end, this is not possible. But if the client application is aware of id creation logic, we can switch to id based api's


## Concerns
-----------
Unit test cases are generally kept to low. This is mainly because of two reasons.
- QCaller doesn't have much complex business logic. Its a simple CRUD library. Hence simple unit tests wouldn't add much value. Proper mocking should be done for request and db. This requires bit more time and it will be impossible to do that as well within the timespan since the scope of things to be done is vast.
	Unit tests should definitely be done and I agree not having much unit tests is a sin, and will definitely be done in next few days.
- Load testing. The main testing to be done for QCaller is definitely load testing. For that a proper load test framework needs to be created with random fake data creation which by itself is a task. Again this tasks will be done within few days on a preferential basis
