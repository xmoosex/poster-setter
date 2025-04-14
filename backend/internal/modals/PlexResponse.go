package modals

import "encoding/xml"

// PlexResponse represents the root MediaContainer element
type PlexResponse struct {
	XMLName             xml.Name            `xml:"MediaContainer"`
	Size                int                 `xml:"size,attr"`
	LibrarySectionID    string              `xml:"librarySectionID,attr"`
	LibrarySectionTitle string              `xml:"librarySectionTitle,attr"`
	ViewGroup           string              `xml:"viewGroup,attr"`      // Shows whether it is a movie, show
	Videos              []PlexVideoItem     `xml:"Video,omitempty"`     // Movies and Episodes
	Directory           []PlexDirectoryItem `xml:"Directory,omitempty"` // Shows and Seasons
}

// Video represents each Video element inside MediaContainer
// It is used for Movies
type PlexVideoItem struct {
	RatingKey      string          `xml:"ratingKey,attr"`
	Key            string          `xml:"key,attr"`
	Type           string          `xml:"type,attr"`
	Title          string          `xml:"title,attr"`
	ContentRating  string          `xml:"contentRating,attr"`
	Summary        string          `xml:"summary,attr"`
	Rating         float64         `xml:"rating,attr"`
	AudienceRating float64         `xml:"audienceRating,attr"`
	UserRating     float64         `xml:"userRating,attr"`
	ViewCount      int             `xml:"viewCount,attr"`
	LastViewedAt   int64           `xml:"lastViewedAt,attr"`
	Year           int             `xml:"year,attr"`
	Tagline        string          `xml:"tagline,attr"`
	Thumb          string          `xml:"thumb,attr"`
	Art            string          `xml:"art,attr"`
	UpdatedAt      int64           `xml:"updatedAt,attr"`
	Media          []PlexMediaItem `xml:"Media"`
	Index          int             `xml:"index,attr,omitempty"`       // Episode Number
	ParentIndex    int             `xml:"parentIndex,attr,omitempty"` // Season Number
}

// Media represents the Media element inside a Video
type PlexMediaItem struct {
	ID   string         `xml:"id,attr"`
	Part []PlexPartItem `xml:"Part"`
}

// Part represents the Part element inside a Media
type PlexPartItem struct {
	ID       string `xml:"id,attr"`
	Duration int64  `xml:"duration,attr"`
	File     string `xml:"file,attr"`
	Size     int64  `xml:"size,attr"`
}

// Directory represents each Directory element inside MediaContainer
// It is used for TV Shows
type PlexDirectoryItem struct {
	RatingKey      string  `xml:"ratingKey,attr"`
	Key            string  `xml:"key,attr"`
	Type           string  `xml:"type,attr"`
	Title          string  `xml:"title,attr"`
	ContentRating  string  `xml:"contentRating,attr"`
	Summary        string  `xml:"summary,attr"`
	Index          int     `xml:"index,attr"`
	AudienceRating float64 `xml:"audienceRating,attr"`
	UserRating     float64 `xml:"userRating,attr"`
	ViewCount      int     `xml:"viewCount,attr"`
	LastViewedAt   int64   `xml:"lastViewedAt,attr"`
	Year           int     `xml:"year,attr"`
	Thumb          string  `xml:"thumb,attr"`
	LeafCount      int     `xml:"leafCount,attr"`  // Episodes
	ChildCount     int     `xml:"childCount,attr"` // Seasons
	UpdatedAt      int64   `xml:"updatedAt,attr"`
	Location       struct {
		Path string `xml:"path,attr"`
	} `xml:"Location"`
}
