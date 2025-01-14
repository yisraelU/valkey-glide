// Copyright Valkey GLIDE Project Contributors - SPDX Identifier: Apache-2.0

package options

// This struct represents the optional arguments for the ZSCAN command.
type ZScanOptions struct {
	BaseScanOptions
	noScores bool
}

func NewZScanOptionsBuilder() *ZScanOptions {
	return &ZScanOptions{}
}

/*
If this value is set to true, the ZSCAN command will be called with NOSCORES option.
In the NOSCORES option, scores are not included in the response.
*/
func (zScanOptions *ZScanOptions) SetNoScores(noScores bool) *ZScanOptions {
	zScanOptions.noScores = noScores
	return zScanOptions
}

func (zScanOptions *ZScanOptions) SetMatch(match string) *ZScanOptions {
	zScanOptions.BaseScanOptions.SetMatch(match)
	return zScanOptions
}

func (zScanOptions *ZScanOptions) SetCount(count int64) *ZScanOptions {
	zScanOptions.BaseScanOptions.SetCount(count)
	return zScanOptions
}

func (options *ZScanOptions) ToArgs() ([]string, error) {
	args := []string{}
	baseArgs, err := options.BaseScanOptions.ToArgs()
	args = append(args, baseArgs...)

	if options.noScores {
		args = append(args, NoScores)
	}
	return args, err
}
