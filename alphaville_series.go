package main

type alphavilleSeries struct {
	UUID                   string                 `json:"uuid"`
	AlternativeIdentifiers alternativeIdentifiers `json:"alternativeIdentifiers,omitempty"`
	PrefLabel              string                 `json:"prefLabel"`
	Type                   string                 `json:"type"`
}

type alternativeIdentifiers struct {
	TME   []string `json:"TME,omitempty"`
	Uuids []string `json:"uuids,omitempty"`
}

type alphavilleSeriesLink struct {
	APIURL string `json:"apiUrl"`
}

type idEntry struct {
	ID string `json:"id"`
}
