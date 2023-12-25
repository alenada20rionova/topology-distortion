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
