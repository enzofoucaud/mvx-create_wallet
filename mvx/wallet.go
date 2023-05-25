package mvx

import (
	"os"
	"strconv"
	"strings"

	"github.com/enzofoucaud/mvx-create_wallet/mvx_api"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/interactors"
	"github.com/rs/zerolog/log"
)

type Wallet struct {
	PrivateKey    []byte
	Address       core.AddressHandler
	AddressBech32 []byte
}

func (w *Wallet) CreateWalletShard(walletName string) error {
	// generating a mnemonic
	iW := interactors.NewWallet()
	mnemonic, err := iW.GenerateMnemonic()
	if err != nil {
		log.Fatal().Err(err).Msg("error generating mnemonic")
		return err
	}
	log.Debug().Interface("mnemonic", mnemonic).Msg("generated mnemonic")

	// generating the private key from the mnemonic using index 0
	index := uint32(0)
	w.PrivateKey = iW.GetPrivateKeyFromMnemonic(mnemonic, 0, index)
	w.Address, err = iW.GetAddressFromPrivateKey(w.PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting address from private key")
		return err
	}

	acc, err := mvx_api.RequestAccountDetail(w.Address.AddressAsBech32String())
	if err != nil {
		log.Fatal().Err(err).Msg("error getting account detail")
		return err
	}
	log.Info().Str("address as bech32", w.Address.AddressAsBech32String()).Int("shard", acc.Shard).Msg("generated private/public key")

	shardInt := strconv.Itoa(acc.Shard)
	if strings.Contains(acc.Address, walletName) {
		err = iW.SavePrivateKeyToPemFile(w.PrivateKey, "./"+walletName+"/"+w.Address.AddressAsBech32String()+".pem")
		if err != nil {
			log.Fatal().Err(err).Msg("error saving private key to pem file")
			return err
		}

		f, err := os.Create("./upong/" + w.Address.AddressAsBech32String() + ".txt")
		if err != nil {
			log.Fatal().Err(err).Msg("error creating file")
			return err
		}
		defer f.Close()
		_, err = f.WriteString(string(mnemonic) + "\n")
		if err != nil {
			log.Fatal().Err(err).Msg("error writing to file")
			return err
		}
	} else {
		err = iW.SavePrivateKeyToPemFile(w.PrivateKey, "./shard"+shardInt+"/"+w.Address.AddressAsBech32String()+".pem")
		if err != nil {
			log.Fatal().Err(err).Msg("error saving private key to pem file")
			return err
		}

		f, err := os.Create("./shard" + shardInt + "/" + w.Address.AddressAsBech32String() + ".txt")
		if err != nil {
			log.Fatal().Err(err).Msg("error creating file")
			return err
		}
		defer f.Close()
		_, err = f.WriteString(string(mnemonic) + "\n")
		if err != nil {
			log.Fatal().Err(err).Msg("error writing to file")
			return err
		}
	}

	log.Info().Msg("saved private key to pem file")

	return nil
}
