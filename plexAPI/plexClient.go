package plexAPI

import (
	"Plex/httpClient"
	"encoding/xml"
	"fmt"
	"net/url"
)

type PlexClient struct {
	httpClient *httpClient.Client
}

func NewPlexClient(host string) *PlexClient {
	httpClient := httpClient.NewClient(host)

	plexClient := &PlexClient{
		httpClient: httpClient,
	}

	return plexClient
}

func (pc *PlexClient) Get(request string) (MediaContainer, error) {
	var mediaContainer MediaContainer

	req, err := pc.httpClient.NewGet(request, nil)
	if err != nil {
		return mediaContainer, err
	}

	// bytes, err := httputil.DumpRequest(req, true)
	// fmt.Printf("%s\n", string(bytes))

	resp, err := pc.httpClient.Invoke(req)
	if err != nil {
		return mediaContainer, err
	}

	defer resp.Body.Close()

	// bytes, err := httputil.DumpResponse(resp, true)
	// fmt.Printf("%s\n", string(bytes))

	err = xml.NewDecoder(resp.Body).Decode(&mediaContainer)
	return mediaContainer, err
}

func (pc *PlexClient) Put(request string) (MediaContainer, error) {
	var mediaContainer MediaContainer
	req, err := pc.httpClient.NewPut(request, nil)
	if err != nil {
		return mediaContainer, err
	}

	resp, err := pc.httpClient.Invoke(req)
	if err != nil {
		return mediaContainer, err
	}

	defer resp.Body.Close()

	err = xml.NewDecoder(resp.Body).Decode(&mediaContainer)
	return mediaContainer, err
}

/*
 * Search
 */
func (pc *PlexClient) SearchPlaylist(title string) (Playlist, error) {
	var playlist Playlist

	mc, err := pc.Get("search?type=15&query=" + url.QueryEscape(title))
	if err != nil {
		return playlist, err
	}

	for _, pl := range mc.Playlists {
		if title == pl.Title {
			return pl, nil
		}
	}

	return playlist, fmt.Errorf("Playlist [%s] not found", title)
}

func (pc *PlexClient) SearchLocal(title string) ([]Video, error) {
	query := fmt.Sprintf("search?local=1&query=%s", url.QueryEscape(title))

	mc, err := pc.Get(query)
	if err != nil {
		return nil, err
	}

	return mc.Videos, nil
}
