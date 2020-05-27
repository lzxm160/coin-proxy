package database

import (
	"encoding/json"

	"github.com/scorum/scorum-go/types"
)

type DynamicGlobalProperties struct {
	ID                       uint32     `json:"id"`
	Time                     types.Time `json:"time"`
	HeadBlockNumber          uint32     `json:"head_block_number"`
	HeadBlockID              string     `json:"head_block_id"`
	CurrentWitness           string     `json:"current_witness"`
	TotalSupply              string     `json:"total_supply"`
	AccountsCurrentSupply    string     `json:"accounts_current_supply"`
	ConfidentialSupply       string     `json:"confidential_supply"`
	TotalVestingFundScorum   string     `json:"total_vesting_fund_scorum"`
	TotalVestingShares       string     `json:"total_vesting_shares"`
	TotalRewardShares2       string     `json:"total_reward_shares2"`
	MaximumBlockSize         int32      `json:"maximum_block_size"`
	CurrentAslot             int32      `json:"current_aslot"`
	RecentSlotsFilled        string     `json:"recent_slots_filled"`
	ParticipationCount       int32      `json:"participation_count"`
	LastIrreversibleBlockNum uint32     `json:"last_irreversible_block_num"`
	VotePowerReserveRate     int32      `json:"vote_power_reserve_rate"`
	InviteQuorum             int32      `json:"invite_quorum"`
	DropoutQuorum            int32      `json:"dropout_quorum"`
	ChangeQuorum             int32      `json:"change_quorum"`
	CurrentReserveRatio      int32      `json:"current_reserve_ratio"`
	AverageBlockSize         int32      `json:"average_block_size"`
	MaxVirtualBandwidth      string     `json:"max_virtual_bandwidth"`
}

type Config struct {
	IsTestNet                                           bool   `json:"IS_TEST_NET"`
	Scorum100Percent                                    int32  `json:"SCORUM_100_PERCENT"`
	Scorum1Percent                                      int32  `json:"SCORUM_1_PERCENT"`
	Scorum1TenthPercent                                 int32  `json:"SCORUM_1_TENTH_PERCENT"`
	ScorumAccountRecoveryRequestExpirationPeriod        string `json:"SCORUM_ACCOUNT_RECOVERY_REQUEST_EXPIRATION_PERIOD"`
	ScorumAddressPrefix                                 string `json:"SCORUM_ADDRESS_PREFIX"`
	ScorumAprPercentMultiplyPerBlock                    string `json:"SCORUM_APR_PERCENT_MULTIPLY_PER_BLOCK"`
	ScorumAprPercentMultiplyPerHour                     string `json:"SCORUM_APR_PERCENT_MULTIPLY_PER_HOUR"`
	ScorumAprPercentMultiplyPerRound                    string `json:"SCORUM_APR_PERCENT_MULTIPLY_PER_ROUND"`
	ScorumAprPercentShiftPerBlock                       int32  `json:"SCORUM_APR_PERCENT_SHIFT_PER_BLOCK"`
	ScorumAprPercentShiftPerHour                        int32  `json:"SCORUM_APR_PERCENT_SHIFT_PER_HOUR"`
	ScorumAprPercentShiftPerRound                       int32  `json:"SCORUM_APR_PERCENT_SHIFT_PER_ROUND"`
	ScorumBandwidthAverageWindowSeconds                 int32  `json:"SCORUM_BANDWIDTH_AVERAGE_WINDOW_SECONDS"`
	ScorumBandwidthPrecision                            int32  `json:"SCORUM_BANDWIDTH_PRECISION"`
	ScorumBlockchainHardforkVersion                     string `json:"SCORUM_BLOCKCHAIN_HARDFORK_VERSION"`
	ScorumBlockchainVersion                             string `json:"SCORUM_BLOCKCHAIN_VERSION"`
	ScorumBlockInterval                                 int32  `json:"SCORUM_BLOCK_INTERVAL"`
	ScorumBlocksPerDay                                  int32  `json:"SCORUM_BLOCKS_PER_DAY"`
	ScorumBlocksPerHour                                 int32  `json:"SCORUM_BLOCKS_PER_HOUR"`
	ScorumBlocksPerYear                                 int32  `json:"SCORUM_BLOCKS_PER_YEAR"`
	ScorumCashoutWindowSeconds                          int32  `json:"SCORUM_CASHOUT_WINDOW_SECONDS"`
	ScorumContentAprPercent                             int32  `json:"SCORUM_CONTENT_APR_PERCENT"`
	ScorumContentRewardPercent                          int32  `json:"SCORUM_CONTENT_REWARD_PERCENT"`
	ScorumCreateAccountDelegationRatio                  int32  `json:"SCORUM_CREATE_ACCOUNT_DELEGATION_RATIO"`
	ScorumCreateAccountDelegationTime                   string `json:"SCORUM_CREATE_ACCOUNT_DELEGATION_TIME"`
	ScorumCreateAccountWithScorumModifier               int32  `json:"SCORUM_CREATE_ACCOUNT_WITH_SCORUM_MODIFIER"`
	ScorumHardforkRequiredWitnesses                     int32  `json:"SCORUM_HARDFORK_REQUIRED_WITNESSES"`
	ScorumInflationNarrowingPeriod                      int32  `json:"SCORUM_INFLATION_NARROWING_PERIOD"`
	ScorumInflationRateStartPercent                     int32  `json:"SCORUM_INFLATION_RATE_START_PERCENT"`
	ScorumInflationRateStopPercent                      int32  `json:"SCORUM_INFLATION_RATE_STOP_PERCENT"`
	ScorumIrreversibleThreshold                         int32  `json:"SCORUM_IRREVERSIBLE_THRESHOLD"`
	ScorumMaxAccountNameLength                          int32  `json:"SCORUM_MAX_ACCOUNT_NAME_LENGTH"`
	ScorumMaxAccountWitnessVotes                        int32  `json:"SCORUM_MAX_ACCOUNT_WITNESS_VOTES"`
	ScorumMaxBlockSize                                  int32  `json:"SCORUM_MAX_BLOCK_SIZE"`
	ScorumMaxCommentDepth                               int32  `json:"SCORUM_MAX_COMMENT_DEPTH"`
	ScorumMaxFeedAgeSeconds                             int32  `json:"SCORUM_MAX_FEED_AGE_SECONDS"`
	ScorumMaxMemoSize                                   int32  `json:"SCORUM_MAX_MEMO_SIZE"`
	ScorumMaxWitnesses                                  int32  `json:"SCORUM_MAX_WITNESSES"`
	ScorumMaxPermlinkLength                             int32  `json:"SCORUM_MAX_PERMLINK_LENGTH"`
	ScorumMaxProxyRecursionDepth                        int32  `json:"SCORUM_MAX_PROXY_RECURSION_DEPTH"`
	ScorumMaxReserveRatio                               int32  `json:"SCORUM_MAX_RESERVE_RATIO"`
	ScorumMaxRunnerWitnesses                            int32  `json:"SCORUM_MAX_RUNNER_WITNESSES"`
	ScorumMaxShareSupply                                string `json:"SCORUM_MAX_SHARE_SUPPLY"`
	ScorumMaxSigCheckDepth                              int32  `json:"SCORUM_MAX_SIG_CHECK_DEPTH"`
	ScorumMaxTimeUntilExpiration                        int32  `json:"SCORUM_MAX_TIME_UNTIL_EXPIRATION"`
	ScorumMaxTransactionSize                            int32  `json:"SCORUM_MAX_TRANSACTION_SIZE"`
	ScorumMaxUndoHistory                                int32  `json:"SCORUM_MAX_UNDO_HISTORY"`
	ScorumMaxVoteChanges                                int32  `json:"SCORUM_MAX_VOTE_CHANGES"`
	ScorumMaxVotedWitnesses                             int32  `json:"SCORUM_MAX_VOTED_WITNESSES"`
	ScorumMaxWithdrawRoutes                             int32  `json:"SCORUM_MAX_WITHDRAW_ROUTES"`
	ScorumMaxWitnessUrlLength                           int32  `json:"SCORUM_MAX_WITNESS_URL_LENGTH"`
	ScorumMinAccountCreationFee                         string `json:"SCORUM_MIN_ACCOUNT_CREATION_FEE"`
	ScorumMinAccountNameLength                          int32  `json:"SCORUM_MIN_ACCOUNT_NAME_LENGTH"`
	ScorumMinBlockSizeLimit                             int32  `json:"SCORUM_MIN_BLOCK_SIZE_LIMIT"`
	ScorumMinContentReward                              string `json:"SCORUM_MIN_CONTENT_REWARD"`
	ScorumMinCurateReward                               string `json:"SCORUM_MIN_CURATE_REWARD"`
	ScorumMinPermlinkLength                             int32  `json:"SCORUM_MIN_PERMLINK_LENGTH"`
	ScorumMinReplyInterval                              int32  `json:"SCORUM_MIN_REPLY_INTERVAL"`
	ScorumMinRootCommentInterval                        int32  `json:"SCORUM_MIN_ROOT_COMMENT_INTERVAL"`
	ScorumMinVoteIntervalSec                            int32  `json:"SCORUM_MIN_VOTE_INTERVAL_SEC"`
	ScorumMinFeeds                                      int32  `json:"SCORUM_MIN_FEEDS"`
	ScorumDefaultReward                                 string `json:"SCORUM_DEFAULT_REWARD"`
	ScorumMinPayout                                     string `json:"SCORUM_MIN_PAYOUT"`
	ScorumMinPowReward                                  string `json:"SCORUM_MIN_POW_REWARD"`
	ScorumMinProducerReward                             string `json:"SCORUM_MIN_PRODUCER_REWARD"`
	ScorumMinTransactionExpirationLimit                 int32  `json:"SCORUM_MIN_TRANSACTION_EXPIRATION_LIMIT"`
	ScorumMinUndoHistory                                int32  `json:"SCORUM_MIN_UNDO_HISTORY"`
	ScorumNumInitDelegates                              int32  `json:"SCORUM_NUM_INIT_DELEGATES"`
	ScorumOwnerAuthHistoryTrackingStartBlockNum         int32  `json:"SCORUM_OWNER_AUTH_HISTORY_TRACKING_START_BLOCK_NUM"`
	ScorumOwnerAuthRecoveryPeriod                       string `json:"SCORUM_OWNER_AUTH_RECOVERY_PERIOD"`
	ScorumOwnerUpdateLimit                              int64  `json:"SCORUM_OWNER_UPDATE_LIMIT"`
	ScorumPowAprPercent                                 int32  `json:"SCORUM_POW_APR_PERCENT"`
	ScorumProducerAprPercent                            int32  `json:"SCORUM_PRODUCER_APR_PERCENT"`
	ScorumProxyToSelfAccount                            string `json:"SCORUM_PROXY_TO_SELF_ACCOUNT"`
	ScorumRecentRsharesDecayRate                        string `json:"SCORUM_RECENT_RSHARES_DECAY_RATE"`
	ScorumReverseAuctionWindowSeconds                   int32  `json:"SCORUM_REVERSE_AUCTION_WINDOW_SECONDS"`
	ScorumRootPostParent                                string `json:"SCORUM_ROOT_POST_PARENT"`
	ScorumSavingsWithdrawRequestLimit                   int32  `json:"SCORUM_SAVINGS_WITHDRAW_REQUEST_LIMIT"`
	ScorumSavingsWithdrawTime                           string `json:"SCORUM_SAVINGS_WITHDRAW_TIME"`
	ScorumSoftMaxCommentDepth                           int32  `json:"SCORUM_SOFT_MAX_COMMENT_DEPTH"`
	ScorumStartMinerVotingBlock                         int32  `json:"SCORUM_START_MINER_VOTING_BLOCK"`
	ScorumStartVestingBlock                             int32  `json:"SCORUM_START_VESTING_BLOCK"`
	ScorumUpvoteLockout                                 int32  `json:"SCORUM_UPVOTE_LOCKOUT"`
	ScorumVestingFundPercent                            int32  `json:"SCORUM_VESTING_FUND_PERCENT"`
	ScorumVestingWithdrawIntervals                      int32  `json:"SCORUM_VESTING_WITHDRAW_INTERVALS"`
	ScorumVestingWithdrawIntervalsPreHf16               int32  `json:"SCORUM_VESTING_WITHDRAW_INTERVALS_PRE_HF_16"`
	ScorumVestingWithdrawIntervalSeconds                int32  `json:"SCORUM_VESTING_WITHDRAW_INTERVAL_SECONDS"`
	ScorumVoteDustThreshold                             int32  `json:"SCORUM_VOTE_DUST_THRESHOLD"`
	ScorumVoteRegenerationSeconds                       string `json:"SCORUM_VOTE_REGENERATION_SECONDS"`
	ScorumSymbol                                        int32  `json:"SCORUM_SYMBOL"`
	VestsSymbol                                         int32  `json:"VESTS_SYMBOL"`
	VirtualScheduleLapLength                            string `json:"VIRTUAL_SCHEDULE_LAP_LENGTH"`
	ScorumRewardsInitialSupply                          string `json:"SCORUM_REWARDS_INITIAL_SUPPLY"`
	ScorumRewardsInitialSupplyPeriodInDays              int32  `json:"SCORUM_REWARDS_INITIAL_SUPPLY_PERIOD_IN_DAYS"`
	ScorumGuarantedRewardSupplyPeriodInDays             int32  `json:"SCORUM_GUARANTED_REWARD_SUPPLY_PERIOD_IN_DAYS"`
	ScorumRewardIncreaseThresholdInDays                 int32  `json:"SCORUM_REWARD_INCREASE_THRESHOLD_IN_DAYS"`
	ScorumAdjustRewardPercent                           int32  `json:"SCORUM_ADJUST_REWARD_PERCENT"`
	ScorumBudgetLimitCountPerOwner                      int32  `json:"SCORUM_BUDGET_LIMIT_COUNT_PER_OWNER"`
	ScorumBudgetLimitDbListSize                         int32  `json:"SCORUM_BUDGET_LIMIT_DB_LIST_SIZE"`
	ScorumBudgetLimitApiListSize                        int32  `json:"SCORUM_BUDGET_LIMIT_API_LIST_SIZE"`
	ScorumRegistrationBonusLimitPerMemberNBlock         int32  `json:"SCORUM_REGISTRATION_BONUS_LIMIT_PER_MEMBER_N_BLOCK"`
	ScorumRegistrationBonusLimitPerMemberPerNBlock      int32  `json:"SCORUM_REGISTRATION_BONUS_LIMIT_PER_MEMBER_PER_N_BLOCK"`
	ScorumRegistrationLimitCountCommitteeMembers        int32  `json:"SCORUM_REGISTRATION_LIMIT_COUNT_COMMITTEE_MEMBERS"`
	ScorumWitnessMissedBlocksThreshold                  int32  `json:"SCORUM_WITNESS_MISSED_BLOCKS_THRESHOLD"`
	ScorumAtomicswapInitiatorRefundLockSecs             int32  `json:"SCORUM_ATOMICSWAP_INITIATOR_REFUND_LOCK_SECS"`
	ScorumAtomicswapParticipantRefundLockSecs           int32  `json:"SCORUM_ATOMICSWAP_PARTICIPANT_REFUND_LOCK_SECS"`
	ScorumAtomicswapLimitRequestedContractsPerOwner     int32  `json:"SCORUM_ATOMICSWAP_LIMIT_REQUESTED_CONTRACTS_PER_OWNER"`
	ScorumAtomicswapLimitRequestedContractsPerRecipient int32  `json:"SCORUM_ATOMICSWAP_LIMIT_REQUESTED_CONTRACTS_PER_RECIPIENT"`
	ScorumAtomicswapContractMetadataMaxLength           int32  `json:"SCORUM_ATOMICSWAP_CONTRACT_METADATA_MAX_LENGTH"`
	ScorumAtomicswapSecretMaxLength                     int32  `json:"SCORUM_ATOMICSWAP_SECRET_MAX_LENGTH"`
}

type Account struct {
	ID                        uint32            `json:"id"`
	Name                      string            `json:"name"`
	Owner                     types.Authority   `json:"owner"`
	Active                    types.Authority   `json:"active"`
	Posting                   types.Authority   `json:"posting"`
	MemoKey                   string            `json:"memo_key"`
	JsonMetadata              string            `json:"json_metadata"`
	Proxy                     string            `json:"proxy"`
	LastOwnerUpdate           types.Time        `json:"last_owner_update"`
	LastAccountUpdate         types.Time        `json:"last_account_update"`
	Created                   types.Time        `json:"created"`
	CreatedByGenesis          bool              `json:"created_by_genesis"`
	OwnerChallenged           bool              `json:"owner_challenged"`
	ActiveChallenged          bool              `json:"active_challenged"`
	LastOwnerProved           types.Time        `json:"last_owner_proved"`
	LastActiveProved          types.Time        `json:"last_active_proved"`
	RecoveryAccount           string            `json:"recovery_account"`
	LastAccountRecovery       types.Time        `json:"last_account_recovery"`
	CommentCount              int32             `json:"comment_count"`
	LifetimeVoteCount         int32             `json:"lifetime_vote_count"`
	PostCount                 int32             `json:"post_count"`
	CanVote                   bool              `json:"can_vote"`
	VotingPower               int32             `json:"voting_power"`
	LastVoteTime              types.Time        `json:"last_vote_time"`
	Balance                   string            `json:"balance"`
	VestingShares             string            `json:"vesting_shares"`
	DelegatedVestingShares    string            `json:"delegated_vesting_shares"`
	ReceivedVestingShares     string            `json:"received_vesting_shares"`
	VestingWithdrawRate       string            `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal     types.Time        `json:"next_vesting_withdrawal"`
	CurationRewards           string            `json:"curation_rewards"`
	PostingRewards            string            `json:"posting_rewards"`
	ProxiedVsfVotes           []json.RawMessage `json:"proxied_vsf_votes"`
	WitnessesVotedFor         json.RawMessage   `json:"witnesses_voted_for"`
	AverageBandwidth          json.RawMessage   `json:"average_bandwidth"`
	LifetimeBandwidth         json.RawMessage   `json:"lifetime_bandwidth"`
	LastBandwidthUpdate       types.Time        `json:"last_bandwidth_update"`
	AverageMarketBandwidth    json.RawMessage   `json:"average_market_bandwidth"`
	LifetimeMarketBandwidth   json.RawMessage   `json:"lifetime_market_bandwidth"`
	LastMarketBandwidthUpdate types.Time        `json:"last_market_bandwidth_update"`
	LastPost                  types.Time        `json:"last_post"`
	LastRootPost              types.Time        `json:"last_root_post"`
	VestingBalance            string            `json:"vesting_balance"`
	TransferHistory           []json.RawMessage `json:"transfer_history"`
	PostHistory               []json.RawMessage `json:"post_history"`
	VoteHistory               []json.RawMessage `json:"vote_history"`
	OtherHistory              []json.RawMessage `json:"other_history"`
	WitnessVotes              []string          `json:"witness_votes"`
	TagsUsage                 []json.RawMessage `json:"tags_usage"`
	GuestBloggers             []json.RawMessage `json:"guest_bloggers"`
}