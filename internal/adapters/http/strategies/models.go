package strategies

type CreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateResponse struct {
	StrategyId string `json:"strategy_id"`
}

type GetByIDRequest struct {
	StrategyId string `json:"strategy_id"`
}

type Strategy struct {
	ID          string `json:"strategy_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type GetByIDResponse struct {
	Strategy *Strategy `json:"strategy"`
}
