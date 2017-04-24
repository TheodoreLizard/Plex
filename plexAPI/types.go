package plexAPI

import "encoding/xml"

type MediaContainer struct {
	XMLName             xml.Name    `xml:"MediaContainer"`
	Size                int         `xml:"size,attr"`
	AllowSync           int         `xml:"allowSync,attr"`
	Identifier          string      `xml:"identifier,attr"`
	LibrarySectionID    int         `xml:"librarySectionID,attr"`
	LibrarySectionTitle string      `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string      `xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string      `xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string      `xml:"mediaTagVersion,attr"`
	Directory           []Directory `xml:"Directory"`
	Playlists           []Playlist  `xml:"Playlist"`
	Videos              []Video     `xml:"Video"`
}

type Directory struct {
	XMLName xml.Name `xml:"Directory"`
	Key     string   `xml:"key,attr"`
	Title   string   `xml:"title,attr"`
	Type    string   `xml:"type,attr"`
}

type Playlist struct {
	XMLName           xml.Name `xml:"Playlist"`
	RatingKey         string   `xml:"ratingKey,attr"`
	Key               string   `xml:"key,attr"`
	GUID              string   `xml:"guid,attr"`
	Type              string   `xml:"type,attr"`
	Title             string   `xml:"title,attr"`
	Summary           string   `xml:"summary,attr"`
	Smart             string   `xml:"smart,attr"`
	PlaylistType      string   `xml:"playlistType,attr"`
	Composite         string   `xml:"composite,attr"`
	ViewCount         int      `xml:"viewCount,attr"`
	LastViewedAt      string   `xml:"lastViewedAt,attr"`
	Duration          int      `xml:"duration,attr"`
	LeafCount         int      `xml:"leafCount,attr"`
	AddedAt           string   `xml:"addedAt,attr"`
	UpdatedAt         string   `xml:"updatedAt,attr"`
	DurationInSeconds string   `xml:"durationInSeconds,attr"`
}

type Video struct {
	XMLName               xml.Name `xml:"Video"`
	Key                   string   `xml:"key,attr"`
	GUID                  string   `xml:"guid,attr"`
	Type                  string   `xml:"type,attr"`
	Title                 string   `xml:"title,attr"`
	Summary               string   `xml:"summary,attr"`
	Year                  string   `xml:"year,attr"`
	Thumb                 string   `xml:"thumb,attr"`
	Art                   string   `xml:"art,attr"`
	PlaylistItemID        string   `xml:"playlistItemID,attr"`
	Duration              string   `xml:"duration,attr"`
	OriginallyAvailableAt string   `xml:"originallyAvailableAt,attr"`
	LibrarySectionID      int      `xml:"librarySectionID,attr"`
	LibrarySectionTitle   string   `xml:"librarySectionTitle,attr"`
	LibrarySectionUUID    string   `xml:"librarySectionUUID,attr"`
	AddedAt               string   `xml:"addedAt,attr"`
	UpdatedAt             string   `xml:"updatedAt,attr"`
	ChapterSource         string   `xml:"chapterSource,attr"`
	Media                 Media    `xml:"Media,omitemtpy"`
	Genre                 Genre    `xml:"Genre,omitempty"`
	Role                  Role     `xml:"Role,omitempty"`
}

type Media struct {
	XMLName               xml.Name `xml:"Media"`
	VideoResolution       string   `xml:"videoResolution,attr"`
	ID                    string   `xml:"id,attr"`
	Duration              string   `xml:"duration,attr"`
	Bitrate               string   `xml:"bitrate,attr"`
	Width                 string   `xml:"width,attr"`
	Height                string   `xml:"height,attr"`
	AspectRatio           string   `xml:"aspectRatio,attr"`
	AudioChannels         string   `xml:"audioChannels,attr"`
	AudioCodec            string   `xml:"audioCodec,attr"`
	VideoCodec            string   `xml:"videoCodec,attr"`
	Container             string   `xml:"container,attr"`
	VideoFrameRate        string   `xml:"videoFrameRate,attr"`
	OptimizedForStreaming string   `xml:"optimizedForStreaming,attr"`
	AudioProfile          string   `xml:"audioProfile,attr"`
	Has64bitOffsets       string   `xml:"has64bitOffsets,attr"`
	VideoProfile          string   `xml:"videoProfile,attr"`
	Part                  []Part   `xml:"Part,omitempty"`
}

type Part struct {
	XMLName               xml.Name `xml:"Part"`
	ID                    string   `xml:"id,attr"`
	Key                   string   `xml:"key,attr"`
	Duration              string   `xml:"duration,attr"`
	File                  string   `xml:"file,attr"`
	Size                  string   `xml:"size,attr"`
	AudioProfile          string   `xml:"audioProfile,attr"`
	Container             string   `xml:"container,attr"`
	Has64bitOffsets       string   `xml:"has64bitOffsets,attr"`
	HasChapterTextStream  string   `xml:"hasChapterTextStream,attr"`
	HasThumbnail          string   `xml:"hasThumbnail,attr"`
	OptimizedForStreaming string   `xml:"optimizedForStreaming,attr"`
	VideoProfile          string   `xml:"videoProfile,attr"`
}

type Genre struct {
	XMLName xml.Name `xml:"Genre"`
	Tag     string   `xml:"tag,attr,omitempty"`
}

type Role struct {
	XMLName xml.Name `xml:"Role"`
	Tag     string   `xml:"tag,attr,omitempty"`
}

//-----------------------------------------------------------------------
type Playlists struct {
	XMLName   xml.Name        `xml:"Playlists"`
	Playlists []SavedPlaylist `xml:"Playlist"`
}

type SavedPlaylist struct {
	XMLName xml.Name     `xml:"Playlist"`
	Title   string       `xml:"title,attr"`
	Key     string       `xml:"key,attr"`
	Videos  []SavedVideo `xml:"Video"`
}

type SavedVideo struct {
	XMLName            xml.Name `xml:"Video"`
	Title              string   `xml:"title,attr"`
	Year               string   `xml:"year,attr"`
	Key                string   `xml:"key,attr"`
	LibrarySectionUUID string   `xml:"librarySectionUUID,attr"`
}
