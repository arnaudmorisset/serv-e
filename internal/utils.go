package internal

/// Reverse a slice of records.
/// This solution come from the official Go wiki.
/// https://github.com/golang/go/wiki/SliceTricks#reversing
func ReverseRecords(records *[]Record) {
	nbRecords := len(*records)
	for i := nbRecords/2 - 1; i >= 0; i-- {
		opp := nbRecords - 1 - i
		(*records)[i], (*records)[opp] = (*records)[opp], (*records)[i]
	}
}
