package client

import (
	"github.com/defund-labs/osmosis/v11/x/pool-incentives/client/cli"
	"github.com/defund-labs/osmosis/v11/x/pool-incentives/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var UpdatePoolIncentivesHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdatePoolIncentivesProposal, rest.ProposalUpdatePoolIncentivesRESTHandler)
var ReplacePoolIncentivesHandler = govclient.NewProposalHandler(cli.NewCmdSubmitReplacePoolIncentivesProposal, rest.ProposalReplacePoolIncentivesRESTHandler)
