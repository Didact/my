package my

type SearchData struct {
	ResponseData struct {
		Results []struct {
			GsearchResultClass string
			UnescapedUrl       string
			Url                string
			CacheUrl           string
			Title              string
			TitleNoFormatting  string
			Conent             string
		}
		Cursor struct {
			ResultCount string
			Pages       []struct {
				Start string
				Label float64
			}
			EstimatedResultCount string
			CurrentPageIndex     float64
			MoreResultsUrl       string
			SearchResultTime     string
		}
	}
	ResponseDetails interface{}
	ResponseStatus  float64
}
