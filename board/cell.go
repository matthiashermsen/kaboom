package board

type Cell struct {
	IsRevealed                 bool
	HasMine                    bool
	HasFlag                    bool
	AmountOfNeighboursWithMine uint
}
