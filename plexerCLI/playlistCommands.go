package main

import (
	"Plex/plexAPI"
	"encoding/xml"
	"fmt"
	"net/url"

	"os"

	"github.com/urfave/cli"
)

func playlistCommand() cli.Command {
	return cli.Command{
		Name:  "playlist",
		Usage: "playlist [save|restore] [filename]",
		Subcommands: []cli.Command{
			savePlaylistCommand(),
			restorePlaylistCommand(),
		},
	}
}

func savePlaylistCommand() cli.Command {
	return cli.Command{
		Name:  "save",
		Usage: "save [filename]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return fmt.Errorf("Please specify filename")
			}
			filename := c.Args().First()

			playlists, err := getPlaylists()
			if err != nil {
				return err
			}

			return saveXML(playlists, filename)
		},
	}
}

func restorePlaylistCommand() cli.Command {
	return cli.Command{
		Name:  "restore",
		Usage: "restore [filename]",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				return fmt.Errorf("Please specify filename")
			}
			filename := c.Args().First()

			var playlists plexAPI.Playlists
			err := loadXML(filename, &playlists)
			if err != nil {
				return err
			}

			return restorePlaylists(playlists)
		},
	}
}

func getPlaylists() (plexAPI.Playlists, error) {
	var playlists plexAPI.Playlists

	plexClient := plexAPI.NewPlexClient(plexAddr, plexToken)
	mc, err := plexClient.Get("playlists")
	if err != nil {
		return playlists, err
	}

	for _, pl := range mc.Playlists {
		playList := plexAPI.SavedPlaylist{
			Title: pl.Title,
		}

		mc, err = plexClient.Get(pl.Key)
		if err != nil {
			return playlists, err
		}

		for _, v := range mc.Videos {
			video := plexAPI.SavedVideo{
				Title: v.Title,
				Year:  v.Year,
			}

			playList.Videos = append(playList.Videos, video)
		}

		playlists.Playlists = append(playlists.Playlists, playList)
	}

	return playlists, nil
}

func restorePlaylists(playlists plexAPI.Playlists) error {
	plexClient := plexAPI.NewPlexClient(plexAddr, plexToken)

	for _, pl := range playlists.Playlists {
		playlist, err := plexClient.SearchPlaylist(pl.Title)
		if err != nil {
			playlist, err = plexClient.CreatePlaylist(pl.Title)
			if err != nil {
				return err
			}
		}

		err = restorePlaylist(plexClient, pl, playlist.Key)
		if err != nil {
			return err
		}

		err = organizePlaylist(plexClient, pl, playlist.Key)
		if err != nil {
			return err
		}
	}

	return nil
}

func restorePlaylist(plexClient *plexAPI.PlexClient, playlist plexAPI.SavedPlaylist, playlistKey string) error {
	fmt.Printf("Restore playlist %s\n", playlist.Title)

	for _, v := range playlist.Videos {
		fmt.Printf("	Restore %s\n", v.Title)

		video, err := plexClient.FindVideo(v)
		if err != nil {
			return err
		}

		videoURI := fmt.Sprintf("library://%s/item/%s", video.LibrarySectionUUID, video.Key)
		request := fmt.Sprintf("%s?uri=%s", playlistKey, url.QueryEscape(videoURI))

		mc, err := plexClient.Put(request)
		if err != nil {
			return err
		}
		if verbose {
			logXML(mc)
		}
	}

	return nil
}

func organizePlaylist(plexClient *plexAPI.PlexClient, playlist plexAPI.SavedPlaylist, playlistKey string) error {
	fmt.Printf("Organize playlist %s\n", playlist.Title)
	mc, err := plexClient.Get(playlistKey)
	if err != nil {
		return err
	}

	var lastItemID string
	for _, v := range playlist.Videos {
		fmt.Printf("	Position %s\n", v.Title)

		video, err := plexClient.FindVideo(v)
		if err != nil {
			return err
		}

		for _, plv := range mc.Videos {
			if video.Key == plv.Key {
				newItemID := plv.PlaylistItemID

				request := playlistKey + "/" + newItemID + "/move"
				if lastItemID != "" {
					request = request + "?after=" + lastItemID
				}
				result, err := plexClient.Put(request)
				if err != nil {
					return err
				}
				if verbose {
					logXML(result)
				}

				lastItemID = newItemID
			}
		}
	}

	return nil
}

/*
 * I/O
 */
func saveXML(data interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "\t")

	return encoder.Encode(data)
}

func loadXML(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	return xml.NewDecoder(file).Decode(data)
}
