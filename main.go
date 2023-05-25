package main

import (
	"create_mvx_wallet/mvx"
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
}

func main() {
	// Log settings
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Parse flags
	walletName := flag.String("wallet", "multiversx", "Allows you to retrieve a wallet containing the specified name")
	logLevel := flag.String("log", "info", "Set the log level: debug, info, warn, error, fatal, panic")
	flag.Parse()

	// Set log level
	switch *logLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	/*
	 Check if folder exists
	*/
	err := createFolders(*walletName)
	if err != nil {
		log.Fatal().Err(err).Msg("error creating folders")
	}

	/*
		Create wallet
	*/
	var i int
	for {
		w := &mvx.Wallet{}
		err := w.CreateWalletShard(*walletName)
		if err != nil {
			log.Fatal().Err(err).Msg("error creating wallet")
		}
		i++
		if i == 60 {
			log.Info().Msg("Sleeping for 1 minute to avoid rate limit..")
			time.Sleep(1 * time.Minute)
			i = 0
		}
	}
}

func createFolders(walletName string) error {
	if _, err := os.Stat("./" + walletName); os.IsNotExist(err) {
		err = os.Mkdir("./"+walletName, 0755)
		if err != nil {
			return err
		}
	}

	for i := 0; i < 3; i++ {
		// Check if folder exists
		shardNumber := strconv.Itoa(i)
		if _, err := os.Stat("./shard" + shardNumber); os.IsNotExist(err) {
			err = os.Mkdir("./shard"+shardNumber, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
