package mvx

import (
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/enzofoucaud/mvx-create_wallet/mvx_api"
	"github.com/multiversx/mx-sdk-go/core"
	"github.com/multiversx/mx-sdk-go/data"
	"github.com/multiversx/mx-sdk-go/interactors"
	"github.com/rs/zerolog/log"
)

func CreateWalletShard(walletName string) error {
	var (
		PrivateKey []byte
		Address    core.AddressHandler
	)

	// creating a new wallet interactor
	wallet := interactors.NewWallet()

	// generating a mnemonic
	mnemonic, err := wallet.GenerateMnemonic()
	if err != nil {
		log.Fatal().Err(err).Msg("error generating mnemonic")
		return err
	}
	log.Debug().Interface("mnemonic", mnemonic).Msg("generated mnemonic")

	// generating the private key from the mnemonic using index 0
	PrivateKey = wallet.GetPrivateKeyFromMnemonic(mnemonic, 0, uint32(0))
	Address, err = wallet.GetAddressFromPrivateKey(PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("error getting address from private key")
		return err
	}
	log.Info().Str("address as bech32", Address.AddressAsBech32String()).Msg("generated private/public key")

	// getting the account detail
	acc, err := mvx_api.RequestAccountDetail(Address.AddressAsBech32String())
	if err != nil {
		log.Fatal().Err(err).Msg("error getting account detail")
		return err
	}
	log.Info().Int("shard", acc.Shard).Msg("got account detail")

	if strings.Contains(acc.Address, walletName) {
		folder := walletName
		err = wallet.SavePrivateKeyToPemFile(PrivateKey, path.Join("./", folder, Address.AddressAsBech32String()+".pem"))
		if err != nil {
			log.Fatal().Err(err).Msg("error saving private key to pem file")
			return err
		}
		SaveTxtWallet(folder, Address.AddressAsBech32String(), mnemonic)
	} else {
		folder := "shard" + strconv.Itoa(acc.Shard)
		err = wallet.SavePrivateKeyToPemFile(PrivateKey, path.Join("./", folder, Address.AddressAsBech32String()+".pem"))
		if err != nil {
			log.Fatal().Err(err).Msg("error saving private key to pem file")
			return err
		}
		SaveTxtWallet(folder, Address.AddressAsBech32String(), mnemonic)
	}

	log.Info().Msg("saved private key to pem file")

	return nil
}

func SaveTxtWallet(folder, address string, mnemonic data.Mnemonic) error {
	f, err := os.Create("./" + folder + "/" + address + ".txt")
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
	return nil
}
