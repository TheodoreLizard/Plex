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

	plexClient := plexAPI.NewPlexClient(plexAddr)
	mc, err := plexClient.Get("playlists")
	if err != nil {
		return playlists, err
	}

	for _, pl := range mc.Playlists {
		playList := plexAPI.SavedPlaylist{
			Title: pl.Title,
			Key:   pl.Key,
		}

		mc, err = plexClient.Get(pl.Key)
		if err != nil {
			return playlists, err
		}

		for _, v := range mc.Videos {
			vmc, err := plexClient.Get(v.Key)
			if err != nil {
				return playlists, err
			}

			video := plexAPI.SavedVideo{
				LibrarySectionUUID: vmc.LibrarySectionUUID,
				Title:              v.Title,
				Year:               v.Year,
				Key:                v.Key,
			}

			playList.Videos = append(playList.Videos, video)
		}

		playlists.Playlists = append(playlists.Playlists, playList)
	}

	return playlists, nil
}

func restorePlaylists(playlists plexAPI.Playlists) error {
	plexClient := plexAPI.NewPlexClient(plexAddr)

	for _, pl := range playlists.Playlists {
		playlist, err := plexClient.SearchPlaylist(pl.Title)
		if err != nil {
			return err
		}

		if pl.Key != playlist.Key {
			return fmt.Errorf("playlist key mismatch, saved: [%v], found: [%v]", pl.Key, playlist.Key)
		}

		for _, v := range pl.Videos {
			fmt.Printf("Restore %s to %s\n", v.Title, pl.Title)

			videos, err := plexClient.SearchLocal(v.Title)
			if err != nil {
				return err
			}

			if len(videos) == 0 {
				return fmt.Errorf("found %d videos matching title: %v", len(videos), v.Title)
			}

			for _, video := range videos {
				if v.Title == video.Title && v.Year == video.Year {
					videoURI := fmt.Sprintf("library://%s/item/%s", video.LibrarySectionUUID, video.Key)
					request := fmt.Sprintf("%s?uri=%s", playlist.Key, url.QueryEscape(videoURI))

					// fmt.Printf("%v\n", request)

					mc, err := plexClient.Put(request)
					if err != nil {
						return err
					}
					if verbose {
						logXML(mc)
					}

					break
				}
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
