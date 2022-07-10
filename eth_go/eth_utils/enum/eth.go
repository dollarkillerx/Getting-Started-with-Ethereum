package enum

type BlockParameter string

const (
	Earliest BlockParameter = "earliest" // 創世塊
	Latest   BlockParameter = "latest"   // 已確定的
	Pending  BlockParameter = "pending"  // 進行中的
)
