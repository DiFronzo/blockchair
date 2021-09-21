package blockchair

import (
	"regexp"
)

// Contains used with GetSupportedCrypto and/or GetSupportedCryptoEth to verify correct crypto.
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ValidateCrypto validate Bitcoin-like crypto.
func (c *Client) ValidateCrypto(crypto string) error {
	if !Contains(GetSupportedCrypto(), crypto) {
		return c.err1(ErrSC)
	}

	return nil
}

// ValidateCryptoEth validate Ethereum crypto.
func (c *Client) ValidateCryptoEth(crypto string) error {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		return c.err1(ErrSCE)
	}

	return nil
}

// ValidateCryptoBoth validate both Bitcoin-like and Ethereum crypto.
func (c *Client) ValidateCryptoBoth(crypto string) error {
	if !Contains(GetSupportedCryptoEth(), crypto) && !Contains(GetSupportedCrypto(), crypto) {
		return c.err1(ErrSCG)
	}

	return nil
}

// ValidateCryptoMultichain validate crypto for multichain address check.
func (c *Client) ValidateCryptoMultichain(crypto string) error {
	if !Contains(GetSupportedCryptoMultichain(), crypto) {
		return c.err1(ErrSCG)
	}

	return nil
}

// ValidateHashEth validate Ethereum hash.
func (c *Client) ValidateHashEth(hash string) error {
	r, _ := regexp.Compile(Hash)
	if !r.MatchString(hash) {
		return c.err4(ErrTHW, hash)
	}

	return nil
}

// ValidateHashesEth validate Ethereum hashes.
func (c *Client) ValidateHashesEth(hashes []string) error {
	r, _ := regexp.Compile(Hash)
	for i := range hashes {
		if !r.MatchString(hashes[i]) {
			return c.err4(ErrTHW, hashes[i])
		}
	}

	return nil
}

// ValidateErc20Token validate ERC-20 token.
func (c *Client) ValidateErc20Token(token string) error {
	r, _ := regexp.Compile("0x[0-9a-fA-F]{40}")
	if !r.MatchString(token) {
		return c.err4(ErrERC, token)
	}

	return nil
}

// ValidateErc20Tokens validate ERC-20 tokens.
func (c *Client) ValidateErc20Tokens(tokens []string) error {
	r, _ := regexp.Compile("0x[0-9a-fA-F]{40}")
	for i := range tokens {
		if !r.MatchString(tokens[i]) {
			return c.err4(ErrERC, tokens[i])
		}
	}

	return nil
}
