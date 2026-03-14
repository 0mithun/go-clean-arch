package strategies

type CreateRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,max=500"`
}

func (r CreateRequest) Messages() map[string]string {
	return map[string]string{
		"name.required": "Please provide strategy name.",
		"name.min":      "Strategy name must be at least 3 characters.",
		"*.required":    "This field is required.",
	}
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
