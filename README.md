# golang-chirpReparser

CHIRP nicely grabs data from remote databases like RFinder

However, it does not filter as I desire.  So I threw this together to make adjustments like:
	- Only stations closer than a given distance (8 miles default)
	- Only stations with a frequency within a range 

Without this I had stations that were 20 miles away at the top of the list and some closer too far down the list to fit in Xceiver memory.
Also, it included stations that I just could not hear as they were outside my radios range.
