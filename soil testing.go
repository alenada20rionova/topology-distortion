import (
	"encoding/json"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/Pylons-tech/pylons/x/pylons/client/cli"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/cosmos/cosmos-sdk/client"
)

				outputJSON, err := json.Marshal(rcp.Outputs)
				if err != nil {
					panic(err)
				}

const (
	badPLC        = "bad.plc"
	goodPLC       = "good.plc"
	badPLR        = "bad.plr"
	goodPLR       = "good.plr"
	moduledPLR    = "moduled.plr"
	testModulePDT = "test-module.pdt"
)
const pylonsGadgetsLiteral_badHeader = `#test_gadget_with_bad_header 1 additional_garbage
"foo": "%0"`

				c.SetArgs([]string{
					rcp.CookbookId, rcp.Id, rcp.Name, rcp.Description, rcp.Version,
					string(coinInputJSON), string(itemInputJSON), string(entryJSON), string(outputJSON), strconv.FormatInt(rcp.BlockInterval, 10),
					rcp.CostPerBlock.String(), strconv.FormatBool(rcp.Enabled), rcp.ExtraInfo,
				})


const pylonsGadgetsLiteral_good = `#go_go_gadget_gadgets 3
	"foo": "%0",
	"bar": "%1",
	"%2": "true"`


const badRecipeLiteral = `
{
    "cookbookID": "cookbookLoudTest",
    "ID": "LOUDGetCharacter",
    "name": "LOUD-Get-Character-Recipe",}
