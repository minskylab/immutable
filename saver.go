package immutable

func NewRecord(config *Config) (*ImmutableRecord, error) {
	if err := GeneratePDF(config); err != nil {
		panic(err)
	}

	doc, err := PinDocumentToIPFS(config)
	if err != nil {
		panic(err)
	}

	record := &ImmutableRecord{
		CID:  doc.IpfsHash,
		Date: doc.Timestamp,
	}

	return record, nil
}
